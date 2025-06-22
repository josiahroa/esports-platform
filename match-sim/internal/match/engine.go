package match

import (
	"log/slog"
	"maps"
	"match-sim/internal/constants"
	"match-sim/internal/event"
	"math/rand"
	"slices"
	"time"
)

func (match Match) SimulateMatch(seed int64, isRealTime bool) {
	matchEvent := event.NewEventLogger(match.ID)

	slog.Info("Simulating match", "match", match.ID)

	slog.Info("Loading teams")
	for _, team := range match.Teams {
		slog.Info("Loading players for team", "team", team.ID)
		for _, player := range team.Players {
			slog.Info("Loading player", "player", player.ID)
		}
	}

	// seed random number generator
	matchSeed := rand.NewSource(seed)
	matchRng := rand.New(matchSeed)

	slog.Info("Match RNG", "matchRng", matchRng)

	// load rules
	rules := DefaultRules()
	slog.Info("Rules", "rules", rules)

	// start game
	gameState := NewGameState(rules, isRealTime)
	gameState.TeamOne = &match.Teams[0]
	gameState.TeamTwo = &match.Teams[1]

	// decide map
	mapKeys := []constants.MapName{}
	for k := range constants.Maps {
		mapKeys = append(mapKeys, k)
	}
	slices.Sort(mapKeys)
	gameMap := constants.Maps[mapKeys[matchRng.Intn(len(mapKeys))]]
	gameState.Map = gameMap
	slog.Info("Map Selection", "map", gameState.Map.Name)

	// decide starting sides
	attackingTeamIndex := matchRng.Intn(2)
	gameState.AttackingTeam = &match.Teams[attackingTeamIndex]
	gameState.DefendingTeam = &match.Teams[1-attackingTeamIndex]

	slog.Info("Attacking", "team", gameState.AttackingTeam.ID)
	slog.Info("Defending", "team", gameState.DefendingTeam.ID)

	// simulate agent selection
	match.Teams[0].AgentSelect(constants.Agents, matchRng)
	match.Teams[1].AgentSelect(constants.Agents, matchRng)

	// initialize player game state
	// used to track the player's stats throughout a match
	gameState.InitPlayerGameState(&match.Teams[0])
	gameState.InitPlayerGameState(&match.Teams[1])

	// start match
	gameState.GameRunning = true
	for gameState.GameRunning {
		// Swap teams at the end of the half
		if gameState.CurrentRound.RoundNumber == 12 {
			gameState.CurrentHalf = 2
			gameState.AttackingTeam, gameState.DefendingTeam = gameState.DefendingTeam, gameState.AttackingTeam
			matchEvent.SetAttackingTeam(gameState.AttackingTeam.ID)
			matchEvent.SetDefendingTeam(gameState.DefendingTeam.ID)
		}

		gameState.StartRound(gameState.AttackingTeam, gameState.DefendingTeam, matchRng)
		// Get round winner and loser and up their score here
		roundWinner := gameState.CurrentRound.RoundWinner
		if roundWinner == gameState.TeamOne {
			gameState.TeamOneRoundsWon++
		} else if roundWinner == gameState.TeamTwo {
			gameState.TeamTwoRoundsWon++
		}

		// Check if game goes to overtime
		// TODO: handle overtime

		// Check if a team has won the match
		if gameState.TeamOneRoundsWon == 13 || gameState.TeamTwoRoundsWon == 13 {
			gameState.GameRunning = false
			slog.Info("Match Ended", "teamOneScore", gameState.TeamOneRoundsWon, "teamTwoScore", gameState.TeamTwoRoundsWon)
			break
		}
	}
}

// Starts a round - runs the round logic for each tick until a team wins the round
func (gameState *GameState) StartRound(attackingTeam *Team, defendingTeam *Team, rng *rand.Rand) {
	gameState.CurrentRound = NewRoundState(gameState.CurrentRound.RoundNumber + 1)

	gameState.CurrentRound.AttackingTeam = attackingTeam
	gameState.CurrentRound.DefendingTeam = defendingTeam

	const tickRate uint = 128
	tickDuration := time.Second / time.Duration(tickRate)

	// Round loop - run this logic for each round until a team reaches 13 rounds won
	totalTicks := uint(gameState.Rules.RoundDurationSeconds) * tickRate

	// Initialize player round state
	gameState.CurrentRound.InitPlayerRoundState(attackingTeam)
	gameState.CurrentRound.InitPlayerRoundState(defendingTeam)

	simulator := NewGameTickSimulator(
		totalTicks,
		gameState,
		rng,
		tickRate,
		tickDuration,
	)

	for tick := uint(0); tick < simulator.TotalTicks; tick++ {
		exit := simulator.SimulateGameTick(&tick)
		if exit {
			break
		}
	}
}

// Before starting a game, init the player game state
func (gameState *GameState) InitPlayerGameState(team *Team) {
	for i := range team.Players {
		playerGameState := NewPlayerGameState(&team.Players[i])
		slog.Info("Init Player Game State", "player", playerGameState.LogValue())
		gameState.PlayerGameState = append(gameState.PlayerGameState, &playerGameState)
	}
}

// Post processing logic for ending a round
func (gameState *GameState) EndRound(winningTeam *Team, losingTeam *Team) {
	gameState.CurrentRound.RoundWinner = winningTeam
	gameState.CurrentRound.RoundLoser = losingTeam
	gameState.Rounds = append(gameState.Rounds, gameState.CurrentRound)

	slog.Info("End Round", "round", gameState.CurrentRound.RoundNumber, "winningTeam", winningTeam.ID, "losingTeam", losingTeam.ID, "roundWinner", winningTeam.ID, "roundLoser", losingTeam.ID)
}

func (roundState *RoundState) InitPlayerRoundState(team *Team) {
	for i := range team.Players {
		playerRoundState := NewPlayerRoundState(team, &team.Players[i], 0)
		// slog.Info("Init Player Round State", "player", playerRoundState.LogValue())
		roundState.PlayerRoundState = append(roundState.PlayerRoundState, &playerRoundState)
	}
}

func (roundState *RoundState) GetAlivePlayers(team *Team) []*PlayerRoundState {
	alivePlayers := make([]*PlayerRoundState, 0, 10)
	for _, p := range roundState.PlayerRoundState {
		if p.Alive && p.Team == team {
			alivePlayers = append(alivePlayers, p)
		}
	}
	return alivePlayers
}
func (roundState *RoundState) SimulatePlayerKill(rng *rand.Rand) PlayerKill {
	allAttackingTeamPlayers := roundState.GetAlivePlayers(roundState.AttackingTeam)
	allDefendingTeamPlayers := roundState.GetAlivePlayers(roundState.DefendingTeam)
	allAlivePlayers := append(allAttackingTeamPlayers, allDefendingTeamPlayers...)

	if len(allAlivePlayers) < 2 {
		return PlayerKill{}
	}

	// Pick a random killer.
	killer := allAlivePlayers[rng.Intn(len(allAlivePlayers))]

	// Pick a random victim from an opposing team.
	opponentsAlive := make([]*PlayerRoundState, 0, 5)
	for _, p := range allAlivePlayers {
		if p.Team != killer.Team {
			opponentsAlive = append(opponentsAlive, p)
		}
	}

	if len(opponentsAlive) == 0 {
		return PlayerKill{}
	}
	victim := opponentsAlive[rng.Intn(len(opponentsAlive))]
	victim.Alive = false

	return PlayerKill{
		Killer: killer.Player,
		Victim: victim.Player,
		Weapon: constants.Weapon{Name: constants.Sheriff}, //TODO: update this to the weapon used
	}
}

// Agent selection simulation logic
func (team Team) AgentSelect(agents map[constants.AgentName]constants.Agent, rng *rand.Rand) {
	availableAgents := make(map[constants.AgentName]constants.Agent)
	maps.Copy(availableAgents, agents)

	for i := range team.Players {
		possibleAgents := filterAgents(availableAgents, team.Players[i].Role)
		agent := possibleAgents[rng.Intn(len(possibleAgents))]
		// TODO: update event system

		team.Players[i].Agent = agent
		slog.Info("Agent Selection", "team", team.ID, "player", team.Players[i].ID, "agent", agent.Name)
		delete(availableAgents, agent.Name)
	}
}

// Filters agents by role - useful for getting a slice of agents for a given role
func filterAgents(agents map[constants.AgentName]constants.Agent, role constants.AgentRole) []constants.Agent {
	// In order for agent selection to be deterministic, we must iterate over the map in a consistent order.
	// We do this by extracting the keys, sorting them, and then iterating over the sorted keys.
	var keys []constants.AgentName
	for k := range agents {
		keys = append(keys, k)
	}
	slices.Sort(keys)

	filteredAgents := []constants.Agent{}
	for _, key := range keys {
		agent := agents[key]
		if agent.Role == role {
			filteredAgents = append(filteredAgents, agent)
		}
	}
	return filteredAgents
}

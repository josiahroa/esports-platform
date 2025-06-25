package match

import (
	"log/slog"
	"maps"
	"match-sim/internal/constants"
	"match-sim/internal/event"
	"math/rand"
	"slices"
)

func (match *Match) SimulateMatch(seed int64, isRealTime bool) {
	matchEvent := event.NewMatchEvent(match.ID)

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
	match.MatchRng = matchRng

	// load rules
	match.Rules = constants.DefaultRules()
	match.IsRealTime = isRealTime

	// start game
	gameState := NewGameState()
	match.GameState = &gameState

	// decide map
	mapKeys := []constants.MapName{}
	for k := range constants.Maps {
		mapKeys = append(mapKeys, k)
	}
	slices.Sort(mapKeys)
	gameMap := constants.Maps[mapKeys[matchRng.Intn(len(mapKeys))]]
	match.Map = gameMap
	slog.Info("Map Selection", "map", match.Map.Name)

	// decide starting sides
	attackingTeamIndex := matchRng.Intn(2)
	match.GameState.AttackingTeam = &match.Teams[attackingTeamIndex]
	match.GameState.DefendingTeam = &match.Teams[1-attackingTeamIndex]
	matchEvent.SetAttackingTeam(match.GameState.AttackingTeam.ID)
	matchEvent.SetDefendingTeam(match.GameState.DefendingTeam.ID)

	// simulate agent selection
	match.Teams[0].AgentSelect(constants.Agents, matchRng)
	match.Teams[1].AgentSelect(constants.Agents, matchRng)

	// initialize player game state
	// used to track the player's stats throughout a match
	match.GameState.InitPlayerGameState(&match.Teams[0])
	match.GameState.InitPlayerGameState(&match.Teams[1])

	// start match
	match.GameState.GameRunning = true
	for match.GameState.GameRunning {
		// Swap teams at the end of the half
		if match.GameState.CurrentRound.RoundNumber == 12 {
			match.GameState.CurrentHalf = 2
			match.GameState.AttackingTeam, match.GameState.DefendingTeam = match.GameState.DefendingTeam, match.GameState.AttackingTeam
			matchEvent.SetAttackingTeam(match.GameState.AttackingTeam.ID)
			matchEvent.SetDefendingTeam(match.GameState.DefendingTeam.ID)
		}

		match.StartRound(match.GameState.AttackingTeam, match.GameState.DefendingTeam)
		// Get round winner and loser and up their score here
		roundWinner := match.GameState.CurrentRound.RoundWinner
		if roundWinner == &match.Teams[0] {
			match.GameState.TeamOneScore++
		} else if roundWinner == &match.Teams[1] {
			match.GameState.TeamTwoScore++
		}

		// Check if game goes to overtime
		// TODO: handle overtime

		// Check if a team has won the match
		if match.GameState.TeamOneScore == 13 || match.GameState.TeamTwoScore == 13 {
			match.GameState.GameRunning = false
			slog.Info("Match Ended", "teamOneScore", match.GameState.TeamOneScore, "teamTwoScore", match.GameState.TeamTwoScore)
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

	slog.Info("End Round", "round", gameState.CurrentRound.RoundNumber, "roundWinner", winningTeam.ID, "roundLoser", losingTeam.ID)
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

package match

import (
	"fmt"
	"log/slog"
	"maps"
	"match-sim/internal/constants"
	"math/rand"
	"slices"
	"time"
)

func (match Match) SimulateMatch(seed int64, isRealTime bool) {
	fmt.Println("Simulating match", match.ID)

	fmt.Println("Loading teams")
	for _, team := range match.Teams {
		fmt.Println("Loading players for team", team.ID)
		for _, player := range team.Players {
			fmt.Println("Loading player", player.ID)
		}
	}

	// seed random number generator
	matchSeed := rand.NewSource(seed)
	matchRng := rand.New(matchSeed)

	fmt.Println("Match RNG", matchRng)

	// load rules
	rules := DefaultRules()
	fmt.Println("Loading rules", rules)

	// start game
	gameState := NewGameState()
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
			gameState.AttackingTeam = gameState.DefendingTeam
			gameState.DefendingTeam = gameState.AttackingTeam
		}

		gameState.StartRound(rules, gameState.AttackingTeam, gameState.DefendingTeam, RoundOptions{IsRealTime: true})
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
			break
		}
	}
}

// Starts a round - runs the round logic for each tick until a team wins the round
func (gameState *GameState) StartRound(rules Rules, attackingTeam *Team, defendingTeam *Team, roundOptions RoundOptions) {
	gameState.CurrentRound = NewRoundState(gameState.CurrentRound.RoundNumber + 1)
	slog.Info("Starting round", "round", gameState.CurrentRound.RoundNumber)

	gameState.CurrentRound.AttackingTeam = attackingTeam
	gameState.CurrentRound.DefendingTeam = defendingTeam

	const TickRate = 128
	const TickDuration = time.Second / TickRate

	// Round loop - run this logic for each round until a team reaches 13 rounds won
	totalTicks := rules.RoundDurationSeconds * TickRate

	// Initialize player round state
	gameState.CurrentRound.InitPlayerRoundState(attackingTeam)
	gameState.CurrentRound.InitPlayerRoundState(defendingTeam)

	for tick := 0; tick < totalTicks; tick++ {

		// Kill logic
		// if a player is alive, they have a chance of killing a player on the other team
		// if a player gets a kill, there is a chance that their teammates can get an assist
		// when a player gets a kill, randomly select an enemy to die

		// Attacker Plant Logic
		// after 15 seconds, there is a chance that the spike can be planted

		// Defuse Logic
		// after the spike has been planted, there is a chance that the defending team can defuse
		// Round ends if the defending team defuses the spike [defenders win]

		// Reset total ticks if spike is planted
		if gameState.CurrentRound.SpikePlanted {
			totalTicks = rules.RoundDurationSpikePlantedSeconds * TickRate
		}

		// Round ends if the attacking team successfully detonates the spike [attackers win]
		if gameState.CurrentRound.SpikeDetonated {
			gameState.EndRound(gameState.CurrentRound.AttackingTeam, gameState.CurrentRound.DefendingTeam)
			return
		}

		// Round ends if the attacking team defeats all defenders [attackers win]

		// Round ends if the defending team defeats all attackers before the spike is planted [defenders win]

		// Round ends if the time runs out before the attacking team plants the spike [defenders win]
		if tick == totalTicks-1 {
			gameState.EndRound(gameState.CurrentRound.DefendingTeam, gameState.CurrentRound.AttackingTeam)
			return
		}

		// fmt.Println("Tick", tick)

		// Sleep for real time simulation

		if roundOptions.IsRealTime {
			time.Sleep(TickDuration)
		}
	}
}

// List of event types that can be send to the event component, TODO: Move to event system, TODO: Move to enum
const (
	Event_SpikeDefused     = "spike_defused"
	Event_SpikePlanted     = "spike_planted"
	Event_SpikeDetonated   = "spike_detonated"
	Event_RoundEnd         = "round_end"
	Event_SetAttackingTeam = "set_attacking_team"
	Event_SetDefendingTeam = "set_defending_team"
	Event_Kill             = "kill"
	Event_Death            = "death"
	Event_Assist           = "assist"
)

// Use this to update game state that should be recorded
func (gameState *GameState) UpdateRecordableEvent(event string) {
	switch event {
	case Event_SpikeDefused:
		fmt.Println("Spike defused")
	case Event_SpikePlanted:
		fmt.Println("Spike planted")
	case Event_SpikeDetonated:
		fmt.Println("Spike detonated")
	case Event_RoundEnd:
		fmt.Println("Round ended")
	default:
		fmt.Println("Unknown event", event)
	}

	// TODO: send update to event system
}

// Before starting a game, init the player game state
func (gameState *GameState) InitPlayerGameState(team *Team) {
	for i := range team.Players {
		playerGameState := NewPlayerGameState(&team.Players[i])
		slog.Info("Init Player Game State", "player", playerGameState.LogValue())
		gameState.PlayerGameState = append(gameState.PlayerGameState, &playerGameState)
	}
}

func (roundState *RoundState) InitPlayerRoundState(team *Team) {
	for i := range team.Players {
		playerRoundState := NewPlayerRoundState(&team.Players[i], 0)
		slog.Info("Init Player Round State", "player", playerRoundState.LogValue())
		roundState.PlayerRoundState = append(roundState.PlayerRoundState, &playerRoundState)
	}
}

// Post processing logic for ending a round
func (gameState *GameState) EndRound(winningTeam *Team, losingTeam *Team) {
	gameState.CurrentRound.RoundWinner = winningTeam
	gameState.CurrentRound.RoundLoser = losingTeam
	gameState.Rounds = append(gameState.Rounds, gameState.CurrentRound)
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

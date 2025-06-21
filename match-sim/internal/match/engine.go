package match

import (
	"fmt"
	"maps"
	"match-sim/internal/constants"
	"math/rand"
	"slices"
	"time"
)

func (match Match) SimulateMatch(seed int64) {
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

	// decide starting sides
	attackingTeamIndex := matchRng.Intn(2)
	gameState.AttackingTeam = &match.Teams[attackingTeamIndex]
	gameState.DefendingTeam = &match.Teams[1-attackingTeamIndex]
	// TODO: Update game state - TODO: Move to event system

	fmt.Println("Attacking team", gameState.AttackingTeam.ID)
	fmt.Println("Defending team", gameState.DefendingTeam.ID)

	match.Teams[0].AgentSelect(constants.Agents, matchRng)
	match.Teams[1].AgentSelect(constants.Agents, matchRng)

	fmt.Println("Team 1 agents", match.Teams[0].Players)
	fmt.Println("Team 2 agents", match.Teams[1].Players)

	// start match
	// gameState.GameRunning = true
	// for gameState.GameRunning {
	// 	gameState.StartRound(rules)

	// 	// Check if a team has won the match
	// 	if gameState.TeamOneWon || gameState.TeamTwoWon {
	// 		gameState.GameRunning = false
	// 		break
	// 	}
	// }
}

// Starts a round - runs the round logic for each tick until a team wins the round
func (gameState *GameState) StartRound(rules Rules) {

	// Before starting the round, check if the half is over
	if gameState.CurrentRound == 12 {
		gameState.CurrentHalf = 2

		// Swap the attacking and defending teams
		attackingTeam := gameState.AttackingTeam
		defendingTeam := gameState.DefendingTeam
		gameState.AttackingTeam = defendingTeam
		gameState.DefendingTeam = attackingTeam
	}

	gameState.CurrentRound++
	const TickRate = 128
	const TickDuration = time.Second / TickRate

	// Round loop - run this logic for each round until a team reaches 13 rounds won
	totalTicks := rules.RoundDurationSeconds * TickRate
	for tick := 0; tick < totalTicks; tick++ {

		// Reset total ticks if spike is planted
		if gameState.SpikePlanted {
			totalTicks = rules.RoundDurationSpikePlantedSeconds * TickRate
		}

		// Round ends if the attacking team successfully detonates the spike [attackers win]

		// Round ends if the attacking team defeats all defenders [attackers win]

		// Round ends if the defending team defuses the spike [defenders win]

		// Round ends if the defending team defeats all attackers before the spike is planted [defenders win]

		// Round ends if the time runs out before the attacking team plants the spike [defenders win]

		fmt.Println("Tick", tick)

		// Sleep for real time simulation
		time.Sleep(TickDuration)
	}
}

// List of event types that can be send to the event component, TODO: Move to event system
const (
	Event_SpikeDefused     = "spike_defused"
	Event_SpikePlanted     = "spike_planted"
	Event_SpikeDetonated   = "spike_detonated"
	Event_RoundEnd         = "round_end"
	Event_SetAttackingTeam = "set_attacking_team"
	Event_SetDefendingTeam = "set_defending_team"
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

func (team Team) AgentSelect(agents map[constants.AgentName]constants.Agent, rng *rand.Rand) {
	availableAgents := make(map[constants.AgentName]constants.Agent)
	maps.Copy(availableAgents, agents)

	for i := range team.Players {
		possibleAgents := filterAgents(availableAgents, team.Players[i].Role)
		agent := possibleAgents[rng.Intn(len(possibleAgents))]
		// TODO: update event system

		team.Players[i].Agent = agent
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

package match

import (
	"fmt"
	"match-sim/internal/constants"
	"math/rand"
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

	// decide sides
	attackingTeamIndex := matchRng.Intn(2)
	attackingTeam := match.Teams[attackingTeamIndex]
	defendingTeam := match.Teams[1-attackingTeamIndex]
	// TODO: Update game state - TODO: Move to event system

	fmt.Println("Attacking team", attackingTeam.ID)
	fmt.Println("Defending team", defendingTeam.ID)

	// team 1 agent selection
	availableAgents := constants.Agents
	for i := range match.Teams[0].Players {
		possibleAgents := filterAgents(availableAgents, match.Teams[0].Players[i].Role)
		agent := possibleAgents[matchRng.Intn(len(possibleAgents))]
		// Update game state - TODO: Move to event system
		match.Teams[0].Players[i].Agent = agent
		availableAgents = removeAgent(availableAgents, agent)
	}

	// team 2 agent selection
	availableAgents = constants.Agents
	for i := range match.Teams[1].Players {
		possibleAgents := filterAgents(availableAgents, match.Teams[1].Players[i].Role)
		agent := possibleAgents[matchRng.Intn(len(possibleAgents))]

		availableAgents = removeAgent(availableAgents, agent)
		// Update game state - TODO: Move to event system
		match.Teams[1].Players[i].Agent = agent
	}

	fmt.Println("Team 1 agents", match.Teams[0].Players)
	fmt.Println("Team 2 agents", match.Teams[1].Players)

	// start match

	// round engine
}

func filterAgents(agents []constants.Agent, role constants.AgentRole) []constants.Agent {
	filteredAgents := []constants.Agent{}
	for _, agent := range agents {
		if agent.Role == role {
			filteredAgents = append(filteredAgents, agent)
		}
	}
	return filteredAgents
}

func removeAgent(agents []constants.Agent, agent constants.Agent) []constants.Agent {
	filteredAgents := []constants.Agent{}
	for _, a := range agents {
		if a.ID != agent.ID {
			filteredAgents = append(filteredAgents, a)
		}
	}
	return filteredAgents
}

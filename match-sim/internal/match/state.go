package match

import "match-sim/internal/constants"

type Player struct {
	ID    string
	Role  constants.AgentRole
	Agent constants.Agent
}

type Team struct {
	ID      string
	Players []Player
}

type Match struct {
	ID    string
	Teams []Team
}

type State struct {
	Matches []Match
}

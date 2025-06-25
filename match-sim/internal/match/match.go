package match

import "match-sim/internal/constants"

// Match type that we receive from the tournament orchestrator
type Match struct {
	ID    string
	Teams []Team
}

type Team struct {
	ID      string
	Name    string
	Players []Player
}

type Player struct {
	ID     string
	Name   string
	TeamID string
	Role   constants.AgentRole // TODO: enum
}

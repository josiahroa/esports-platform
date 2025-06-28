package match

import "match-sim/internal/domain/valorant"

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
	Role   valorant.AgentRole
}

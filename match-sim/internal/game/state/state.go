package state

import "match-sim/internal/match"

// both the client and the server will be aware of the game state for the simulation
// the simulator reads the game state for simulation but never writes to it

type Player struct {
	Player       *match.Player
	Agent        string // TODO: enum
	Role         string // TODO: enum
	Health       uint8
	Armor        uint8
	Credits      uint16
	Eliminations uint8
	Assists      uint8
	Deaths       uint8
}

// TODO: move to state
type Round struct {
	RoundNumber uint8
	RoundType   string // TODO: enum
	RoundTime   uint16
	RoundWinner *match.Team
}

type Game struct {
	AttackingTeam     *match.Team // current attacking team
	DefendingTeam     *match.Team // current defending team
	Match             *match.Match
	PlayerState       map[string]Player
	Map               string // TODO: enum
	CurrentRoundIndex uint8
	CurrentHalf       bool // 0(false) = first half, 1(true) = second half
	Rounds            []*Round
	TeamOneScore      uint8
	TeamTwoScore      uint8
}

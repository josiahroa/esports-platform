package state

import (
	"match-sim/internal/domain/valorant"
	"match-sim/internal/game"
	"match-sim/internal/match"
)

// both the client and the server will be aware of the game state for the simulation
// the simulator reads the game state for simulation but never writes to it

type Player struct {
	Player       *match.Player
	Agent        valorant.AgentName
	Role         valorant.AgentRole
	Health       uint8
	Armor        uint8
	Credits      uint16
	Eliminations uint8
	Assists      uint8
	Deaths       uint8
}

type Game struct {
	AttackingTeam     *match.Team // current attacking team
	DefendingTeam     *match.Team // current defending team
	Match             *match.Match
	PlayerState       map[string]Player
	Map               valorant.MapName
	CurrentRoundIndex uint8
	CurrentHalf       bool // 0(false) = first half, 1(true) = second half
	Rounds            []*game.Round
	TeamOneScore      uint8
	TeamTwoScore      uint8
}

func NewGame(match *match.Match) *Game {
	game := &Game{
		AttackingTeam:     &match.Teams[0],
		DefendingTeam:     &match.Teams[1],
		Match:             match,
		PlayerState:       map[string]Player{},
		Map:               valorant.MapName_SUNSET,
		CurrentRoundIndex: 0,
		CurrentHalf:       false,
		Rounds:            []*game.Round{},
		TeamOneScore:      0,
		TeamTwoScore:      0,
	}

	// retrieve players from match and load into player state
	for _, team := range match.Teams {
		for _, player := range team.Players {
			game.PlayerState[player.ID] = Player{
				Player: &player,
			}
		}
	}

	return game
}

func (g *Game) AddRound(round *game.Round) {
	g.Rounds = append(g.Rounds, round)
}

func (g *Game) IncrementRound() uint8 {
	g.CurrentRoundIndex++
	return g.CurrentRoundIndex
}

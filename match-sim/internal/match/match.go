package match

import (
	"log/slog"
	"match-sim/internal/constants"
	"math/rand"
)

type Match struct {
	ID         string
	Teams      []Team
	GameState  *GameState
	Winner     *Team
	IsRealTime bool
	Rules      Rules
	MatchRng   *rand.Rand
	Map        constants.Map
}

type Player struct {
	ID    string
	Name  string
	Role  constants.AgentRole
	Agent constants.Agent
}

type Team struct {
	ID      string
	Name    string
	Players []Player
}

type PlayerKill struct {
	Killer *Player
	Victim *Player
	Weapon constants.Weapon
}

func (p *PlayerKill) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("killer_id", p.Killer.ID),
		slog.String("victim_id", p.Victim.ID),
		// slog.String("weapon", p.Weapon.Name), //TODO: add weapon
	)
}

type PlayerGameState struct {
	Player  *Player
	Kills   uint8
	Deaths  uint8
	Assists uint8
	Score   int
}

func NewPlayerGameState(player *Player) PlayerGameState {
	return PlayerGameState{
		Player:  player,
		Kills:   0,
		Deaths:  0,
		Assists: 0,
		Score:   0,
	}
}

func (p *PlayerGameState) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("player_id", p.Player.ID),
		slog.Uint64("kills", uint64(p.Kills)),
		slog.Uint64("deaths", uint64(p.Deaths)),
		slog.Uint64("assists", uint64(p.Assists)),
		slog.Int("score", p.Score),
	)
}

type GameState struct {
	AttackingTeam   *Team
	CurrentHalf     uint8
	CurrentRound    *RoundState
	DefendingTeam   *Team
	GameRunning     bool
	Rounds          []*RoundState
	PlayerGameState []*PlayerGameState
	TeamOneScore    uint8
	TeamTwoScore    uint8
}

func NewGameState() GameState {
	return GameState{
		CurrentHalf: 1,
		CurrentRound: &RoundState{
			RoundNumber: 0, // Start at 0 to indicate no round has been started yet
		},
		GameRunning:     false,
		Rounds:          []*RoundState{},
		PlayerGameState: []*PlayerGameState{},
		TeamOneScore:    0,
		TeamTwoScore:    0,
	}
}

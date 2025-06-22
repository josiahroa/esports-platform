package match

import (
	"log/slog"
	"match-sim/internal/constants"
)

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
	ID        string
	Teams     []Team
	GameState GameState
	Winner    *Team
}

func (m *Match) GetID() string {
	return m.ID
}

type State struct {
	Matches []Match
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

func (p *PlayerGameState) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("player_id", p.Player.ID),
		slog.Uint64("kills", uint64(p.Kills)),
		slog.Uint64("deaths", uint64(p.Deaths)),
		slog.Uint64("assists", uint64(p.Assists)),
		slog.Int("score", p.Score),
	)
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

type GameState struct {
	AttackingTeam    *Team
	CurrentHalf      uint8
	CurrentRound     RoundState
	DefendingTeam    *Team
	GameRunning      bool
	Rounds           []RoundState
	PlayerGameState  []*PlayerGameState
	TeamOne          *Team
	TeamOneRoundsWon uint8
	TeamOneScore     uint8
	TeamTwo          *Team
	TeamTwoRoundsWon uint8
	TeamTwoScore     uint8
	Map              constants.Map
	IsRealTime       bool
	Rules            Rules
}

func NewGameState(rules Rules, isRealTime bool) GameState {
	return GameState{
		CurrentHalf: 1,
		CurrentRound: RoundState{
			RoundNumber: 0, // Start at 0 to indicate no round has been started yet
		},
		GameRunning:      false,
		Rounds:           []RoundState{},
		PlayerGameState:  []*PlayerGameState{},
		TeamOne:          &Team{},
		TeamOneRoundsWon: 0,
		TeamOneScore:     0,
		TeamTwo:          &Team{},
		TeamTwoRoundsWon: 0,
		TeamTwoScore:     0,
		IsRealTime:       isRealTime,
		Rules:            rules,
	}
}

type PlayerRoundState struct {
	Team            *Team
	Player          *Player
	StartingCredits uint
	EndCredits      uint
	Alive           bool
}

func NewPlayerRoundState(team *Team, player *Player, startingCredits uint) PlayerRoundState {
	return PlayerRoundState{
		Team:            team,
		Player:          player,
		StartingCredits: startingCredits,
		EndCredits:      0,
		Alive:           true,
	}
}

type RoundOptions struct {
	IsRealTime bool
	IsOvertime bool
}

type RoundState struct {
	AttackingTeam        *Team
	DefendingTeam        *Team
	Half                 uint8
	RoundDuration        int
	RoundLoser           *Team
	RoundNumber          uint8
	RoundWinner          *Team
	PlayerRoundState     []*PlayerRoundState
	SpikePlanted         bool
	SpikePlantedLocation constants.PlantSite
	SpikePlantedBy       *Player
	SpikeDefused         bool
	SpikeDefusedBy       *Player
	SpikeDetonated       bool
}

func (p *PlayerRoundState) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("player_id", p.Player.ID),
		slog.Uint64("starting_credits", uint64(p.StartingCredits)),
		slog.Uint64("end_credits", uint64(p.EndCredits)),
		slog.Bool("alive", bool(p.Alive)),
	)
}

func NewRoundState(roundNumber uint8) RoundState {
	return RoundState{
		AttackingTeam:        nil,
		DefendingTeam:        nil,
		Half:                 1,
		RoundDuration:        100,
		RoundLoser:           nil,
		RoundNumber:          roundNumber,
		RoundWinner:          nil,
		PlayerRoundState:     []*PlayerRoundState{},
		SpikePlanted:         false,
		SpikePlantedLocation: constants.PlantSiteA,
		SpikeDetonated:       false,
	}
}

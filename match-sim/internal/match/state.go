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
	ID        string
	Teams     []Team
	GameState GameState
	Winner    *Team
}

type State struct {
	Matches []Match
}

type PlayerKill struct {
	Killer *Player
	Victim *Player
	Weapon constants.Weapon
}

type PlayerState struct {
	ID              string
	Player          *Player
	StartingCredits int
	EndCredits      int
	Kills           uint8
	Deaths          uint8
	Assists         uint8
	Score           int
}

type GameState struct {
	AttackingTeam    *Team
	CurrentHalf      uint8
	CurrentRound     RoundState
	DefendingTeam    *Team
	GameRunning      bool
	Rounds           []RoundState
	TeamOne          *Team
	TeamOneRoundsWon uint8
	TeamOneScore     uint8
	TeamTwo          *Team
	TeamTwoRoundsWon uint8
	TeamTwoScore     uint8
	Map              constants.Map
}

func NewGameState() GameState {
	return GameState{
		CurrentHalf: 1,
		CurrentRound: RoundState{
			RoundNumber: 0, // Start at 0 to indicate no round has been started yet
		},
		GameRunning:      false,
		Rounds:           []RoundState{},
		TeamOne:          &Team{},
		TeamOneRoundsWon: 0,
		TeamOneScore:     0,
		TeamTwo:          &Team{},
		TeamTwoRoundsWon: 0,
		TeamTwoScore:     0,
	}
}

type RoundState struct {
	AttackingTeam        *Team
	DefendingTeam        *Team
	Half                 uint8
	RoundDuration        int
	RoundLoser           *Team
	RoundNumber          uint8
	RoundWinner          *Team
	Players              []*PlayerState
	SpikePlanted         bool
	SpikePlantedLocation constants.PlantSite
	SpikeDetonated       bool
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
		Players:              []*PlayerState{},
		SpikePlanted:         false,
		SpikePlantedLocation: constants.PlantSiteA,
		SpikeDetonated:       false,
	}
}

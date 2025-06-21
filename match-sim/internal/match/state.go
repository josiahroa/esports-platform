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
	Kills           int
	Deaths          int
	Assists         int
	Score           int
}

type RoundState struct {
	AttackingTeam *Team
	DefendingTeam *Team
	Half          int
	RoundDuration int
	RoundLoser    *Team
	RoundNumber   int
	RoundWinner   *Team
	Players       []PlayerState
}

type GameState struct {
	AttackingTeam    *Team
	CurrentHalf      int
	CurrentRound     int
	DefendingTeam    *Team
	GameRunning      bool
	Rounds           []RoundState
	SpikePlanted     bool
	TeamOneRoundsWon int
	TeamOneScore     int
	TeamOneWon       bool
	TeamTwoRoundsWon int
	TeamTwoScore     int
	TeamTwoWon       bool
}

func NewGameState() GameState {
	return GameState{
		AttackingTeam:    nil,
		CurrentHalf:      1,
		CurrentRound:     1,
		DefendingTeam:    nil,
		GameRunning:      false,
		Rounds:           []RoundState{},
		SpikePlanted:     false,
		TeamOneRoundsWon: 0,
		TeamOneScore:     0,
		TeamOneWon:       false,
		TeamTwoRoundsWon: 0,
		TeamTwoScore:     0,
		TeamTwoWon:       false,
	}
}

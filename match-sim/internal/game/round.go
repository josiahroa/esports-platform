package game

import "match-sim/internal/match"

type Round struct {
	RoundNumber uint8
	RoundTime   uint16
	RoundWinner *match.Team
}

func CreateRound(roundNumber uint8) *Round {
	return &Round{
		RoundNumber: roundNumber,
		RoundTime:   0,
		RoundWinner: nil,
	}
}

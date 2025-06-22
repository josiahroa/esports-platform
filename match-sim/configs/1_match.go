package configs

import (
	"match-sim/internal/constants"
	"match-sim/internal/match"
)

var Match1 = match.Match{
	ID: "1",
	Teams: []match.Team{
		{ID: "11232323", Name: "Gamers", Players: []match.Player{
			{ID: "1", Name: "Player 1", Role: constants.Controller},
			{ID: "2", Name: "Player 2", Role: constants.Initiator},
			{ID: "3", Name: "Player 3", Role: constants.Duelist},
			{ID: "4", Name: "Player 4", Role: constants.Sentinel},
			{ID: "5", Name: "Player 5", Role: constants.Controller},
		}},
		{ID: "23451324", Name: "Pro Players", Players: []match.Player{
			{ID: "6", Name: "Player 6", Role: constants.Controller},
			{ID: "7", Name: "Player 7", Role: constants.Initiator},
			{ID: "8", Name: "Player 8", Role: constants.Duelist},
			{ID: "9", Name: "Player 9", Role: constants.Sentinel},
			{ID: "10", Name: "Player 10", Role: constants.Sentinel},
		}},
	},
}

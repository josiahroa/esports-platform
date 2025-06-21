package configs

import (
	"match-sim/internal/constants"
	"match-sim/internal/match"
)

var Match1 = match.Match{
	ID: "1",
	Teams: []match.Team{
		{ID: "11232323", Players: []match.Player{
			{ID: "1", Role: constants.Controller},
			{ID: "2", Role: constants.Initiator},
			{ID: "3", Role: constants.Duelist},
			{ID: "4", Role: constants.Sentinel},
			{ID: "5", Role: constants.Controller},
		}},
		{ID: "23451324", Players: []match.Player{
			{ID: "6", Role: constants.Controller},
			{ID: "7", Role: constants.Initiator},
			{ID: "8", Role: constants.Duelist},
			{ID: "9", Role: constants.Sentinel},
			{ID: "10", Role: constants.Sentinel},
		}},
	},
}

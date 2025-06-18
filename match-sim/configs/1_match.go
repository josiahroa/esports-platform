package configs

import "match-sim/internal/match"

var Match1 = match.Match{
	ID: "1",
	Teams: []match.Team{
		{ID: "11232323", Players: []match.Player{
			{ID: "1", Role: "Controller"},
			{ID: "2", Role: "Initiator"},
			{ID: "3", Role: "Duelist"},
			{ID: "4", Role: "Sentinel"},
			{ID: "5", Role: "Controller"},
		}},
		{ID: "23451324", Players: []match.Player{
			{ID: "6", Role: "Controller"},
			{ID: "7", Role: "Initiator"},
			{ID: "8", Role: "Duelist"},
			{ID: "9", Role: "Sentinel"},
			{ID: "10", Role: "Sentinel"},
		}},
	},
}

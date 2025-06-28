package main

import (
	"match-sim/internal/domain/valorant"
	"match-sim/internal/game/client"
	"match-sim/internal/game/server"
	"match-sim/internal/match"
	"math/rand"
)

var Match1 = match.Match{
	ID: "1",
	Teams: []match.Team{
		{ID: "11232323", Name: "Gamers", Players: []match.Player{
			{ID: "1", Name: "Player 1", Role: valorant.AgentRole_CONTROLLER},
			{ID: "2", Name: "Player 2", Role: valorant.AgentRole_INITIATOR},
			{ID: "3", Name: "Player 3", Role: valorant.AgentRole_DUELIST},
			{ID: "4", Name: "Player 4", Role: valorant.AgentRole_SENTINEL},
			{ID: "5", Name: "Player 5", Role: valorant.AgentRole_CONTROLLER},
		}},
		{ID: "23451324", Name: "Pro Players", Players: []match.Player{
			{ID: "6", Name: "Player 6", Role: valorant.AgentRole_CONTROLLER},
			{ID: "7", Name: "Player 7", Role: valorant.AgentRole_INITIATOR},
			{ID: "8", Name: "Player 8", Role: valorant.AgentRole_DUELIST},
			{ID: "9", Name: "Player 9", Role: valorant.AgentRole_SENTINEL},
			{ID: "10", Name: "Player 10", Role: valorant.AgentRole_SENTINEL},
		}},
	},
}

func main() {
	seed := int64(12341234)
	rng := rand.New(rand.NewSource(seed))

	// Load match from configs
	match := Match1

	client.SimulateGame(&match, &server.Config{
		IsRealTime: false,
		Seed:       rng,
	})
}

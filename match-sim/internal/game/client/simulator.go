package client

import (
	"log/slog"
	"match-sim/internal/game"
	"match-sim/internal/game/server"
	"match-sim/internal/match"
)

// Simulates a game and sends events to the gameserver
func SimulateMatch(match *match.Match, config *server.Config) {
	// start the game server, we call its creation here because
	// we are mocking the connection of the client to the server
	// in this simulator.
	server := server.StartServer(match, config)

	slog.Info("Match started", "gameState", server.GameState)

	// Map selection
	match.Map = server.SelectMap()
	slog.Info("Map selected", "map", match.Map)

	// TODO: Select agents
	// simulate each player selecting an agent from the server's agent pool
	// when a player selects an agent, the server will update the agent pool
	// and the client will receive the updated agent pool
	for _, player := range match.Teams[0].Players {
		agentPool := server.GetAgentPool_Role(player.Role)
		selected := false
		for !selected {
			selected = server.TeamOne_AgentSelect(agentPool[int32(config.Rng.Intn(len(agentPool)))].Agent)
		}
	}

	// add the first round to the game state
	server.GameState.AddRound(game.CreateRound(0))

	// MVP: simulate basic game logic like win conditions and eliminations
	gameRunning := true
	for gameRunning {
		// check for win conditions
		if server.GameState.TeamOneScore == 13 || server.GameState.TeamTwoScore == 13 {
			gameRunning = false
		}

		// check for half change
		currentRound := server.GameState.Rounds[server.GameState.CurrentRoundIndex]
		if currentRound.RoundNumber == 12 {
			server.GameState.CurrentHalf = !server.GameState.CurrentHalf
		}

		// simulate round
		SimulateRound(server)

		// increment the round
		currentRoundIndex := server.GameState.IncrementRound()
		server.GameState.AddRound(game.CreateRound(currentRoundIndex))
	}
}

func SimulateRound(server *server.Server) {
	// TOOD: Simulate weapon purchases

	// TODO: Simulate player actions

	// TODO: Simulate round end
}

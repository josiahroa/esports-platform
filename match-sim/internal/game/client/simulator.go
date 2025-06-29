package client

import (
	"log/slog"
	"match-sim/internal/game"
	"match-sim/internal/game/server"
	"match-sim/internal/match"
	// "math/rand"
)

// Simulates a game and sends events to the gameserver
func SimulateMatch(match *match.Match, config *server.Config) {
	// start the game server, we call its creation here because
	// we are mocking the connection of the client to the server
	// in this simulator.
	server := server.StartServer(match)

	slog.Info("Match started", "gameState", server.GameState)

	// TODO: Select map

	// TODO: Select agents

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

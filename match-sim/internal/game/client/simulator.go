package client

import (
	"log/slog"
	"match-sim/internal/game/server"
	"match-sim/internal/match"
)

// Simulates a game and sends events to the gameserver

func SimulateGame(match *match.Match, config *server.Config) {

	// create a random number generator

	// start the game server
	gameState := server.StartServer(match, config)

	slog.Info("Game started", "gameState", gameState)

	// MVP: simulate basic game logic like win conditions and eliminations
}

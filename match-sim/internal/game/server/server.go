package server

// Receives events from the gameclient, maintains game state, and sends asyncronously sends events to kinesis
// since this is a simulator, we will not be responding to the game client.
// the simulator(game client) will be aware of the game state in this scenario.

// expose a public function that will receive events from the gameclient
// this will mimic a UDP connection between the gameclient and the game server

import (
	"log/slog"
	"match-sim/internal/game/state"
	"match-sim/internal/match"
	"math/rand"
)

// TODO: move to state

type Config struct {
	IsRealTime bool
	Seed       *rand.Rand
}

type Server struct {
	GameState *state.Game
}

func StartServer(match *match.Match) *Server {
	gameState := state.NewGame(match)

	return &Server{
		GameState: gameState,
	}
}

func (s *Server) SelectMap() {
	// randomly selects a map for the match
}

func (s *Server) SelectAgent(playerID string, agentID int) {
	// receive agent selection per player from game client

	slog.Info("Agent Selection", "playerID", playerID, "agentID", agentID)
}

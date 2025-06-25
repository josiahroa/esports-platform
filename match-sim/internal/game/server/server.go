package server

// Receives events from the gameclient, maintains game state, and sends asyncronously sends events to kinesis
// since this is a simulator, we will not be responding to the game client.
// the simulator(game client) will be aware of the game state in this scenario.

// expose a public function that will receive events from the gameclient
// this will mimic a UDP connection between the gameclient and the game server

import (
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
	Config    *Config
}

func StartServer(match *match.Match, config *Config) *Server {
	gameState := &state.Game{
		Match:             match,
		PlayerState:       map[string]state.Player{},
		Map:               "TEST",
		CurrentRoundIndex: 0,
		CurrentHalf:       false,
		Rounds:            []*state.Round{},
		TeamOneScore:      0,
		TeamTwoScore:      0,
	}

	// load players from match
	for _, team := range match.Teams {
		for _, player := range team.Players {
			gameState.PlayerState[player.ID] = state.Player{
				Player: &player,
			}
		}
	}

	return &Server{
		GameState: gameState,
		Config:    config,
	}
}

func (s *Server) SelectMap() {
	// randomly selects a map for the match
}

func (s *Server) SelectAgent(playerID string, agent string) {
	// receive agent selection per player from game client
}

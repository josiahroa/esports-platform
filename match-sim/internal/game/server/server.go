package server

// Receives events from the gameclient, maintains game state, and sends asyncronously sends events to kinesis
// since this is a simulator, we will not be responding to the game client.
// the simulator(game client) will be aware of the game state in this scenario.

// expose a public function that will receive events from the gameclient
// this will mimic a UDP connection between the gameclient and the game server

import (
	"fmt"
	"match-sim/internal/domain/valorant"
	"match-sim/internal/game/state"
	"match-sim/internal/match"
	"math/rand"
)

// A struct specific to the agent pool so that we can track which agents are selected
type AgentPool_Agent struct {
	Agent    *valorant.Agent
	Selected bool
}

func (a *AgentPool_Agent) String() string {
	return fmt.Sprintf("Agent: %s, Selected: %t", a.Agent.Name, a.Selected)
}

type Config struct {
	IsRealTime bool
	Rng        *rand.Rand
}

type Server struct {
	GameState         *state.Game
	Config            *Config
	TeamOne_AgentPool map[int32]*AgentPool_Agent
	TeamTwo_AgentPool map[int32]*AgentPool_Agent
}

func StartServer(match *match.Match, config *Config) *Server {
	gameState := state.NewGame(match)

	teamOneAgentPool := newAgentPool()
	teamTwoAgentPool := newAgentPool()

	return &Server{
		GameState:         gameState,
		Config:            config,
		TeamOne_AgentPool: teamOneAgentPool,
		TeamTwo_AgentPool: teamTwoAgentPool,
	}
}

func (s *Server) SelectMap() (matchMap *valorant.Map) {
	matchMap = valorant.CreateMaps()[s.Config.Rng.Intn(len(valorant.CreateMaps()))]
	return matchMap
}

func (s *Server) TeamOne_AgentSelect(agent *valorant.Agent) bool {
	// check if selected agent is already selected
	if s.TeamOne_AgentPool[agent.Id].Selected {
		return false
	}

	s.TeamOne_AgentPool[agent.Id].Selected = true
	return true
}

func newAgentPool_Agent(agent *valorant.Agent) (agentPool_Agent *AgentPool_Agent) {
	agentPool_Agent = &AgentPool_Agent{
		Agent:    agent,
		Selected: false,
	}
	return agentPool_Agent
}

func newAgentPool() (agentPool map[int32]*AgentPool_Agent) {
	agentPool = make(map[int32]*AgentPool_Agent)
	for _, agent := range valorant.CreateAgents() {
		agentPool[agent.Id] = newAgentPool_Agent(agent)
	}
	return agentPool
}

// Provides an agent pool for a specified role
func (s *Server) GetAgentPool_Role(role valorant.AgentRole) (agentPool map[int32]*AgentPool_Agent) {
	agentPool = make(map[int32]*AgentPool_Agent)
	for _, agent := range s.TeamOne_AgentPool {
		if agent.Agent.Role == role {
			agentPool[agent.Agent.Id] = agent
		}
	}
	return agentPool
}

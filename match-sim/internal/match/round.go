package match

import (
	"math/rand"
	"time"
)

type GameTickSimulator struct {
	TotalTicks   uint
	GameState    *GameState
	Rng          *rand.Rand
	TickRate     uint
	TickDuration time.Duration
}

func NewGameTickSimulator(
	totalTicks uint,
	gameState *GameState,
	rng *rand.Rand,
	tickRate uint,
	tickDuration time.Duration,
) *GameTickSimulator {
	return &GameTickSimulator{
		TotalTicks:   totalTicks,
		GameState:    gameState,
		Rng:          rng,
		TickRate:     tickRate,
		TickDuration: tickDuration,
	}
}

func (t *GameTickSimulator) SimulateGameTick(tick *uint) bool {
	const killProbability = 0.0004 // Results in an average of ~5 kills per 100s round
	currentRound := &t.GameState.CurrentRound

	if t.Rng.Float64() < killProbability {
		currentRound.SimulatePlayerKill(t.Rng)
	}

	// Attacker Plant Logic
	// after 15 seconds, there is a chance that the spike can be planted
	if *tick >= 15*t.TickRate && t.Rng.Float64() < 0.0001 {
		attackingTeamPlayers := currentRound.GetAlivePlayers(currentRound.AttackingTeam)
		if len(attackingTeamPlayers) > 0 {
			currentRound.SpikePlanted = true
			currentRound.SpikePlantedBy = attackingTeamPlayers[t.Rng.Intn(len(attackingTeamPlayers))].Player
		}
	}

	// Defuse Logic
	// after the spike has been planted, there is a chance that the defending team can defuse
	// Round ends if the defending team defuses the spike [defenders win]
	if currentRound.SpikePlanted && t.Rng.Float64() < 0.0001 {
		defendingTeamPlayers := currentRound.GetAlivePlayers(currentRound.DefendingTeam)
		if len(defendingTeamPlayers) > 0 {
			currentRound.SpikeDefused = true
			currentRound.SpikeDefusedBy = defendingTeamPlayers[t.Rng.Intn(len(defendingTeamPlayers))].Player

			t.GameState.EndRound(currentRound.DefendingTeam, currentRound.AttackingTeam)
			return true
		}
	}

	// Reset total ticks if spike is planted
	if currentRound.SpikePlanted {
		*tick = 0
		t.TotalTicks = uint(t.GameState.Rules.RoundDurationSpikePlantedSeconds) * t.TickRate
	}

	// Round ends if the attacking team successfully detonates the spike [attackers win]
	if currentRound.SpikeDetonated {
		t.GameState.EndRound(currentRound.AttackingTeam, currentRound.DefendingTeam)
		return true
	}

	// Round ends if the attacking team defeats all defenders [attackers win]
	if len(currentRound.GetAlivePlayers(currentRound.DefendingTeam)) == 0 {
		t.GameState.EndRound(currentRound.AttackingTeam, currentRound.DefendingTeam)
		return true
	}

	// Round ends if the defending team defeats all attackers before the spike is planted [defenders win]
	if len(currentRound.GetAlivePlayers(currentRound.AttackingTeam)) == 0 && !currentRound.SpikePlanted {
		t.GameState.EndRound(currentRound.DefendingTeam, currentRound.AttackingTeam)
		return true
	}

	// Round ends if the time runs out before the attacking team plants the spike [defenders win]
	if *tick == t.TotalTicks-1 {
		t.GameState.EndRound(currentRound.DefendingTeam, currentRound.AttackingTeam)
		return true
	}

	// Sleep for real time simulation
	if t.GameState.IsRealTime {
		time.Sleep(t.TickDuration)
	}
	return false
}

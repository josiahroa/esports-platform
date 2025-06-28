package match

// import (
// 	"match-sim/internal/constants"
// 	"math/rand"
// 	"time"

// 	"log/slog"
// )

// type PlayerRoundState struct {
// 	Team            *Team
// 	Player          *Player
// 	StartingCredits uint
// 	EndCredits      uint
// 	Alive           bool
// }

// func NewPlayerRoundState(team *Team, player *Player, startingCredits uint) PlayerRoundState {
// 	return PlayerRoundState{
// 		Team:            team,
// 		Player:          player,
// 		StartingCredits: startingCredits,
// 		EndCredits:      0,
// 		Alive:           true,
// 	}
// }

// type RoundOptions struct {
// 	IsRealTime bool
// 	IsOvertime bool
// }

// type RoundState struct {
// 	AttackingTeam        *Team
// 	DefendingTeam        *Team
// 	Half                 uint8
// 	RoundDuration        int
// 	RoundLoser           *Team
// 	RoundNumber          uint8
// 	RoundWinner          *Team
// 	PlayerRoundState     []*PlayerRoundState
// 	SpikePlanted         bool
// 	SpikePlantedLocation constants.PlantSite
// 	SpikePlantedBy       *Player
// 	SpikeDefused         bool
// 	SpikeDefusedBy       *Player
// 	SpikeDetonated       bool
// }

// func (p *PlayerRoundState) LogValue() slog.Value {
// 	return slog.GroupValue(
// 		slog.String("player_id", p.Player.ID),
// 		slog.Uint64("starting_credits", uint64(p.StartingCredits)),
// 		slog.Uint64("end_credits", uint64(p.EndCredits)),
// 		slog.Bool("alive", bool(p.Alive)),
// 	)
// }

// func NewRoundState(roundNumber uint8) RoundState {
// 	return RoundState{
// 		AttackingTeam:        nil,
// 		DefendingTeam:        nil,
// 		Half:                 1,
// 		RoundDuration:        100,
// 		RoundLoser:           nil,
// 		RoundNumber:          roundNumber,
// 		RoundWinner:          nil,
// 		PlayerRoundState:     []*PlayerRoundState{},
// 		SpikePlanted:         false,
// 		SpikePlantedLocation: constants.PlantSiteA,
// 		SpikeDetonated:       false,
// 	}
// }

// type GameTickSimulator struct {
// 	TotalTicks   uint
// 	Match        *Match
// 	Rng          *rand.Rand
// 	TickRate     uint
// 	TickDuration time.Duration
// }

// func NewGameTickSimulator(
// 	totalTicks uint,
// 	match *Match,
// 	tickRate uint,
// 	tickDuration time.Duration,
// ) *GameTickSimulator {
// 	return &GameTickSimulator{
// 		TotalTicks:   totalTicks,
// 		Match:        match,
// 		TickRate:     tickRate,
// 		TickDuration: tickDuration,
// 	}
// }

// // Starts a round - runs the round logic for each tick until a team wins the round
// func (match *Match) StartRound(attackingTeam *Team, defendingTeam *Team) {
// 	currentRound := NewRoundState(match.GameState.CurrentRound.RoundNumber + 1)
// 	match.GameState.CurrentRound = &currentRound
// 	match.GameState.CurrentRound.AttackingTeam = attackingTeam
// 	match.GameState.CurrentRound.DefendingTeam = defendingTeam

// 	const tickRate uint = 128
// 	tickDuration := time.Second / time.Duration(tickRate)

// 	// Round loop - run this logic for each round until a team reaches 13 rounds won
// 	totalTicks := uint(match.Rules.RoundDurationSeconds) * tickRate

// 	// Initialize player round state
// 	match.GameState.CurrentRound.InitPlayerRoundState(attackingTeam)
// 	match.GameState.CurrentRound.InitPlayerRoundState(defendingTeam)

// 	simulator := NewGameTickSimulator(
// 		totalTicks,
// 		match,
// 		tickRate,
// 		tickDuration,
// 	)

// 	for tick := uint(0); tick < simulator.TotalTicks; tick++ {
// 		exit := simulator.SimulateGameTick(&tick)
// 		if exit {
// 			break
// 		}
// 	}
// }

// func (t *GameTickSimulator) SimulateGameTick(tick *uint) bool {
// 	const killProbability = 0.0004 // Results in an average of ~5 kills per 100s round
// 	currentRound := t.Match.GameState.CurrentRound
// 	matchRng := t.Match.MatchRng

// 	if matchRng.Float64() < killProbability {
// 		currentRound.SimulatePlayerKill(matchRng)
// 	}

// 	// Attacker Plant Logic
// 	// after 15 seconds, there is a chance that the spike can be planted
// 	if *tick >= 15*t.TickRate && matchRng.Float64() < 0.0001 {
// 		attackingTeamPlayers := currentRound.GetAlivePlayers(currentRound.AttackingTeam)
// 		if len(attackingTeamPlayers) > 0 {
// 			currentRound.SpikePlanted = true
// 			currentRound.SpikePlantedBy = attackingTeamPlayers[matchRng.Intn(len(attackingTeamPlayers))].Player
// 		}
// 	}

// 	// Defuse Logic
// 	// after the spike has been planted, there is a chance that the defending team can defuse
// 	// Round ends if the defending team defuses the spike [defenders win]
// 	if currentRound.SpikePlanted && matchRng.Float64() < 0.0001 {
// 		defendingTeamPlayers := currentRound.GetAlivePlayers(currentRound.DefendingTeam)
// 		if len(defendingTeamPlayers) > 0 {
// 			currentRound.SpikeDefused = true
// 			currentRound.SpikeDefusedBy = defendingTeamPlayers[matchRng.Intn(len(defendingTeamPlayers))].Player

// 			t.Match.GameState.EndRound(currentRound.DefendingTeam, currentRound.AttackingTeam)
// 			return true
// 		}
// 	}

// 	// Reset total ticks if spike is planted
// 	if currentRound.SpikePlanted {
// 		*tick = 0
// 		t.TotalTicks = uint(t.Match.Rules.RoundDurationSpikePlantedSeconds) * t.TickRate
// 	}

// 	// Round ends if the attacking team successfully detonates the spike [attackers win]
// 	if currentRound.SpikeDetonated {
// 		t.Match.GameState.EndRound(currentRound.AttackingTeam, currentRound.DefendingTeam)
// 		return true
// 	}

// 	// Round ends if the attacking team defeats all defenders [attackers win]
// 	if len(currentRound.GetAlivePlayers(currentRound.DefendingTeam)) == 0 {
// 		t.Match.GameState.EndRound(currentRound.AttackingTeam, currentRound.DefendingTeam)
// 		return true
// 	}

// 	// Round ends if the defending team defeats all attackers before the spike is planted [defenders win]
// 	if len(currentRound.GetAlivePlayers(currentRound.AttackingTeam)) == 0 && !currentRound.SpikePlanted {
// 		t.Match.GameState.EndRound(currentRound.DefendingTeam, currentRound.AttackingTeam)
// 		return true
// 	}

// 	// Round ends if the time runs out before the attacking team plants the spike [defenders win]
// 	if *tick == t.TotalTicks-1 {
// 		t.Match.GameState.EndRound(currentRound.DefendingTeam, currentRound.AttackingTeam)
// 		return true
// 	}

// 	// Sleep for real time simulation
// 	if t.Match.IsRealTime {
// 		time.Sleep(t.TickDuration)
// 	}
// 	return false
// }

// func (roundState *RoundState) InitPlayerRoundState(team *Team) {
// 	for i := range team.Players {
// 		playerRoundState := NewPlayerRoundState(team, &team.Players[i], 0)
// 		roundState.PlayerRoundState = append(roundState.PlayerRoundState, &playerRoundState)
// 	}
// }

// func (roundState *RoundState) GetAlivePlayers(team *Team) []*PlayerRoundState {
// 	alivePlayers := make([]*PlayerRoundState, 0, 10)
// 	for _, p := range roundState.PlayerRoundState {
// 		if p.Alive && p.Team == team {
// 			alivePlayers = append(alivePlayers, p)
// 		}
// 	}
// 	return alivePlayers
// }
// func (roundState *RoundState) SimulatePlayerKill(rng *rand.Rand) PlayerKill {
// 	allAttackingTeamPlayers := roundState.GetAlivePlayers(roundState.AttackingTeam)
// 	allDefendingTeamPlayers := roundState.GetAlivePlayers(roundState.DefendingTeam)
// 	allAlivePlayers := append(allAttackingTeamPlayers, allDefendingTeamPlayers...)

// 	if len(allAlivePlayers) < 2 {
// 		return PlayerKill{}
// 	}

// 	// Pick a random killer.
// 	killer := allAlivePlayers[rng.Intn(len(allAlivePlayers))]

// 	// Pick a random victim from an opposing team.
// 	opponentsAlive := make([]*PlayerRoundState, 0, 5)
// 	for _, p := range allAlivePlayers {
// 		if p.Team != killer.Team {
// 			opponentsAlive = append(opponentsAlive, p)
// 		}
// 	}

// 	if len(opponentsAlive) == 0 {
// 		return PlayerKill{}
// 	}
// 	victim := opponentsAlive[rng.Intn(len(opponentsAlive))]
// 	victim.Alive = false

// 	return PlayerKill{
// 		Killer: killer.Player,
// 		Victim: victim.Player,
// 		Weapon: constants.Weapon{Name: constants.Sheriff}, //TODO: update this to the weapon used
// 	}
// }

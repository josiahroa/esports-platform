package event

// the game client uses events to send to the server
// the game server uses events to send to kinesis

// for the purposes of this simulator, we will be using the same event object for both scenarios.
// this means that a PlayerKillEvent sent from the game client to the game server will be the same as a PlayerKillEvent sent from the game server to kinesis.

import (
	"log/slog"
	"time"
)

type MatchEvent struct {
	matchID string
}

func NewMatchEvent(matchID string) *MatchEvent {
	return &MatchEvent{matchID: matchID}
}

func (l *MatchEvent) SetAttackingTeam(teamID string) {
	slog.Info("Set Attacking Team", "matchID", l.matchID, "teamID", teamID, "timestamp", time.Now().Format(time.RFC3339))
}

func (l *MatchEvent) SetDefendingTeam(teamID string) {
	slog.Info("Set Defending Team", "matchID", l.matchID, "teamID", teamID, "timestamp", time.Now().Format(time.RFC3339))
}

func (l *MatchEvent) MapSelected(mapName string) {
	slog.Info("Map Selected", "matchID", l.matchID, "mapName", mapName, "timestamp", time.Now().Format(time.RFC3339))
}

type RoundEvent struct {
	matchID string
	roundID int
}

func NewRoundEvent(matchID string, roundID int) *RoundEvent {
	return &RoundEvent{matchID: matchID, roundID: roundID}
}

func (l *RoundEvent) SpikeDefused(timelineTick int, playerID string) {
	slog.Info("Spike Defused", "playerID", playerID, "timestamp", time.Now().Format(time.RFC3339), "timelineTick", timelineTick)
}

func (l *RoundEvent) Clutch(timelineTick int, playerID string) {

}

func (l *RoundEvent) Ace(timelineTick int, playerID string) {}

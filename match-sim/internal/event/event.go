package event

import (
	"log/slog"
	"time"
)

type Event struct {
	matchID string
	roundID int
}

func NewEventLogger(matchID string) *Event {
	return &Event{matchID: matchID}
}

func NewRoundEvent(matchID string, roundID int) *Event {
	return &Event{matchID: matchID, roundID: roundID}
}

func (l *Event) SetAttackingTeam(teamID string) {
	slog.Info("Set Attacking Team", "matchID", l.matchID, "teamID", teamID, "timestamp", time.Now().Format(time.RFC3339))
}

func (l *Event) SetDefendingTeam(teamID string) {
	slog.Info("Set Defending Team", "matchID", l.matchID, "teamID", teamID, "timestamp", time.Now().Format(time.RFC3339))
}

func (l *Event) MapSelected(mapName string) {
	slog.Info("Map Selected", "matchID", l.matchID, "mapName", mapName, "timestamp", time.Now().Format(time.RFC3339))
}

func SpikeDefused(timelineTick int, playerID string) {
	slog.Info("Spike Defused", "playerID", playerID, "timestamp", time.Now().Format(time.RFC3339), "timelineTick", timelineTick)
}

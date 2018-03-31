package entity

import (
	"time"
)

// Match struct
type Match struct {
	ID               int64
	LeagueID         string
	ReporterPlayerID string
	reporterPlayer   Player
	CreatedAt        time.Time
	MatchPlayers     []MatchPlayer
}

// ReporterPlayer func
func (m Match) ReporterPlayer() Player {
	return m.reporterPlayer
}

// WinnerMatchPlayers func
func (m Match) WinnerMatchPlayers() []MatchPlayer {
	hasWon := true
	return m.getMatchPlayersByHasWon(hasWon)
}

// LoserMatchPlayers func
func (m Match) LoserMatchPlayers() []MatchPlayer {
	hasWon := false
	return m.getMatchPlayersByHasWon(hasWon)
}

func (m Match) getMatchPlayersByHasWon(hasWon bool) []MatchPlayer {
	matchPlayers := []MatchPlayer{}
	for i := range m.MatchPlayers {
		if m.MatchPlayers[i].HasWon == hasWon {
			matchPlayers = append(matchPlayers, m.MatchPlayers[i])
		}
	}
	return matchPlayers
}

// NewMatch factory
func NewMatch(match Match, reporterPlayer Player) Match {
	match.reporterPlayer = reporterPlayer
	return match
}

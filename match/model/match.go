package model

import (
	"time"

	leagueModel "github.com/szokodiakos/r8m8/league/model"
	"github.com/szokodiakos/r8m8/player/model"
)

// Match struct
type Match struct {
	ID             int64              `db:"id"`
	League         leagueModel.League `db:"league"`
	CreatedAt      time.Time          `db:"created_at"`
	ReporterPlayer model.Player       `db:"reporter_player"`
	MatchPlayers   []MatchPlayer
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
		if matchPlayers[i].HasWon == hasWon {
			matchPlayers = append(matchPlayers, m.MatchPlayers[i])
		}
	}
	return matchPlayers
}

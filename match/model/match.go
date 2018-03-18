package model

import (
	"time"

	"github.com/szokodiakos/r8m8/player/model"
)

// Match struct
type Match struct {
	ID                 int64        `db:"id"`
	LeagueID           int64        `db:"league_id"`
	ReporterPlayerID   int64        `db:"reporter_player_id"`
	CreatedAt          time.Time    `db:"created_at"`
	ReporterPlayer     model.Player `db:"reporter_player"`
	WinnerMatchPlayers []MatchPlayer
	LoserMatchPlayers  []MatchPlayer
}

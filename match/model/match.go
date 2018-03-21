package model

import (
	"time"

	leagueModel "github.com/szokodiakos/r8m8/league/model"
	"github.com/szokodiakos/r8m8/player/model"
)

// Match struct
type Match struct {
	ID                 int64              `db:"id"`
	League             leagueModel.League `db:"league"`
	CreatedAt          time.Time          `db:"created_at"`
	ReporterPlayer     model.Player       `db:"reporter_player"`
	WinnerMatchPlayers []MatchPlayer
	LoserMatchPlayers  []MatchPlayer
}

package model

import (
	leagueModel "github.com/szokodiakos/r8m8/league/model"
)

// MatchPlayer struct
type MatchPlayer struct {
	LeaguePlayer leagueModel.LeaguePlayer `db:"league_player"`
	RatingChange int                      `db:"rating_change"`
	HasWon       bool                     `db:"has_won"`
}

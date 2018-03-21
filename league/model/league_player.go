package model

import (
	playerModel "github.com/szokodiakos/r8m8/player/model"
)

// LeaguePlayer struct
type LeaguePlayer struct {
	Player     playerModel.Player `db:"player"`
	Rating     int                `db:"rating"`
	WinCount   int                `db:"won_match_count"`
	MatchCount int                `db:"total_match_count"`
}

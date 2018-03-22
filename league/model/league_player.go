package model

import (
	playerModel "github.com/szokodiakos/r8m8/player/model"
)

// LeaguePlayer struct
type LeaguePlayer struct {
	Player     playerModel.Player `db:"player"`
	League     League             `db:"league"`
	Rating     int                `db:"rating"`
	winCount   int                `db:"win_count"`
	matchCount int                `db:"match_count"`
}

// GetWinCount func
func (l LeaguePlayer) GetWinCount() int {
	return l.winCount
}

// GetMatchCount func
func (l LeaguePlayer) GetMatchCount() int {
	return l.matchCount
}

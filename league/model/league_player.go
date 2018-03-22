package model

import (
	playerModel "github.com/szokodiakos/r8m8/player/model"
)

// LeaguePlayer struct
type LeaguePlayer struct {
	Player     playerModel.Player `db:"player"`
	League     League             `db:"league"`
	Rating     int                `db:"rating"`
	WinCount   int                `db:"win_count"`
	MatchCount int                `db:"match_count"`
}

// GetWinCount func
func (l LeaguePlayer) GetWinCount() int {
	return l.WinCount
}

// GetMatchCount func
func (l LeaguePlayer) GetMatchCount() int {
	return l.MatchCount
}

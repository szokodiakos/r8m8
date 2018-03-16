package model

import (
	playerModel "github.com/szokodiakos/r8m8/player/model"
)

// MatchStats struct
type MatchStats struct {
	ReporterPlayer          playerModel.Player
	WinnerMatchPlayersStats []MatchPlayerStats
	LoserMatchPlayersStats  []MatchPlayerStats
}

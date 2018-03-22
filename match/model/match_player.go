package model

import (
	playerModel "github.com/szokodiakos/r8m8/player/model"
)

// MatchPlayer struct
type MatchPlayer struct {
	Player       playerModel.Player `db:"player"`
	RatingChange int                `db:"rating_change"`
	HasWon       bool               `db:"has_won"`
}

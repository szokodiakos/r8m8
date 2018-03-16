package model

import "github.com/szokodiakos/r8m8/player/model"

// Details struct
type Details struct {
	PlayerID     int64 `db:"player_id"`
	MatchID      int64 `db:"match_id"`
	RatingChange int   `db:"rating_change"`
	HasWon       bool  `db:"has_won"`
	model.Player `db:"player"`
}

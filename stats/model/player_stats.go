package model

import (
	playerModel "github.com/szokodiakos/r8m8/player/model"
	ratingModel "github.com/szokodiakos/r8m8/rating/model"
)

// PlayerStats struct
type PlayerStats struct {
	Player     playerModel.Player `db:"player"`
	Rating     ratingModel.Rating `db:"rating"`
	WinCount   int                `db:"won_match_count"`
	MatchCount int                `db:"total_match_count"`
}

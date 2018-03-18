package model

import (
	detailsModel "github.com/szokodiakos/r8m8/details/model"
	playerModel "github.com/szokodiakos/r8m8/player/model"
	ratingModel "github.com/szokodiakos/r8m8/rating/model"
)

// MatchPlayer struct
type MatchPlayer struct {
	Player  playerModel.Player   `db:"player"`
	Rating  ratingModel.Rating   `db:"rating"`
	Details detailsModel.Details `db:"details"`
}

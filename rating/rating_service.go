package rating

import (
	"github.com/szokodiakos/r8m8/player"
)

// Service interface
type Service interface {
	CalculateRating(winnerDBPlayers []player.DBPlayer, loserDBPlayers []player.DBPlayer) ([]player.DBPlayer, []player.DBPlayer)
}

type ratingService struct {
}

func (r *ratingService) CalculateRating(winnerDBPlayers []player.DBPlayer, loserDBPlayers []player.DBPlayer) ([]player.DBPlayer, []player.DBPlayer) {
	return winnerDBPlayers, loserDBPlayers
}

// NewService factory
func NewService() Service {
	return &ratingService{}
}

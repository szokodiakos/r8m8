package rating

import (
	"github.com/szokodiakos/r8m8/player"
)

// Service interface
type Service interface {
	CalculateRating(winnerPlayers []player.Player, loserPlayers []player.Player) ([]player.Player, []player.Player)
}

type ratingService struct {
}

func (rs *ratingService) CalculateRating(winnerPlayers []player.Player, loserPlayers []player.Player) ([]player.Player, []player.Player) {
	return winnerPlayers, loserPlayers
}

// NewService factory
func NewService() Service {
	return &ratingService{}
}

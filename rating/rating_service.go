package rating

import (
	"github.com/szokodiakos/r8m8/player"
)

// Service interface
type Service interface {
	CalculateRating(winnerRepoPlayers []player.RepoPlayer, loserRepoPlayers []player.RepoPlayer) ([]player.RepoPlayer, []player.RepoPlayer)
}

type ratingService struct {
}

func (r *ratingService) CalculateRating(winnerRepoPlayers []player.RepoPlayer, loserRepoPlayers []player.RepoPlayer) ([]player.RepoPlayer, []player.RepoPlayer) {
	return winnerRepoPlayers, loserRepoPlayers
}

// NewService factory
func NewService() Service {
	return &ratingService{}
}

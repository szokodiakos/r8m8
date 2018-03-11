package match

import (
	"github.com/szokodiakos/r8m8/player"
	"github.com/szokodiakos/r8m8/transaction"
)

// DetailsService interface
type DetailsService interface {
	AddMultiple(transaction transaction.Transaction, matchID int64, repoPlayers []player.RepoPlayer, adjustedRepoPlayers []player.RepoPlayer) error
}

type matchDetailsService struct {
	matchDetailsRepository DetailsRepository
}

func (m *matchDetailsService) AddMultiple(transaction transaction.Transaction, matchID int64, repoPlayers []player.RepoPlayer, adjustedRepoPlayers []player.RepoPlayer) error {
	for i := range repoPlayers {
		ratingChange := getRatingChange(repoPlayers[i], adjustedRepoPlayers[i])
		matchDetails := Details{
			PlayerID:     repoPlayers[i].ID,
			MatchID:      matchID,
			RatingChange: ratingChange,
		}

		if err := m.matchDetailsRepository.Create(transaction, matchDetails); err != nil {
			return err
		}
	}

	return nil
}

func getRatingChange(repoPlayer player.RepoPlayer, adjustedRepoPlayer player.RepoPlayer) int {
	return adjustedRepoPlayer.Rating - repoPlayer.Rating
}

// NewDetailsService factory
func NewDetailsService(matchDetailsRepository DetailsRepository) DetailsService {
	return &matchDetailsService{
		matchDetailsRepository: matchDetailsRepository,
	}
}

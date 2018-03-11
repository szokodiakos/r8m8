package match

import (
	"github.com/szokodiakos/r8m8/player"
	"github.com/szokodiakos/r8m8/transaction"
)

// DetailsService interface
type DetailsService interface {
	AddMultiple(transaction transaction.Transaction, matchID int64, dbPlayers []player.DBPlayer, adjustedDBPlayers []player.DBPlayer) error
}

type matchDetailsService struct {
	matchDetailsRepository DetailsRepository
}

func (m *matchDetailsService) AddMultiple(transaction transaction.Transaction, matchID int64, dbPlayers []player.DBPlayer, adjustedDBPlayers []player.DBPlayer) error {
	for i := range dbPlayers {
		ratingChange := getRatingChange(dbPlayers[i], adjustedDBPlayers[i])
		matchDetails := Details{
			PlayerID:     dbPlayers[i].ID,
			MatchID:      matchID,
			RatingChange: ratingChange,
		}

		if err := m.matchDetailsRepository.Create(transaction, matchDetails); err != nil {
			return err
		}
	}

	return nil
}

func getRatingChange(dbPlayer player.DBPlayer, adjustedDBPlayer player.DBPlayer) int {
	return adjustedDBPlayer.Rating - dbPlayer.Rating
}

// NewDetailsService factory
func NewDetailsService(matchDetailsRepository DetailsRepository) DetailsService {
	return &matchDetailsService{
		matchDetailsRepository: matchDetailsRepository,
	}
}

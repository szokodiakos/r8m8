package match

import (
	"github.com/szokodiakos/r8m8/player"
	"github.com/szokodiakos/r8m8/transaction"
)

// DetailsService interface
type DetailsService interface {
	AddMultiple(transaction transaction.Transaction, matchID int64, players []player.Player, adjustedPlayers []player.Player) error
}

type matchDetailsService struct {
	matchDetailsRepository DetailsRepository
}

func (mds *matchDetailsService) AddMultiple(transaction transaction.Transaction, matchID int64, players []player.Player, adjustedPlayers []player.Player) error {
	for i, player := range players {
		ratingChange := mds.getRatingChange(player, adjustedPlayers[i])
		matchDetails := Details{
			PlayerID:     player.ID,
			MatchID:      matchID,
			RatingChange: ratingChange,
		}

		if err := mds.matchDetailsRepository.Create(transaction, matchDetails); err != nil {
			return err
		}
	}

	return nil
}

func (mds *matchDetailsService) getRatingChange(player player.Player, adjustedPlayer player.Player) int {
	return adjustedPlayer.Rating - player.Rating
}

// NewDetailsService factory
func NewDetailsService(matchDetailsRepository DetailsRepository) DetailsService {
	return &matchDetailsService{
		matchDetailsRepository: matchDetailsRepository,
	}
}

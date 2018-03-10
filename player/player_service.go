package player

import "github.com/szokodiakos/r8m8/transaction"

// Service interface
type Service interface {
	AddMultiple(transaction transaction.Transaction, count int) ([]int64, error)
	UpdateRatingsForMultiple(transaction transaction.Transaction, players []Player) error
}

type playerService struct {
	playerRepository Repository
}

func (ps *playerService) AddMultiple(transaction transaction.Transaction, count int) ([]int64, error) {
	playerIDs := make([]int64, 0, count)
	for i := 0; i < count; i++ {
		playerID, err := ps.playerRepository.Create(transaction)

		if err != nil {
			return playerIDs, err
		}

		playerIDs = append(playerIDs, playerID)
	}

	return playerIDs, nil
}

func (ps *playerService) UpdateRatingsForMultiple(transaction transaction.Transaction, players []Player) error {
	for i := range players {
		if err := ps.playerRepository.UpdateRatingByID(transaction, players[i].ID, players[i].Rating); err != nil {
			return err
		}
	}

	return nil
}

// NewService factory
func NewService(playerRepository Repository) Service {
	return &playerService{
		playerRepository: playerRepository,
	}
}

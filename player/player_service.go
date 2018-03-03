package player

// Service interface
type Service interface {
	AddMultiple(count int) ([]int64, error)
}

type playerService struct {
	playerRepository Repository
}

func (ps *playerService) AddMultiple(count int) ([]int64, error) {
	playerIDs := make([]int64, 0, count)
	for i := 0; i < count; i++ {
		playerID, err := ps.playerRepository.Create()

		if err != nil {
			return playerIDs, err
		}

		playerIDs = append(playerIDs, playerID)
	}

	return playerIDs, nil
}

// NewService factory
func NewService(playerRepository Repository) Service {
	return &playerService{
		playerRepository: playerRepository,
	}
}

package match

// Service interface
type Service interface {
	AddMatch() (Match, error)
}

type matchService struct {
	repository Repository
}

func (ms *matchService) AddMatch() (Match, error) {
	var createdMatch Match
	return createdMatch, nil
}

// NewService creates a service
func NewService(matchRepository Repository) Service {
	return &matchService{
		repository: matchRepository,
	}
}

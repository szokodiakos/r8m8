package match

// Service interface
type Service interface {
	AddMatch(match Match) Match
}

type matchService struct {
	repository Repository
}

func (ms *matchService) AddMatch(match Match) Match {
	var m Match
	return m
}

// NewService creates a service
func NewService(matchRepository Repository) Service {
	return &matchService{
		repository: matchRepository,
	}
}

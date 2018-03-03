package match

import "github.com/szokodiakos/r8m8/transaction"

// Service interface
type Service interface {
	AddMatch(match Match) (Match, error)
}

type matchService struct {
	transactionService transaction.Service
	repository         Repository
}

func (ms *matchService) AddMatch(match Match) (Match, error) {
	var createdMatch Match

	tr, err := ms.transactionService.Start()
	if err != nil {
		return createdMatch, err
	}
	defer ms.transactionService.CommitOrRollback(tr)

	return createdMatch, nil
}

// NewService creates a service
func NewService(transactionService transaction.Service, matchRepository Repository) Service {
	return &matchService{
		transactionService: transactionService,
		repository:         matchRepository,
	}
}

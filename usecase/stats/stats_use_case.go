package stats

import (
	"github.com/szokodiakos/r8m8/league"
	"github.com/szokodiakos/r8m8/transaction"
)

// UseCase interface
type UseCase interface {
	Handle(input Input) (Output, error)
}

type statsUseCase struct {
	transactionService transaction.Service
	leagueRepository   league.Repository
}

func (s *statsUseCase) Handle(input Input) (Output, error) {
	var output Output
	league := input.League

	tr, err := s.transactionService.Start()
	if err != nil {
		return output, err
	}

	_, err = s.leagueRepository.GetByID(tr, league.ID)
	if err != nil {
		return output, s.transactionService.Rollback(tr, err)
	}

	output = Output{}
	err = s.transactionService.Commit(tr)
	return output, err
}

// NewUseCase factory
func NewUseCase(
	transactionService transaction.Service,
	leagueRepository league.Repository,
) UseCase {
	return &statsUseCase{
		transactionService: transactionService,
		leagueRepository:   leagueRepository,
	}
}

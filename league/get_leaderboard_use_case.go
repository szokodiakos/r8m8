package league

import (
	"github.com/szokodiakos/r8m8/league/model"
	"github.com/szokodiakos/r8m8/transaction"
)

// GetLeaderboardUseCase interface
type GetLeaderboardUseCase interface {
	Handle(input model.GetLeaderboardInput) (model.GetLeaderboardOutput, error)
}

type getLeaderboardUseCase struct {
	transactionService transaction.Service
	leagueRepository   Repository
}

func (g *getLeaderboardUseCase) Handle(input model.GetLeaderboardInput) (model.GetLeaderboardOutput, error) {
	var output model.GetLeaderboardOutput
	league := input.League

	tr, err := g.transactionService.Start()
	if err != nil {
		return output, err
	}

	repoLeague, err := g.leagueRepository.GetByUniqueName(tr, league.UniqueName)
	if err != nil {
		return output, g.transactionService.Rollback(tr, err)
	}

	output = model.GetLeaderboardOutput{
		League: repoLeague,
	}
	err = g.transactionService.Commit(tr)
	return output, err
}

// NewGetLeaderboardUseCase factory
func NewGetLeaderboardUseCase(
	transactionService transaction.Service,
	leagueRepository Repository,
) GetLeaderboardUseCase {
	return &getLeaderboardUseCase{
		transactionService: transactionService,
		leagueRepository:   leagueRepository,
	}
}

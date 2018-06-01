package leaderboard

import (
	"github.com/szokodiakos/r8m8/league"
	"github.com/szokodiakos/r8m8/transaction"
)

// UseCase interface
type UseCase interface {
	Handle(input Input) (Output, error)
}

type leaderboardUseCase struct {
	transactionService transaction.Service
	leagueRepository   league.Repository
}

func (g *leaderboardUseCase) Handle(input Input) (Output, error) {
	var output Output
	league := input.League

	tr, err := g.transactionService.Start()
	if err != nil {
		return output, err
	}

	repoLeague, err := g.leagueRepository.GetByID(tr, league.ID)
	if err != nil {
		return output, g.transactionService.Rollback(tr, err)
	}

	repoLeague.LeaguePlayers = repoLeague.GetTopLeaguePlayers()

	output = Output{
		League: repoLeague,
	}
	err = g.transactionService.Commit(tr)
	return output, err
}

// NewUseCase factory
func NewUseCase(
	transactionService transaction.Service,
	leagueRepository league.Repository,
) UseCase {
	return &leaderboardUseCase{
		transactionService: transactionService,
		leagueRepository:   leagueRepository,
	}
}

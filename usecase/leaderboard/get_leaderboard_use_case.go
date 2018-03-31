package leaderboard

import (
	"github.com/szokodiakos/r8m8/entity"
	"github.com/szokodiakos/r8m8/transaction"
)

// UseCase interface
type UseCase interface {
	Handle(input Input) (Output, error)
}

type getLeaderboardUseCase struct {
	transactionService transaction.Service
	leagueRepository   entity.LeagueRepository
}

func (g *getLeaderboardUseCase) Handle(input Input) (Output, error) {
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

	repoLeague.LeaguePlayers = repoLeague.GetTop10LeaguePlayers()

	output = Output{
		League: repoLeague,
	}
	err = g.transactionService.Commit(tr)
	return output, err
}

// NewGetLeaderboardUseCase factory
func NewGetLeaderboardUseCase(
	transactionService transaction.Service,
	leagueRepository entity.LeagueRepository,
) UseCase {
	return &getLeaderboardUseCase{
		transactionService: transactionService,
		leagueRepository:   leagueRepository,
	}
}

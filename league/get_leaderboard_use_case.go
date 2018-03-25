package league

import (
	"github.com/szokodiakos/r8m8/entity"
	"github.com/szokodiakos/r8m8/transaction"
)

// GetLeaderboardUseCase interface
type GetLeaderboardUseCase interface {
	Handle(input GetLeaderboardInput) (GetLeaderboardOutput, error)
}

type getLeaderboardUseCase struct {
	transactionService transaction.Service
	leagueRepository   entity.LeagueRepository
}

func (g *getLeaderboardUseCase) Handle(input GetLeaderboardInput) (GetLeaderboardOutput, error) {
	var output GetLeaderboardOutput
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

	output = GetLeaderboardOutput{
		League: repoLeague,
	}
	err = g.transactionService.Commit(tr)
	return output, err
}

// NewGetLeaderboardUseCase factory
func NewGetLeaderboardUseCase(
	transactionService transaction.Service,
	leagueRepository entity.LeagueRepository,
) GetLeaderboardUseCase {
	return &getLeaderboardUseCase{
		transactionService: transactionService,
		leagueRepository:   leagueRepository,
	}
}

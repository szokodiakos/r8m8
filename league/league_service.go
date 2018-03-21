package league

import (
	"github.com/szokodiakos/r8m8/league/errors"
	"github.com/szokodiakos/r8m8/league/model"
	"github.com/szokodiakos/r8m8/transaction"
)

// Service interface
type Service interface {
	GetOrAdd(tr transaction.Transaction, league model.League) (model.League, error)
}

type leagueService struct {
	leagueRepository Repository
}

func (l *leagueService) GetOrAdd(tr transaction.Transaction, league model.League) (model.League, error) {
	repoLeague, err := l.leagueRepository.GetByUniqueName(tr, league.UniqueName)
	if err != nil {
		return l.handleGetLeagueError(tr, league, err)
	}
	return repoLeague, err
}

func (l *leagueService) handleGetLeagueError(tr transaction.Transaction, league model.League, err error) (model.League, error) {
	switch err.(type) {
	case *errors.LeagueNotFoundError:
		err = l.leagueRepository.Create(tr, league)
		if err != nil {
			return league, err
		}
		return l.GetOrAdd(tr, league)
	default:
		return league, err
	}
}

// NewService factory
func NewService(leagueRepository Repository) Service {
	return &leagueService{
		leagueRepository: leagueRepository,
	}
}

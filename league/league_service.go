package league

import (
	"github.com/szokodiakos/r8m8/transaction"
)

// Service interface
type Service interface {
	GetOrAddLeague(tr transaction.Transaction, league League) (League, error)
}

type leagueService struct {
	leagueRepository Repository
}

func (l *leagueService) GetOrAddLeague(tr transaction.Transaction, league League) (League, error) {
	repoLeague, err := l.leagueRepository.GetByUniqueName(tr, league.UniqueName)
	if err != nil {
		return repoLeague, err
	}

	if repoLeague == (League{}) {
		err = l.leagueRepository.Create(tr, league)
		if err != nil {
			return repoLeague, err
		}

		repoLeague, err = l.leagueRepository.GetByUniqueName(tr, league.UniqueName)
		if err != nil {
			return repoLeague, err
		}
	}

	return repoLeague, err
}

// NewService factory
func NewService(leagueRepository Repository) Service {
	return &leagueService{
		leagueRepository: leagueRepository,
	}
}

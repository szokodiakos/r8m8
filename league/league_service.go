package league

import (
	"github.com/szokodiakos/r8m8/league/errors"
	"github.com/szokodiakos/r8m8/league/model"
	"github.com/szokodiakos/r8m8/player"
	playerModel "github.com/szokodiakos/r8m8/player/model"
	"github.com/szokodiakos/r8m8/transaction"
)

// Service interface
type Service interface {
	GetOrAdd(tr transaction.Transaction, league model.League) (model.League, error)
	AddAnyMissingPlayers(tr transaction.Transaction, league model.League, players []playerModel.Player) error
}

type leagueService struct {
	playerService       player.Service
	leaguePlayerService PlayerService
	leagueRepository    Repository
}

func (l *leagueService) GetOrAdd(tr transaction.Transaction, league model.League) (model.League, error) {
	repoLeague, err := l.leagueRepository.GetByUniqueName(tr, league.UniqueName)
	if err != nil {
		return l.createIfNotExists(tr, league, err)
	}
	return repoLeague, err
}

func (l *leagueService) createIfNotExists(tr transaction.Transaction, league model.League, err error) (model.League, error) {
	switch err.(type) {
	case *errors.LeagueNotFoundError:
		return l.leagueRepository.Create(tr, league)
	default:
		return league, err
	}
}

func (l *leagueService) AddAnyMissingPlayers(tr transaction.Transaction, league model.League, players []playerModel.Player) error {
	err := l.playerService.AddAnyMissing(tr, players)
	if err != nil {
		return err
	}

	err = l.leaguePlayerService.AddAnyMissing(tr, players, league)
	return err
}

// NewService factory
func NewService(playerService player.Service, leaguePlayerService PlayerService, leagueRepository Repository) Service {
	return &leagueService{
		playerService:       playerService,
		leaguePlayerService: leaguePlayerService,
		leagueRepository:    leagueRepository,
	}
}

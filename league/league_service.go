package league

import (
	"github.com/szokodiakos/r8m8/entity"
	"github.com/szokodiakos/r8m8/league/errors"
	"github.com/szokodiakos/r8m8/player"
	"github.com/szokodiakos/r8m8/transaction"
)

// Service interface
type Service interface {
	GetOrAddLeague(tr transaction.Transaction, league entity.League, players []entity.Player) (entity.League, error)
}

type leagueService struct {
	leagueRepository LeagueRepository
	playerRepository player.PlayerRepository
	initialRating    int
}

func (l *leagueService) GetOrAddLeague(tr transaction.Transaction, league entity.League, players []entity.Player) (entity.League, error) {
	repoLeague, err := l.leagueRepository.GetByID(tr, league.ID)
	if err != nil {
		switch err.(type) {
		case *errors.LeagueNotFoundError:
			return l.addLeague(tr, league, players)
		default:
			return league, err
		}
	}
	return repoLeague, err
}

func (l *leagueService) addLeague(tr transaction.Transaction, league entity.League, players []entity.Player) (entity.League, error) {
	repoPlayers, err := l.addPlayers(tr, players)
	if err != nil {
		return league, err
	}

	leaguePlayers := l.createLeaguePlayers(repoPlayers)
	league.LeaguePlayers = leaguePlayers

	return l.leagueRepository.Add(tr, league)
}

func (l *leagueService) addPlayers(tr transaction.Transaction, players []entity.Player) ([]entity.Player, error) {
	repoPlayers := make([]entity.Player, len(players))
	for i := range players {
		repoPlayer, err := l.playerRepository.Add(tr, players[i])
		if err != nil {
			return repoPlayers, err
		}
		repoPlayers[i] = repoPlayer
	}
	return repoPlayers, nil
}

func (l *leagueService) createLeaguePlayers(repoPlayers []entity.Player) []entity.LeaguePlayer {
	leaguePlayers := make([]entity.LeaguePlayer, len(repoPlayers))
	for i := range repoPlayers {
		leaguePlayers[i] = entity.LeaguePlayer{
			PlayerID: repoPlayers[i].ID,
			Rating:   l.initialRating,
		}
	}
	return leaguePlayers
}

// NewService factory
func NewService(
	leagueRepository LeagueRepository,
	playerRepository player.PlayerRepository,
	initialRating int,
) Service {
	return &leagueService{
		leagueRepository: leagueRepository,
		playerRepository: playerRepository,
		initialRating:    initialRating,
	}
}

package league

import (
	"github.com/szokodiakos/r8m8/league/errors"
	"github.com/szokodiakos/r8m8/league/model"
	"github.com/szokodiakos/r8m8/player"
	playerModel "github.com/szokodiakos/r8m8/player/model"
	"github.com/szokodiakos/r8m8/transaction"
)

// PlayerService interface
type PlayerService interface {
	GetMultipleByUniqueNamesInOrder(tr transaction.Transaction, uniqueNames []string) ([]model.LeaguePlayer, error)
	UpdateMultiple(tr transaction.Transaction, leaguePlayers []model.LeaguePlayer) error
	AddAnyMissing(tr transaction.Transaction, players []playerModel.Player, league model.League) error
}

type leaguePlayerService struct {
	playerService          player.Service
	playerRepository       player.Repository
	leaguePlayerRepository PlayerRepository
	initialRating          int
}

func (l *leaguePlayerService) GetMultipleByUniqueNamesInOrder(tr transaction.Transaction, uniqueNames []string) ([]model.LeaguePlayer, error) {
	repoLeaguePlayers, err := l.leaguePlayerRepository.GetMultipleByPlayerUniqueNames(tr, uniqueNames)
	if err != nil {
		return repoLeaguePlayers, err
	}

	orderedRepoLeaguePlayers, err := sortLeaguePlayersByUniqueNames(repoLeaguePlayers, uniqueNames)
	return orderedRepoLeaguePlayers, err
}

func sortLeaguePlayersByUniqueNames(leaguePlayers []model.LeaguePlayer, uniqueNames []string) ([]model.LeaguePlayer, error) {
	orderedLeaguePlayers := make([]model.LeaguePlayer, len(leaguePlayers))
	for i := range uniqueNames {
		player, err := getLeaguePlayerByUniqueName(leaguePlayers, uniqueNames[i])
		if err != nil {
			return orderedLeaguePlayers, err
		}
		orderedLeaguePlayers[i] = player
	}
	return orderedLeaguePlayers, nil
}

func getLeaguePlayerByUniqueName(leaguePlayers []model.LeaguePlayer, uniqueName string) (model.LeaguePlayer, error) {
	for i := range leaguePlayers {
		if leaguePlayers[i].Player.UniqueName == uniqueName {
			return leaguePlayers[i], nil
		}
	}
	return model.LeaguePlayer{}, &errors.LeaguePlayerNotFoundError{
		UniqueName: uniqueName,
	}
}

func (l *leaguePlayerService) UpdateMultiple(tr transaction.Transaction, leaguePlayers []model.LeaguePlayer) error {
	for i := range leaguePlayers {
		err := l.leaguePlayerRepository.Update(tr, leaguePlayers[i])
		if err != nil {
			return err
		}
	}
	return nil
}

func (l *leaguePlayerService) AddAnyMissing(tr transaction.Transaction, players []playerModel.Player, league model.League) error {
	uniqueNames := l.playerService.MapToUniqueNames(players)
	repoLeaguePlayers, err := l.leaguePlayerRepository.GetMultipleByPlayerUniqueNames(tr, uniqueNames)
	if err != nil {
		return err
	}

	if isMissingLeaguePlayerExists(players, repoLeaguePlayers) {
		repoPlayers, err := l.playerRepository.GetMultipleByUniqueNames(tr, uniqueNames)
		if err != nil {
			return err
		}

		missingLeaguePlayers := l.getMissingLeaguePlayers(repoPlayers, repoLeaguePlayers, league)
		return l.addMultiple(tr, missingLeaguePlayers)
	}

	return nil
}

func isMissingLeaguePlayerExists(players []playerModel.Player, repoLeaguePlayers []model.LeaguePlayer) bool {
	return (len(players) != len(repoLeaguePlayers))
}

func (l *leaguePlayerService) getMissingLeaguePlayers(repoPlayers []playerModel.Player, repoLeaguePlayers []model.LeaguePlayer, league model.League) []model.LeaguePlayer {
	missingLeaguePlayers := []model.LeaguePlayer{}

	for i := range repoPlayers {
		err := testRepoLeaguePlayerMissing(repoPlayers[i], repoLeaguePlayers)
		switch err.(type) {
		case *errors.LeaguePlayerNotFoundError:
			missingLeaguePlayers = append(missingLeaguePlayers, model.LeaguePlayer{
				Player: repoPlayers[i],
				League: league,
				Rating: l.initialRating,
			})
		}
	}

	return missingLeaguePlayers
}

func testRepoLeaguePlayerMissing(repoPlayer playerModel.Player, repoLeaguePlayers []model.LeaguePlayer) error {
	for i := range repoLeaguePlayers {
		if repoLeaguePlayers[i].Player.UniqueName == repoPlayer.UniqueName {
			return nil
		}
	}

	return &errors.LeaguePlayerNotFoundError{
		UniqueName: repoPlayer.UniqueName,
	}
}

func (l *leaguePlayerService) addMultiple(tr transaction.Transaction, leaguePlayers []model.LeaguePlayer) error {
	for i := range leaguePlayers {
		err := l.leaguePlayerRepository.Create(tr, leaguePlayers[i])
		if err != nil {
			return err
		}
	}
	return nil
}

// NewPlayerService factory
func NewPlayerService(playerService player.Service, playerRepository player.Repository, leaguePlayerRepository PlayerRepository, initialRating int) PlayerService {
	return &leaguePlayerService{
		playerService:          playerService,
		playerRepository:       playerRepository,
		leaguePlayerRepository: leaguePlayerRepository,
		initialRating:          initialRating,
	}
}

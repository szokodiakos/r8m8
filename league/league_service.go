package league

import (
	"github.com/szokodiakos/r8m8/league/errors"
	"github.com/szokodiakos/r8m8/league/model"
	"github.com/szokodiakos/r8m8/player"
	playerModel "github.com/szokodiakos/r8m8/player/model"
	"github.com/szokodiakos/r8m8/rating"
	ratingModel "github.com/szokodiakos/r8m8/rating/model"
	"github.com/szokodiakos/r8m8/transaction"
)

// Service interface
type Service interface {
	GetOrAdd(tr transaction.Transaction, league model.League) (model.League, error)
	AddAnyMissingPlayers(tr transaction.Transaction, league model.League, players []playerModel.Player) error
}

type leagueService struct {
	playerService    player.Service
	leagueRepository Repository
	playerRepository player.Repository
	ratingRepository rating.Repository
	initialRating    int
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
		return l.leagueRepository.Create(tr, league)
	default:
		return league, err
	}
}

func (l *leagueService) AddAnyMissingPlayers(tr transaction.Transaction, league model.League, players []playerModel.Player) error {
	repoPlayers, err := l.playerService.GetRepoPlayers(tr, players)
	if err != nil {
		return err
	}

	if isMissingPlayerExists(players, repoPlayers) {
		missingPlayers := getMissingPlayers(players, repoPlayers)
		err := l.addMissingPlayers(tr, league, missingPlayers)
		if err != nil {
			return err
		}
	}

	return nil
}

func mapToUniqueNames(players []playerModel.Player) []string {
	uniqueNames := make([]string, len(players))
	for i := range players {
		uniqueNames[i] = players[i].UniqueName
	}
	return uniqueNames
}

func isMissingPlayerExists(players []playerModel.Player, repoPlayers []playerModel.Player) bool {
	return (len(players) != len(repoPlayers))
}

func getMissingPlayers(players []playerModel.Player, repoPlayers []playerModel.Player) []playerModel.Player {
	missingPlayers := make([]playerModel.Player, 0, len(repoPlayers))

	for i := range players {
		repoPlayer := getRepoCounterpart(players[i], repoPlayers)

		if repoPlayer == (playerModel.Player{}) {
			missingPlayers = append(missingPlayers, players[i])
		}
	}
	return missingPlayers
}

func getRepoCounterpart(player playerModel.Player, repoPlayers []playerModel.Player) playerModel.Player {
	var repoPlayer playerModel.Player

	for i := range repoPlayers {
		if player.UniqueName == repoPlayers[i].UniqueName {
			repoPlayer = repoPlayers[i]
		}
	}

	return repoPlayer
}

func (l *leagueService) addMissingPlayers(tr transaction.Transaction, league model.League, players []playerModel.Player) error {
	for i := range players {

		player, err := l.playerRepository.Create(tr, players[i])
		if err != nil {
			return err
		}

		rating := ratingModel.Rating{
			PlayerID: player.ID,
			LeagueID: league.ID,
			Rating:   l.initialRating,
		}
		err = l.ratingRepository.Create(tr, rating)
		if err != nil {
			return err
		}
	}

	return nil
}

// NewService factory
func NewService(playerService player.Service, leagueRepository Repository, playerRepository player.Repository, ratingRepository rating.Repository, initialRating int) Service {
	return &leagueService{
		playerService:    playerService,
		leagueRepository: leagueRepository,
		playerRepository: playerRepository,
		ratingRepository: ratingRepository,
		initialRating:    initialRating,
	}
}

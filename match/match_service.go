package match

import (
	"github.com/szokodiakos/r8m8/match/errors"
	"github.com/szokodiakos/r8m8/player"
	"github.com/szokodiakos/r8m8/rating"
	"github.com/szokodiakos/r8m8/transaction"
)

// Service interface
type Service interface {
	Add(transaction transaction.Transaction, players []player.Player) error
}

type matchService struct {
	matchRepository     Repository
	ratingService       rating.Service
	playerService       player.Service
	matchDetailsService DetailsService
}

func (m *matchService) Add(transaction transaction.Transaction, players []player.Player) error {
	if isPlayerCountUneven(players) {
		return errors.NewUnevenMatchPlayersError()
	}

	repoPlayers, err := m.playerService.GetOrAddPlayers(transaction, players)
	if err != nil {
		return err
	}

	winnerRepoPlayers := getWinnerRepoPlayers(repoPlayers)
	loserRepoPlayers := getLoserRepoPlayers(repoPlayers)
	adjustedWinnerRepoPlayers, adjustedLoserRepoPlayers := m.ratingService.CalculateRating(winnerRepoPlayers, loserRepoPlayers)
	adjustedRepoPlayers := append(adjustedWinnerRepoPlayers, adjustedLoserRepoPlayers...)

	if err := m.playerService.UpdateRatingsForMultiple(transaction, adjustedRepoPlayers); err != nil {
		return err
	}

	matchID, err := m.matchRepository.Create(transaction)
	if err != nil {
		return err
	}

	err = m.matchDetailsService.AddMultiple(transaction, matchID, repoPlayers, adjustedRepoPlayers)
	return err
}

func isPlayerCountUneven(players []player.Player) bool {
	return (len(players) % 2) != 0
}

func getWinnerRepoPlayers(players []player.RepoPlayer) []player.RepoPlayer {
	lowerhalf := players[:(len(players) / 2)]
	return lowerhalf
}

func getLoserRepoPlayers(players []player.RepoPlayer) []player.RepoPlayer {
	upperhalf := players[(len(players) / 2):]
	return upperhalf
}

// NewService creates a service
func NewService(matchRepository Repository, ratingService rating.Service, playerService player.Service, matchDetailsService DetailsService) Service {
	return &matchService{
		matchRepository:     matchRepository,
		ratingService:       ratingService,
		playerService:       playerService,
		matchDetailsService: matchDetailsService,
	}
}

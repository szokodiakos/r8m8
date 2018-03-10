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

func (ms *matchService) Add(transaction transaction.Transaction, players []player.Player) error {
	if ms.isPlayerCountUneven(players) {
		return errors.NewUnevenMatchPlayersError()
	}

	winnerPlayers := ms.getWinnerPlayers(players)
	loserPlayers := ms.getLoserPlayers(players)
	adjustedWinnerPlayers, adjustedLoserPlayers := ms.ratingService.CalculateRating(winnerPlayers, loserPlayers)
	adjustedPlayers := append(adjustedWinnerPlayers, adjustedLoserPlayers...)

	if err := ms.playerService.UpdateRatingsForMultiple(transaction, adjustedPlayers); err != nil {
		return err
	}

	matchID, err := ms.matchRepository.Create(transaction)
	if err != nil {
		return err
	}

	err = ms.matchDetailsService.AddMultiple(transaction, matchID, players, adjustedPlayers)
	return err
}

func (ms *matchService) isPlayerCountUneven(players []player.Player) bool {
	return (len(players) % 2) != 0
}

func (ms *matchService) getWinnerPlayers(players []player.Player) []player.Player {
	lowerhalfPlayers := players[:(len(players) / 2)]
	return lowerhalfPlayers
}

func (ms *matchService) getLoserPlayers(players []player.Player) []player.Player {
	upperhalfPlayers := players[(len(players) / 2):]
	return upperhalfPlayers
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

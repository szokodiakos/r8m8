package match

import (
	"github.com/szokodiakos/r8m8/match/errors"
	"github.com/szokodiakos/r8m8/player"
)

// Service interface
type Service interface {
	Add(players []player.Player) (Match, error)
}

type matchService struct {
	matchRepository Repository
}

func (ms *matchService) Add(players []player.Player) (Match, error) {
	var createdMatch Match

	if ms.isPlayerCountUneven(players) {
		return createdMatch, errors.NewUnevenMatchPlayersError()
	}
	// winnerPlayers := ms.getWinnerPlayers(players)
	// loserPlayers := ms.getLoserPlayers(players)

	// matchID, err := ms.matchRepository.Create()
	// if err != nil {
	// 	return createdMatch, err
	// }

	return createdMatch, nil
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
func NewService(matchRepository Repository) Service {
	return &matchService{
		matchRepository: matchRepository,
	}
}

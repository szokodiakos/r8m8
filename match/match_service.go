package match

import (
	"github.com/szokodiakos/r8m8/match/model"
	"github.com/szokodiakos/r8m8/transaction"
)

// Service interface
type Service interface {
	GetByID(tr transaction.Transaction, matchID int64) (model.Match, error)
}

type matchService struct {
	matchRepository       Repository
	matchPlayerRepository PlayerRepository
}

func (m *matchService) GetByID(tr transaction.Transaction, matchID int64) (model.Match, error) {
	match, err := m.matchRepository.GetByID(tr, matchID)
	if err != nil {
		return match, err
	}

	matchPlayers, err := m.matchPlayerRepository.GetMultipleByMatchID(tr, matchID)
	if err != nil {
		return match, err
	}

	winnerMatchPlayers, loserMatchPlayers := sortMatchPlayers(matchPlayers)
	match.WinnerMatchPlayers = winnerMatchPlayers
	match.LoserMatchPlayers = loserMatchPlayers

	return match, nil
}

func sortMatchPlayers(matchPlayers []model.MatchPlayer) ([]model.MatchPlayer, []model.MatchPlayer) {
	winnerMatchPlayers := []model.MatchPlayer{}
	loserMatchPlayers := []model.MatchPlayer{}
	for i := range matchPlayers {
		if matchPlayers[i].Details.HasWon == true {
			winnerMatchPlayers = append(winnerMatchPlayers, matchPlayers[i])
		} else {
			loserMatchPlayers = append(loserMatchPlayers, matchPlayers[i])
		}
	}
	return winnerMatchPlayers, loserMatchPlayers
}

// NewService creates a service
func NewService(
	matchRepository Repository,
	matchPlayerRepository PlayerRepository,
) Service {
	return &matchService{
		matchRepository:       matchRepository,
		matchPlayerRepository: matchPlayerRepository,
	}
}

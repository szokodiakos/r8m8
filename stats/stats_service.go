package stats

import (
	"github.com/szokodiakos/r8m8/league"
	"github.com/szokodiakos/r8m8/transaction"
)

// Service interface
type Service interface {
	GetLeaderboard(tr transaction.Transaction, league league.League) (Leaderboard, error)
}

type statsService struct {
	statsRepository Repository
}

func (s *statsService) GetLeaderboard(tr transaction.Transaction, league league.League) (Leaderboard, error) {
	leaderboard := Leaderboard{
		DisplayName: league.DisplayName,
	}

	leaderboardPlayers, err := s.statsRepository.GetLeaderboardPlayersByLeagueUniqueName(tr, league.UniqueName)
	if err != nil {
		return leaderboard, err
	}

	leaderboard.Players = leaderboardPlayers
	return leaderboard, nil
}

// NewService factory
func NewService(statsRepository Repository) Service {
	return &statsService{
		statsRepository: statsRepository,
	}
}

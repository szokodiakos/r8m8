package stats

import (
	"github.com/szokodiakos/r8m8/league"
	"github.com/szokodiakos/r8m8/transaction"
)

// Service interface
type Service interface {
	GetLeaderboard(tr transaction.Transaction, league league.League) (Leaderboard, error)
	GetMatchStats(tr transaction.Transaction, matchID int64) (MatchStats, error)
}

type statsService struct {
	playerStatsRepository PlayerRepository
}

func (s *statsService) GetLeaderboard(tr transaction.Transaction, league league.League) (Leaderboard, error) {
	leaderboard := Leaderboard{
		DisplayName: league.DisplayName,
	}

	repoPlayersStats, err := s.playerStatsRepository.GetMultipleByLeagueUniqueName(tr, league.UniqueName)
	if err != nil {
		return leaderboard, err
	}

	leaderboard.PlayersStats = repoPlayersStats
	return leaderboard, nil
}

func (s *statsService) GetMatchStats(tr transaction.Transaction, matchID int64) (MatchStats, error) {
	return MatchStats{}, nil
}

// NewService factory
func NewService(playerStatsRepository PlayerRepository) Service {
	return &statsService{
		playerStatsRepository: playerStatsRepository,
	}
}

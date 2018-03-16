package stats

import (
	"github.com/szokodiakos/r8m8/league"
	"github.com/szokodiakos/r8m8/player"
	"github.com/szokodiakos/r8m8/transaction"
)

// Service interface
type Service interface {
	GetLeaderboard(tr transaction.Transaction, league league.League) (Leaderboard, error)
	GetMatchStats(tr transaction.Transaction, matchID int64) (MatchStats, error)
}

type statsService struct {
	playerStatsRepository PlayerRepository
	playerRepository      player.Repository
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
	var matchStats MatchStats

	// repoPlayer, err := s.playerRepository.GetReporterPlayerByMatchID(tr, matchID)
	// if err != nil {
	// 	return matchStats, err
	// }

	// matchStats.ReporterDisplayName = repoPlayer.DisplayName

	// players, err := s.playerRepository.GetMultipleByMatchID(tr, matchID)
	// if err != nil {
	// 	return matchStats, err
	// }

	return matchStats, nil
}

// NewService factory
func NewService(playerStatsRepository PlayerRepository, playerRepository player.Repository) Service {
	return &statsService{
		playerStatsRepository: playerStatsRepository,
		playerRepository:      playerRepository,
	}
}

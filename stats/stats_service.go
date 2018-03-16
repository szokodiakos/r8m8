package stats

import (
	leagueModel "github.com/szokodiakos/r8m8/league/model"
	"github.com/szokodiakos/r8m8/player"
	"github.com/szokodiakos/r8m8/stats/model"
	"github.com/szokodiakos/r8m8/transaction"
)

// Service interface
type Service interface {
	GetLeaderboard(tr transaction.Transaction, league leagueModel.League) (model.Leaderboard, error)
	// GetMatchStats(tr transaction.Transaction, matchID int64) (MatchStats, error)
}

type statsService struct {
	playerStatsRepository PlayerRepository
	playerRepository      player.Repository
}

func (s *statsService) GetLeaderboard(tr transaction.Transaction, league leagueModel.League) (model.Leaderboard, error) {
	leaderboard := model.Leaderboard{
		DisplayName: league.DisplayName,
	}

	playersStats, err := s.playerStatsRepository.GetMultipleByLeagueUniqueName(tr, league.UniqueName)
	if err != nil {
		return leaderboard, err
	}

	leaderboard.PlayersStats = playersStats
	return leaderboard, nil
}

// func (s *statsService) GetMatchStats(tr transaction.Transaction, matchID int64) (MatchStats, error) {
// 	var matchStats MatchStats

// 	reporterPlayer, err := s.playerRepository.GetReporterPlayerByMatchID(tr, matchID)
// 	if err != nil {
// 		return matchStats, err
// 	}

// 	matchStats.ReporterDisplayName = reporterPlayer.DisplayName

// 	players, err := s.playerRepository.GetMultipleByMatchID(tr, matchID)
// 	if err != nil {
// 		return matchStats, err
// 	}

// 	return matchStats, nil
// }

// NewService factory
func NewService(playerStatsRepository PlayerRepository, playerRepository player.Repository) Service {
	return &statsService{
		playerStatsRepository: playerStatsRepository,
		playerRepository:      playerRepository,
	}
}

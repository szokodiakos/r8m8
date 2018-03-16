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
	GetMatchStats(tr transaction.Transaction, matchID int64) (model.MatchStats, error)
}

type statsService struct {
	playerStatsRepository      PlayerRepository
	playerRepository           player.Repository
	matchPlayerStatsRepository MatchPlayerStatsRepository
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

func (s *statsService) GetMatchStats(tr transaction.Transaction, matchID int64) (model.MatchStats, error) {
	var matchStats model.MatchStats

	reporterPlayer, err := s.playerRepository.GetReporterPlayerByMatchID(tr, matchID)
	if err != nil {
		return matchStats, err
	}

	matchStats.ReporterPlayer = reporterPlayer

	matchPlayersStats, err := s.matchPlayerStatsRepository.GetMultipleByMatchID(tr, matchID)
	if err != nil {
		return matchStats, err
	}

	winnerMatchPlayersStats, loserMatchPlayersStats := sortMatchPlayersStats(matchPlayersStats)
	matchStats.WinnerMatchPlayersStats = winnerMatchPlayersStats
	matchStats.LoserMatchPlayersStats = loserMatchPlayersStats

	return matchStats, nil
}

func sortMatchPlayersStats(matchPlayersStats []model.MatchPlayerStats) ([]model.MatchPlayerStats, []model.MatchPlayerStats) {
	winnerMatchPlayersStats := []model.MatchPlayerStats{}
	loserMatchPlayersStats := []model.MatchPlayerStats{}
	for i := range matchPlayersStats {
		if matchPlayersStats[i].Details.HasWon == true {
			winnerMatchPlayersStats = append(winnerMatchPlayersStats, matchPlayersStats[i])
		} else {
			loserMatchPlayersStats = append(loserMatchPlayersStats, matchPlayersStats[i])
		}
	}
	return winnerMatchPlayersStats, loserMatchPlayersStats
}

// NewService factory
func NewService(playerStatsRepository PlayerRepository, playerRepository player.Repository, matchPlayerStatsRepository MatchPlayerStatsRepository) Service {
	return &statsService{
		playerStatsRepository:      playerStatsRepository,
		playerRepository:           playerRepository,
		matchPlayerStatsRepository: matchPlayerStatsRepository,
	}
}

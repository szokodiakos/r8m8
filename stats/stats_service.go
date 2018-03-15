package stats

import (
	"github.com/szokodiakos/r8m8/league"
)

// Service interface
type Service interface {
	GetLeaderboard(league league.League) (Leaderboard, error)
}

type statsService struct {
	leagueRepository league.Repository
}

func (s *statsService) GetLeaderboard(league league.League) (Leaderboard, error) {
	leaderboard := Leaderboard{
		DisplayName: league.DisplayName,
	}

	// repoLeague, err := s.leagueRepository.GetByUniqueName(league.UniqueName)
	return leaderboard, nil
}

// NewService factory
func NewService(leagueRepository league.Repository) Service {
	return &statsService{
		leagueRepository: leagueRepository,
	}
}

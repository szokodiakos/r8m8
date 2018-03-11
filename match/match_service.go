package match

import (
	"github.com/szokodiakos/r8m8/league"
	"github.com/szokodiakos/r8m8/match/errors"
	"github.com/szokodiakos/r8m8/player"
	"github.com/szokodiakos/r8m8/rating"
	"github.com/szokodiakos/r8m8/transaction"
)

// Service interface
type Service interface {
	Add(transaction transaction.Transaction, players []player.Player, league league.League) error
}

type matchService struct {
	matchRepository Repository
	ratingService   rating.Service
	playerService   player.Service
	leagueService   league.Service
}

func (m *matchService) Add(transaction transaction.Transaction, players []player.Player, league league.League) error {
	if isPlayerCountUneven(players) {
		return errors.NewUnevenMatchPlayersError()
	}

	repoLeague, err := m.leagueService.GetOrAddLeague(transaction, league)
	if err != nil {
		return err
	}

	leagueID := repoLeague.ID
	repoPlayers, err := m.playerService.GetOrAddPlayers(transaction, players, leagueID)
	if err != nil {
		return err
	}

	matchID, err := m.matchRepository.Create(transaction, repoLeague.ID)
	if err != nil {
		return err
	}

	repoPlayerIDs := mapToIDs(repoPlayers)
	err = m.ratingService.UpdateRatings(transaction, repoPlayerIDs)
	if err != nil {
		return err
	}

	return nil
}

func isPlayerCountUneven(players []player.Player) bool {
	return (len(players) % 2) != 0
}

func mapToIDs(repoPlayers []player.RepoPlayer) []int64 {
	IDs := make([]int64, len(repoPlayers))
	for i := range repoPlayers {
		IDs[i] = repoPlayers[i].ID
	}
	return IDs
}

// NewService creates a service
func NewService(matchRepository Repository, ratingService rating.Service, playerService player.Service, leagueService league.Service) Service {
	return &matchService{
		matchRepository: matchRepository,
		ratingService:   ratingService,
		playerService:   playerService,
		leagueService:   leagueService,
	}
}

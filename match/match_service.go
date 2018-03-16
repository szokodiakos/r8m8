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
	Add(tr transaction.Transaction, players []player.Player, league league.League, reporterPlayer player.Player) error
}

type matchService struct {
	matchRepository Repository
	ratingService   rating.Service
	playerService   player.Service
	leagueService   league.Service
}

func (m *matchService) Add(tr transaction.Transaction, players []player.Player, league league.League, reporterPlayer player.Player) error {
	if isPlayerCountUneven(players) {
		return &errors.UnevenMatchPlayersError{}
	}

	if isReporterPlayerNotInLeague(reporterPlayer, players) {
		return &errors.ReporterPlayerNotInLeagueError{}
	}

	repoLeague, err := m.leagueService.GetOrAddLeague(tr, league)
	if err != nil {
		return err
	}

	leagueID := repoLeague.ID
	repoPlayers, err := m.playerService.GetOrAddPlayers(tr, players, leagueID)
	if err != nil {
		return err
	}

	reporterRepoPlayer := getReporterRepoPlayer(reporterPlayer, repoPlayers)
	reporterRepoPlayerID := reporterRepoPlayer.ID
	matchID, err := m.matchRepository.Create(tr, leagueID, reporterRepoPlayerID)
	if err != nil {
		return err
	}

	repoPlayerIDs := mapToIDs(repoPlayers)
	err = m.ratingService.UpdateRatings(tr, repoPlayerIDs, matchID)
	return err
}

func isPlayerCountUneven(players []player.Player) bool {
	return (len(players) % 2) != 0
}

func isReporterPlayerNotInLeague(reporterPlayer player.Player, players []player.Player) bool {
	missingFromLeague := true
	for i := range players {
		if players[i].UniqueName == reporterPlayer.UniqueName {
			missingFromLeague = false
		}
	}
	return missingFromLeague
}

func getReporterRepoPlayer(reporterPlayer player.Player, repoPlayers []player.Player) player.Player {
	var reporterRepoPlayer player.Player

	for i := range repoPlayers {
		if repoPlayers[i].UniqueName == reporterPlayer.UniqueName {
			reporterRepoPlayer = repoPlayers[i]
		}
	}

	return reporterRepoPlayer
}

func mapToIDs(players []player.Player) []int64 {
	IDs := make([]int64, len(players))
	for i := range players {
		IDs[i] = players[i].ID
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

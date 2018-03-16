package match

import (
	"github.com/szokodiakos/r8m8/league"
	leagueModel "github.com/szokodiakos/r8m8/league/model"
	"github.com/szokodiakos/r8m8/match/errors"
	"github.com/szokodiakos/r8m8/player"
	playerModel "github.com/szokodiakos/r8m8/player/model"
	"github.com/szokodiakos/r8m8/rating"
	"github.com/szokodiakos/r8m8/transaction"
)

// Service interface
type Service interface {
	Add(tr transaction.Transaction, players []playerModel.Player, league leagueModel.League, reporterPlayer playerModel.Player) error
}

type matchService struct {
	matchRepository Repository
	ratingService   rating.Service
	playerService   player.Service
	leagueService   league.Service
}

func (m *matchService) Add(tr transaction.Transaction, players []playerModel.Player, league leagueModel.League, reporterPlayer playerModel.Player) error {
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

func isPlayerCountUneven(players []playerModel.Player) bool {
	return (len(players) % 2) != 0
}

func isReporterPlayerNotInLeague(reporterPlayer playerModel.Player, players []playerModel.Player) bool {
	missingFromLeague := true
	for i := range players {
		if players[i].UniqueName == reporterPlayer.UniqueName {
			missingFromLeague = false
		}
	}
	return missingFromLeague
}

func getReporterRepoPlayer(reporterPlayer playerModel.Player, repoPlayers []playerModel.Player) playerModel.Player {
	var reporterRepoPlayer playerModel.Player

	for i := range repoPlayers {
		if repoPlayers[i].UniqueName == reporterPlayer.UniqueName {
			reporterRepoPlayer = repoPlayers[i]
		}
	}

	return reporterRepoPlayer
}

func mapToIDs(players []playerModel.Player) []int64 {
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

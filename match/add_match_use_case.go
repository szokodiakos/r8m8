package match

import (
	"github.com/szokodiakos/r8m8/league"
	"github.com/szokodiakos/r8m8/match/errors"
	"github.com/szokodiakos/r8m8/match/model"
	"github.com/szokodiakos/r8m8/player"
	playerModel "github.com/szokodiakos/r8m8/player/model"
	"github.com/szokodiakos/r8m8/rating"
	"github.com/szokodiakos/r8m8/transaction"
)

// AddMatchUseCase interface
type AddMatchUseCase interface {
	Handle(input model.AddMatchInput) (model.AddMatchOutput, error)
}

type addMatchUseCase struct {
	transactionService transaction.Service
	matchRepository    Repository
	ratingService      rating.Service
	playerService      player.Service
	leagueService      league.Service
	matchService       Service
}

func (a *addMatchUseCase) Handle(input model.AddMatchInput) (model.AddMatchOutput, error) {
	var output model.AddMatchOutput

	if isPlayerCountUneven(input.Players) {
		return output, &errors.UnevenMatchPlayersError{}
	}

	if isReporterPlayerNotInLeague(input.ReporterPlayer, input.Players) {
		return output, &errors.ReporterPlayerNotInLeagueError{}
	}

	tr, err := a.transactionService.Start()
	if err != nil {
		return output, err
	}

	repoLeague, err := a.leagueService.GetOrAddLeague(tr, input.League)
	if err != nil {
		return output, a.transactionService.Rollback(tr, err)
	}

	leagueID := repoLeague.ID
	repoPlayers, err := a.playerService.GetOrAddPlayersByLeagueID(tr, input.Players, leagueID)
	if err != nil {
		return output, a.transactionService.Rollback(tr, err)
	}

	reporterRepoPlayer := getReporterRepoPlayer(input.ReporterPlayer, repoPlayers)
	reporterRepoPlayerID := reporterRepoPlayer.ID
	matchID, err := a.matchRepository.Create(tr, leagueID, reporterRepoPlayerID)
	if err != nil {
		return output, a.transactionService.Rollback(tr, err)
	}

	repoPlayerIDs := mapToIDs(repoPlayers)
	err = a.ratingService.UpdateRatings(tr, repoPlayerIDs, matchID)
	if err != nil {
		return output, a.transactionService.Rollback(tr, err)
	}

	match, err := a.matchService.GetByID(tr, matchID)
	if err != nil {
		return output, a.transactionService.Rollback(tr, err)
	}

	output = model.AddMatchOutput{
		Match: match,
	}
	err = a.transactionService.Commit(tr)
	return output, err
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

// NewAddMatchUseCase factory
func NewAddMatchUseCase(
	transactionService transaction.Service,
	matchRepository Repository,
	ratingService rating.Service,
	playerService player.Service,
	leagueService league.Service,
	matchService Service,
) AddMatchUseCase {
	return &addMatchUseCase{
		transactionService: transactionService,
		matchRepository:    matchRepository,
		ratingService:      ratingService,
		playerService:      playerService,
		leagueService:      leagueService,
		matchService:       matchService,
	}
}

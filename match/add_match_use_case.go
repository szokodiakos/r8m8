package match

import (
	"github.com/szokodiakos/r8m8/league"
	"github.com/szokodiakos/r8m8/match/errors"
	"github.com/szokodiakos/r8m8/match/model"
	"github.com/szokodiakos/r8m8/player"
	playerErrors "github.com/szokodiakos/r8m8/player/errors"
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
	leagueService      league.Service
	matchService       Service
	playerRepository   player.Repository
}

func (a *addMatchUseCase) Handle(input model.AddMatchInput) (model.AddMatchOutput, error) {
	var output model.AddMatchOutput

	if isPlayerCountUneven(input.Players) {
		return output, &errors.UnevenMatchPlayersError{}
	}

	tr, err := a.transactionService.Start()
	if err != nil {
		return output, err
	}

	repoLeague, err := a.leagueService.GetOrAdd(tr, input.League)
	if err != nil {
		return output, a.transactionService.Rollback(tr, err)
	}

	err = a.leagueService.AddAnyMissingPlayers(tr, repoLeague, input.Players)
	if err != nil {
		return output, a.transactionService.Rollback(tr, err)
	}

	repoReporterPlayer, err := a.playerRepository.GetByUniqueName(tr, input.ReporterPlayer.UniqueName)
	if err != nil {
		switch err.(type) {
		case *playerErrors.PlayerNotFoundError:
			return output, a.transactionService.Rollback(tr, &errors.ReporterPlayerNotInLeagueError{})
		default:
			return output, a.transactionService.Rollback(tr, err)
		}
	}

	match := model.Match{
		League:         repoLeague,
		ReporterPlayer: repoReporterPlayer,
	}
	repoMatch, err := a.matchService.CreateWithPlayers(tr, match, input.Players)
	if err != nil {
		return output, a.transactionService.Rollback(tr, err)
	}

	output = model.AddMatchOutput{
		Match: repoMatch,
	}
	err = a.transactionService.Commit(tr)
	return output, err
}

func isPlayerCountUneven(players []playerModel.Player) bool {
	return (len(players) % 2) != 0
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

// NewAddMatchUseCase factory
func NewAddMatchUseCase(
	transactionService transaction.Service,
	matchRepository Repository,
	ratingService rating.Service,
	leagueService league.Service,
	matchService Service,
	playerRepository player.Repository,
) AddMatchUseCase {
	return &addMatchUseCase{
		transactionService: transactionService,
		matchRepository:    matchRepository,
		ratingService:      ratingService,
		leagueService:      leagueService,
		matchService:       matchService,
		playerRepository:   playerRepository,
	}
}

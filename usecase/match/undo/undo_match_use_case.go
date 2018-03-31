package undo

import (
	"github.com/szokodiakos/r8m8/league"
	"github.com/szokodiakos/r8m8/match"
	"github.com/szokodiakos/r8m8/transaction"
)

// UseCase interface
type UseCase interface {
	Handle(input Input) (Output, error)
}

type undoMatchUseCase struct {
	transactionService  transaction.Service
	leaguePlayerService league.PlayerService
	matchRepository     match.Repository
	leagueRepository    league.Repository
}

func (u *undoMatchUseCase) Handle(input Input) (output Output, err error) {
	tr, err := u.transactionService.Start()
	if err != nil {
		return
	}

	defer func() {
		switch err {
		case nil:
			err = u.transactionService.Commit(tr)
			return
		default:
			err = u.transactionService.Rollback(tr, err)
			return
		}
	}()

	repoMatch, err := u.matchRepository.GetLatestByReporterPlayerID(tr, input.ReporterPlayer.ID)
	if err != nil {
		return
	}

	repoLeague, err := u.leagueRepository.GetByID(tr, repoMatch.LeagueID)
	if err != nil {
		return
	}

	adjustedLeaguePlayers := u.leaguePlayerService.UndoRatingChangesForLeaguePlayers(repoLeague.LeaguePlayers, repoMatch.MatchPlayers)
	repoLeague.LeaguePlayers = adjustedLeaguePlayers

	err = u.leagueRepository.Update(tr, repoLeague)
	if err != nil {
		return
	}

	err = u.matchRepository.Remove(tr, repoMatch)
	if err != nil {
		return
	}

	output = Output{
		ReporterPlayer: input.ReporterPlayer,
		LeaguePlayers:  adjustedLeaguePlayers,
		MatchPlayers:   repoMatch.MatchPlayers,
	}

	return
}

// NewUndoMatchUseCase factory
func NewUndoMatchUseCase(transactionService transaction.Service, leaguePlayerService league.PlayerService, matchRepository match.Repository, leagueRepository league.Repository) UseCase {
	return &undoMatchUseCase{
		transactionService:  transactionService,
		leaguePlayerService: leaguePlayerService,
		matchRepository:     matchRepository,
		leagueRepository:    leagueRepository,
	}
}

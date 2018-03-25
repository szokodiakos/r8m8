package undo

import (
	"github.com/szokodiakos/r8m8/entity"
	"github.com/szokodiakos/r8m8/transaction"
)

// UndoMatchUseCase interface
type UndoMatchUseCase interface {
	Handle(input UndoMatchInput) (UndoMatchOutput, error)
}

type undoMatchUseCase struct {
	transactionService transaction.Service
	matchRepository    entity.MatchRepository
	leagueRepository   entity.LeagueRepository
}

func (u *undoMatchUseCase) Handle(input UndoMatchInput) (UndoMatchOutput, error) {
	var output UndoMatchOutput

	tr, err := u.transactionService.Start()
	if err != nil {
		return output, err
	}

	repoMatch, err := u.matchRepository.GetLatestByReporterPlayerID(tr, input.ReporterPlayer.ID)
	if err != nil {
		return output, u.transactionService.Rollback(tr, err)
	}

	repoLeague, err := u.leagueRepository.GetByID(tr, repoMatch.LeagueID)
	if err != nil {
		return output, u.transactionService.Rollback(tr, err)
	}

	adjustedLeaguePlayers := undoRatingChangesForLeaguePlayers(repoLeague.LeaguePlayers, repoMatch.MatchPlayers)
	repoLeague.LeaguePlayers = adjustedLeaguePlayers

	err = u.leagueRepository.Update(tr, repoLeague)
	if err != nil {
		return output, u.transactionService.Rollback(tr, err)
	}

	err = u.matchRepository.Remove(tr, repoMatch)
	if err != nil {
		return output, u.transactionService.Rollback(tr, err)
	}

	err = u.transactionService.Commit(tr)
	output = UndoMatchOutput{
		ReporterPlayer: input.ReporterPlayer,
		LeaguePlayers:  adjustedLeaguePlayers,
		MatchPlayers:   repoMatch.MatchPlayers,
	}

	return output, err
}

func undoRatingChangesForLeaguePlayers(leaguePlayers []entity.LeaguePlayer, matchPlayers []entity.MatchPlayer) []entity.LeaguePlayer {
	adjustedLeaguePlayers := make([]entity.LeaguePlayer, len(leaguePlayers))
	copy(adjustedLeaguePlayers, leaguePlayers)

	for i := range matchPlayers {
		for j := range leaguePlayers {
			if isLeaguePlayerParticipatedInMatch(matchPlayers[i], leaguePlayers[j]) {
				adjustedLeaguePlayers[j].Rating -= matchPlayers[i].RatingChange
			}
		}
	}

	return adjustedLeaguePlayers
}

func isLeaguePlayerParticipatedInMatch(matchPlayer entity.MatchPlayer, leaguePlayer entity.LeaguePlayer) bool {
	return matchPlayer.PlayerID == leaguePlayer.PlayerID
}

// NewUndoMatchUseCase factory
func NewUndoMatchUseCase(transactionService transaction.Service, matchRepository entity.MatchRepository, leagueRepository entity.LeagueRepository) UndoMatchUseCase {
	return &undoMatchUseCase{
		transactionService: transactionService,
		matchRepository:    matchRepository,
		leagueRepository:   leagueRepository,
	}
}

package match

import (
	"github.com/szokodiakos/r8m8/match/model"
	"github.com/szokodiakos/r8m8/transaction"
)

// PlayerRepository interface
type PlayerRepository interface {
	GetMultipleByMatchID(tr transaction.Transaction, matchID int64) ([]model.MatchPlayer, error)
}

package player

import (
	"github.com/szokodiakos/r8m8/player/model"
	"github.com/szokodiakos/r8m8/transaction"
)

// Repository interface
type Repository interface {
	Create(tr transaction.Transaction, player model.Player) (int64, error)
	GetMultipleByUniqueNames(tr transaction.Transaction, uniqueNames []string) ([]model.Player, error)
	GetReporterPlayerByMatchID(tr transaction.Transaction, matchID int64) (model.Player, error)
}

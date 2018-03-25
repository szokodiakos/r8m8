package entity

import (
	"github.com/szokodiakos/r8m8/transaction"
)

// PlayerRepository interface
type PlayerRepository interface {
	Add(tr transaction.Transaction, player Player) (Player, error)
	GetMultipleByIDs(tr transaction.Transaction, ids []string) ([]Player, error)
	GetByID(tr transaction.Transaction, id string) (Player, error)
}

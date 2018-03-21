package player

import (
	"github.com/szokodiakos/r8m8/player/model"
	"github.com/szokodiakos/r8m8/transaction"
)

// Repository interface
type Repository interface {
	Create(tr transaction.Transaction, player model.Player) (model.Player, error)
	GetMultipleByUniqueNames(tr transaction.Transaction, uniqueNames []string) ([]model.Player, error)
	GetByUniqueName(tr transaction.Transaction, uniqueName string) (model.Player, error)
}

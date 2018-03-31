package player

import (
	"github.com/szokodiakos/r8m8/entity"
	"github.com/szokodiakos/r8m8/transaction"
)

// PlayerRepository interface
type PlayerRepository interface {
	Add(tr transaction.Transaction, player entity.Player) (entity.Player, error)
	GetMultipleByIDs(tr transaction.Transaction, ids []string) ([]entity.Player, error)
	GetByID(tr transaction.Transaction, id string) (entity.Player, error)
}

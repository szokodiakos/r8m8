package league

import (
	"github.com/szokodiakos/r8m8/league/model"
	"github.com/szokodiakos/r8m8/transaction"
)

// Repository interface
type Repository interface {
	GetByUniqueName(tr transaction.Transaction, uniqueName string) (model.League, error)
	Create(tr transaction.Transaction, league model.League) (model.League, error)
}

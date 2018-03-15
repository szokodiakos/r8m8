package league

import (
	"github.com/szokodiakos/r8m8/transaction"
)

// Repository interface
type Repository interface {
	GetByUniqueName(tr transaction.Transaction, uniqueName string) (RepoLeague, error)
	Create(tr transaction.Transaction, league League) error
}

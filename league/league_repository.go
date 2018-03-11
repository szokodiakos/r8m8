package league

import (
	"github.com/szokodiakos/r8m8/transaction"
)

// Repository interface
type Repository interface {
	GetByUniqueName(transaction transaction.Transaction, uniqueName string) (RepoLeague, error)
	Create(transaction transaction.Transaction, league League) error
}

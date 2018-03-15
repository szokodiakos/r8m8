package details

import "github.com/szokodiakos/r8m8/transaction"

// Repository interface
type Repository interface {
	Create(tr transaction.Transaction, details Details) error
}

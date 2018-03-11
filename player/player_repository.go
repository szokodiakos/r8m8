package player

import "github.com/szokodiakos/r8m8/transaction"

// Repository interface
type Repository interface {
	Create(transaction transaction.Transaction, player Player) (int64, error)
	GetMultipleByUniqueNames(transaction transaction.Transaction, uniqueNames []string) ([]RepoPlayer, error)
}

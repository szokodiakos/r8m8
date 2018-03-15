package player

import "github.com/szokodiakos/r8m8/transaction"

// Repository interface
type Repository interface {
	Create(tr transaction.Transaction, player Player) (int64, error)
	GetMultipleByUniqueNames(tr transaction.Transaction, uniqueNames []string) ([]RepoPlayer, error)
}

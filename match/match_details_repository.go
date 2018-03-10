package match

import "github.com/szokodiakos/r8m8/transaction"

// DetailsRepository interface
type DetailsRepository interface {
	Create(transaction transaction.Transaction, matchDetails Details) error
}

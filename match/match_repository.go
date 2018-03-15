package match

import "github.com/szokodiakos/r8m8/transaction"

// Repository interface
type Repository interface {
	Create(tr transaction.Transaction, leagueID int64, reporterPlayerID int64) (int64, error)
}

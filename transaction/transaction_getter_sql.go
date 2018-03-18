package transaction

import "github.com/szokodiakos/r8m8/sql"

// GetSQLTransaction func
func GetSQLTransaction(transaction Transaction) sql.Transaction {
	return transaction.concreteTransaction
}

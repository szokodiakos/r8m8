package transaction

import "github.com/szokodiakos/r8m8/sql"

// Transaction struct
type Transaction struct {
	concreteTransaction sql.Transaction
}

package match

import (
	"github.com/szokodiakos/r8m8/sql"
	"github.com/szokodiakos/r8m8/transaction"
)

type matchRepositorySQL struct {
}

func (mrs *matchRepositorySQL) Create(transaction transaction.Transaction, leagueID int64) (int64, error) {
	var createdID int64

	query := `
		INSERT INTO matches
			(league_id, created_at)
		VALUES
			($1, current_timestamp)
		RETURNING id;
	`

	sqlTransaction := transaction.ConcreteTransaction.(sql.Transaction)
	err := sqlTransaction.Get(&createdID, query, leagueID)
	return createdID, err
}

// NewRepositorySQL factory
func NewRepositorySQL() Repository {
	return &matchRepositorySQL{}
}

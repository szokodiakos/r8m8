package match

import (
	"github.com/szokodiakos/r8m8/transaction"
)

type matchRepositorySQL struct {
}

func (m *matchRepositorySQL) Create(tr transaction.Transaction, leagueID int64, reporterPlayerID int64) (int64, error) {
	var createdID int64

	query := `
		INSERT INTO matches
			(league_id, reporter_player_id, created_at)
		VALUES
			($1, $2, current_timestamp)
		RETURNING id;
	`

	sqlTransaction := transaction.GetSQLTransaction(tr)
	err := sqlTransaction.Get(&createdID, query, leagueID, reporterPlayerID)
	return createdID, err
}

// NewRepositorySQL factory
func NewRepositorySQL() Repository {
	return &matchRepositorySQL{}
}

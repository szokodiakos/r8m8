package match

import (
	"github.com/szokodiakos/r8m8/sql"
	"github.com/szokodiakos/r8m8/transaction"
)

type matchDetailsRepositorySQL struct {
}

func (mdrs *matchDetailsRepositorySQL) Create(transaction transaction.Transaction, matchDetails Details) error {
	query := `
		INSERT INTO match_details
			(player_id, match_id, rating_change)
		VALUES
			($1, $2, $3);
	`

	sqlTransaction := transaction.ConcreteTransaction.(sql.Transaction)
	_, err := sqlTransaction.Exec(query, matchDetails.PlayerID, matchDetails.MatchID, matchDetails.RatingChange)
	return err
}

// NewDetailsRepositorySQL factory
func NewDetailsRepositorySQL() DetailsRepository {
	return &matchDetailsRepositorySQL{}
}

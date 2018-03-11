package details

import (
	"github.com/szokodiakos/r8m8/sql"
	"github.com/szokodiakos/r8m8/transaction"
)

type detailsRepositorySQL struct {
}

func (mdrs *detailsRepositorySQL) Create(transaction transaction.Transaction, details Details) error {
	query := `
		INSERT INTO details
			(player_id, match_id, rating_change, has_won)
		VALUES
			($1, $2, $3, $4);
	`

	sqlTransaction := transaction.ConcreteTransaction.(sql.Transaction)
	_, err := sqlTransaction.Exec(query, details.PlayerID, details.MatchID, details.RatingChange, details.HasWon)
	return err
}

// NewRepositorySQL factory
func NewRepositorySQL() Repository {
	return &detailsRepositorySQL{}
}

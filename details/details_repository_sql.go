package details

import (
	"github.com/szokodiakos/r8m8/details/model"
	"github.com/szokodiakos/r8m8/transaction"
)

type detailsRepositorySQL struct {
}

func (d *detailsRepositorySQL) Create(tr transaction.Transaction, details model.Details) error {
	query := `
		INSERT INTO details
			(player_id, match_id, rating_change, has_won)
		VALUES
			($1, $2, $3, $4);
	`

	sqlTransaction := transaction.GetSQLTransaction(tr)
	_, err := sqlTransaction.Exec(query, details.PlayerID, details.MatchID, details.RatingChange, details.HasWon)
	return err
}

// NewRepositorySQL factory
func NewRepositorySQL() Repository {
	return &detailsRepositorySQL{}
}

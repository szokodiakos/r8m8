package rating

import (
	"github.com/lib/pq"
	"github.com/szokodiakos/r8m8/rating/model"
	"github.com/szokodiakos/r8m8/transaction"
)

type ratingRepositorySQL struct {
}

func (r *ratingRepositorySQL) Create(tr transaction.Transaction, rating model.Rating) error {
	query := `
		INSERT INTO ratings
			(player_id, league_id, rating)
		VALUES
			($1, $2, $3);
	`

	sqlTransaction := transaction.GetSQLTransaction(tr)
	_, err := sqlTransaction.Exec(query, rating.PlayerID, rating.LeagueID, rating.Rating)
	return err
}

func (r *ratingRepositorySQL) GetMultipleByPlayerIDs(tr transaction.Transaction, playerIDs []int64) ([]model.Rating, error) {
	ratings := []model.Rating{}

	query := `
		SELECT
			r.player_id,
			r.league_id,
			r.rating
		FROM
			ratings r
		WHERE
			r.player_id = ANY($1);
	`

	sqlTransaction := transaction.GetSQLTransaction(tr)
	err := sqlTransaction.Select(&ratings, query, pq.Array(playerIDs))

	return ratings, err
}

func (r *ratingRepositorySQL) UpdateRating(tr transaction.Transaction, rating model.Rating) error {
	query := `
		UPDATE ratings
		SET
			rating = $1
		WHERE
			player_id = $2 AND
			league_id = $3;
	`

	sqlTransaction := transaction.GetSQLTransaction(tr)
	_, err := sqlTransaction.Exec(query, rating.Rating, rating.PlayerID, rating.LeagueID)
	return err
}

// NewRepositorySQL factory
func NewRepositorySQL() Repository {
	return &ratingRepositorySQL{}
}

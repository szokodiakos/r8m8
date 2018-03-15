package rating

import (
	"github.com/lib/pq"
	"github.com/szokodiakos/r8m8/transaction"
)

type ratingRepositorySQL struct {
}

func (r *ratingRepositorySQL) Create(tr transaction.Transaction, rating Rating) error {
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

func (r *ratingRepositorySQL) GetMultipleByPlayerIDs(tr transaction.Transaction, playerIDs []int64) ([]RepoRating, error) {
	repoRatings := []RepoRating{}

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
	err := sqlTransaction.Select(&repoRatings, query, pq.Array(playerIDs))

	return repoRatings, err
}

func (r *ratingRepositorySQL) UpdateRating(tr transaction.Transaction, rating Rating) error {
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

package rating

import (
	"github.com/lib/pq"
	"github.com/szokodiakos/r8m8/sql"
	"github.com/szokodiakos/r8m8/transaction"
)

type ratingRepositorySQL struct {
}

func (r *ratingRepositorySQL) Create(transaction transaction.Transaction, rating Rating) error {
	query := `
		INSERT INTO ratings
			(player_id, league_id, rating)
		VALUES
			($1, $2, $3);
	`

	sqlTransaction := transaction.ConcreteTransaction.(sql.Transaction)
	_, err := sqlTransaction.Exec(query, rating.PlayerID, rating.LeagueID, rating.Rating)
	return err
}

func (r *ratingRepositorySQL) GetMultipleByPlayerIDs(transaction transaction.Transaction, playerIDs []int64) ([]RepoRating, error) {
	var repoRatings []RepoRating

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

	sqlTransaction := transaction.ConcreteTransaction.(sql.Transaction)
	rows, err := sqlTransaction.Query(query, pq.Array(playerIDs))
	if err != nil {
		return repoRatings, err
	}

	for rows.Next() {
		var playerID, leagueID int64
		var rating int

		if err := rows.Scan(&playerID, &leagueID, &rating); err != nil {
			return repoRatings, err
		}

		repoRating := RepoRating{
			PlayerID: playerID,
			LeagueID: leagueID,
			Rating:   rating,
		}
		repoRatings = append(repoRatings, repoRating)
	}

	return repoRatings, nil
}

func (r *ratingRepositorySQL) UpdateRating(transaction transaction.Transaction, rating Rating) error {
	query := `
		UPDATE ratings
		SET
			rating = $1
		WHERE
			player_id = $2 AND
			league_id = $3;
	`

	sqlTransaction := transaction.ConcreteTransaction.(sql.Transaction)
	_, err := sqlTransaction.Exec(query, rating.Rating, rating.PlayerID, rating.LeagueID)
	return err
}

// NewRatingRepositorySQL factory
func NewRatingRepositorySQL() Repository {
	return &ratingRepositorySQL{}
}

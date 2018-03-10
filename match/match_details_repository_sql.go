package match

import "github.com/szokodiakos/r8m8/sql"

type matchDetailsRepositorySQL struct {
	db sql.DB
}

func (mdrs *matchDetailsRepositorySQL) Create(matchDetails Details) error {
	query := `
		INSERT INTO match_details
			(player_id, match_id, rating_change)
		VALUES
			($1, $2, $3);
	`

	_, err := mdrs.db.Exec(query, matchDetails.PlayerID, matchDetails.MatchID, matchDetails.RatingChange)
	return err
}

// NewDetailsRepositorySQL factory
func NewDetailsRepositorySQL(db sql.DB) DetailsRepository {
	return &matchDetailsRepositorySQL{
		db: db,
	}
}

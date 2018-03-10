package match

import "database/sql"

type matchDetailsRepositorySQL struct {
	db *sql.DB
}

func (mdrs *matchDetailsRepositorySQL) Create(matchDetails Details) error {
	query := `
		INSERT INTO match_details
			(player_id, match_id, rating_change)
		VALUES
			(?, ?, ?);
	`

	_, err := mdrs.db.Exec(query, matchDetails.PlayerID, matchDetails.MatchID, matchDetails.RatingChange)
	return err
}

// NewDetailsRepositorySQL factory
func NewDetailsRepositorySQL(db *sql.DB) DetailsRepository {
	return &matchDetailsRepositorySQL{}
}

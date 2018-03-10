package match

import "github.com/szokodiakos/r8m8/sql"

type matchRepositorySQL struct {
	db sql.DB
}

func (mrs *matchRepositorySQL) Create() (int64, error) {
	var createdID int64
	query := `
		INSERT INTO matches
			(created_at)
		VALUES
			(utc_timestamp())
		RETURNING id;
	`

	res := mrs.db.QueryRow(query)
	err := res.Scan(&createdID)
	if err != nil {
		return createdID, err
	}

	return createdID, nil
}

// NewRepositorySQL factory
func NewRepositorySQL(db sql.DB) Repository {
	return &matchRepositorySQL{
		db: db,
	}
}

package match

import (
	"database/sql"
)

type matchRepositorySQL struct {
	db *sql.DB
}

func (mrs *matchRepositorySQL) Create() (int64, error) {
	var createdID int64
	query := `
		INSERT INTO matches
			(created_at)
		VALUES
			(utc_timestamp());
	`

	res, err := mrs.db.Exec(query)
	if err != nil {
		return createdID, err
	}

	createdID, err = res.LastInsertId()
	if err != nil {
		return createdID, err
	}

	return createdID, nil
}

// NewRepositorySQL factory
func NewRepositorySQL(db *sql.DB) Repository {
	return &matchRepositorySQL{
		db: db,
	}
}

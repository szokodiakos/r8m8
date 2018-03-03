package match

import (
	"database/sql"
)

type matchRepositorySQL struct {
	db *sql.DB
}

func (mrs *matchRepositorySQL) Create() error {
	query := `
		INSERT INTO matches
			(created_at)
		VALUES
			(utc_timestamp());
	`
	_, err := mrs.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

// NewRepositorySQL factory
func NewRepositorySQL(db *sql.DB) Repository {
	return &matchRepositorySQL{
		db: db,
	}
}

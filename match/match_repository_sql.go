package match

import (
	"database/sql"
)

type matchRepositorySQL struct {
	db *sql.DB
}

func (mrs *matchRepositorySQL) Create() error {
	_, err := mrs.db.Exec("INSERT INTO matches (created_at) VALUES (utc_timestamp())")
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

package match

import (
	"database/sql"
	"fmt"
)

type matchRepositorySQL struct {
	db *sql.DB
}

func (mrs *matchRepositorySQL) Create() error {
	result, err := mrs.db.Exec("INSERT INTO matches (created_at) VALUES (utc_timestamp())")
	if err != nil {
		return err
	}
	fmt.Println("res", result)
	return nil
}

// NewRepository factory
func NewRepository(db *sql.DB) Repository {
	return &matchRepositorySQL{
		db: db,
	}
}

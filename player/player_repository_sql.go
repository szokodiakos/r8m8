package player

import (
	"database/sql"
)

type playerRepositorySQL struct {
	db *sql.DB
}

func (prs *playerRepositorySQL) Create() (int64, error) {
	var createdID int64

	query := `
		INSERT INTO players DEFAULT VALUES;
	`

	res, err := prs.db.Exec(query)
	if err != nil {
		return createdID, err
	}

	createdID, err = res.LastInsertId()
	if err != nil {
		return createdID, err
	}

	return createdID, nil
}

// NewRepository factory
func NewRepository(db *sql.DB) Repository {
	return &playerRepositorySQL{
		db: db,
	}
}

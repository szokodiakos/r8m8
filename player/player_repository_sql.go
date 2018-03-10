package player

import "github.com/szokodiakos/r8m8/sql"

type playerRepositorySQL struct {
	db sql.DB
}

func (prs *playerRepositorySQL) Create() (int64, error) {
	var createdID int64

	query := `
		INSERT INTO players DEFAULT VALUES RETURNING id;
	`

	res := prs.db.QueryRow(query)
	err := res.Scan(&createdID)
	if err != nil {
		return createdID, err
	}

	return createdID, nil
}

func (prs *playerRepositorySQL) UpdateRatingByID(ID int64, rating int) error {
	query := `
		UPDATE players
		SET rating = ?
		WHERE id = ?
	`

	_, err := prs.db.Exec(query, ID, rating)
	return err
}

// NewRepository factory
func NewRepository(db sql.DB) Repository {
	return &playerRepositorySQL{
		db: db,
	}
}

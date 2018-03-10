package player

import (
	"github.com/szokodiakos/r8m8/sql"
	"github.com/szokodiakos/r8m8/transaction"
)

type playerRepositorySQL struct {
}

func (prs *playerRepositorySQL) Create(transaction transaction.Transaction) (int64, error) {
	var createdID int64

	query := `
		INSERT INTO players DEFAULT VALUES RETURNING id;
	`

	sqlTransaction := transaction.ConcreteTransaction.(sql.Transaction)
	res := sqlTransaction.QueryRow(query)
	err := res.Scan(&createdID)
	if err != nil {
		return createdID, err
	}

	return createdID, nil
}

func (prs *playerRepositorySQL) UpdateRatingByID(transaction transaction.Transaction, ID int64, rating int) error {
	query := `
		UPDATE players
		SET rating = $1
		WHERE id = $2
	`

	sqlTransaction := transaction.ConcreteTransaction.(sql.Transaction)
	_, err := sqlTransaction.Exec(query, ID, rating)
	return err
}

// NewRepository factory
func NewRepository() Repository {
	return &playerRepositorySQL{}
}

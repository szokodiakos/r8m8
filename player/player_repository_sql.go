package player

import (
	"github.com/lib/pq"
	"github.com/szokodiakos/r8m8/transaction"
)

type playerRepositorySQL struct {
}

func (p *playerRepositorySQL) GetMultipleByUniqueNames(tr transaction.Transaction, uniqueNames []string) ([]RepoPlayer, error) {
	repoPlayers := []RepoPlayer{}

	query := `
		SELECT
			p.id,
			p.unique_name,
			p.display_name
		FROM
			players p
		WHERE
			p.unique_name = ANY($1);
	`

	sqlTransaction := transaction.GetSQLTransaction(tr)
	err := sqlTransaction.Select(&repoPlayers, query, pq.Array(uniqueNames))
	return repoPlayers, err
}

func (p *playerRepositorySQL) Create(tr transaction.Transaction, player Player) (int64, error) {
	var createdID int64

	query := `
		INSERT INTO players
			(unique_name, display_name)
		VALUES
			($1, $2)
		RETURNING id;
	`

	sqlTransaction := transaction.GetSQLTransaction(tr)
	err := sqlTransaction.Get(&createdID, query, player.UniqueName, player.DisplayName)
	return createdID, err
}

// NewRepository factory
func NewRepository() Repository {
	return &playerRepositorySQL{}
}

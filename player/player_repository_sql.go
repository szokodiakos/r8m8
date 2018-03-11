package player

import (
	"github.com/lib/pq"
	"github.com/szokodiakos/r8m8/sql"
	"github.com/szokodiakos/r8m8/transaction"
)

type playerRepositorySQL struct {
}

func (p *playerRepositorySQL) GetMultipleByUniqueNames(transaction transaction.Transaction, uniqueNames []string) ([]RepoPlayer, error) {
	var repoPlayers = make([]RepoPlayer, 0, len(uniqueNames))
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

	sqlTransaction := transaction.ConcreteTransaction.(sql.Transaction)
	rows, err := sqlTransaction.Query(query, pq.Array(uniqueNames))
	if err != nil {
		return repoPlayers, err
	}

	for rows.Next() {
		var id int64
		var uniqueName, displayName string

		if err := rows.Scan(&id, &uniqueName, &displayName); err != nil {
			return repoPlayers, err
		}

		repoPlayer := RepoPlayer{
			ID:          id,
			UniqueName:  uniqueName,
			DisplayName: displayName,
		}
		repoPlayers = append(repoPlayers, repoPlayer)
	}
	return repoPlayers, nil
}

func (p *playerRepositorySQL) Create(transaction transaction.Transaction, player Player) (int64, error) {
	var createdID int64
	query := `
		INSERT INTO players
			(unique_name, display_name)
		VALUES
			($1, $2)
		RETURNING id;
	`

	sqlTransaction := transaction.ConcreteTransaction.(sql.Transaction)
	res := sqlTransaction.QueryRow(query)
	err := res.Scan(&createdID)
	if err != nil {
		return createdID, err
	}

	return createdID, nil
}

// NewRepository factory
func NewRepository() Repository {
	return &playerRepositorySQL{}
}

package player

import (
	"github.com/lib/pq"
	"github.com/szokodiakos/r8m8/sql"
	"github.com/szokodiakos/r8m8/transaction"
)

type playerRepositorySQL struct {
}

func (p *playerRepositorySQL) GetMultipleByUniqueName(transaction transaction.Transaction, uniqueNames []string) ([]RepoPlayer, error) {
	var repoPlayers = make([]RepoPlayer, 0, len(uniqueNames))
	query := `
		SELECT
			p.id,
			p.rating,
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
		var rating int
		var uniqueName, displayName string

		if err := rows.Scan(&id, &rating, &uniqueName, &displayName); err != nil {
			return repoPlayers, err
		}

		repoPlayer := RepoPlayer{
			ID:          id,
			Rating:      rating,
			UniqueName:  uniqueName,
			DisplayName: displayName,
		}
		repoPlayers = append(repoPlayers, repoPlayer)
	}
	return repoPlayers, nil
}

func (p *playerRepositorySQL) Create(transaction transaction.Transaction, player Player) error {
	query := `
		INSERT INTO players
			(unique_name, display_name)
		VALUES
			($1, $2);
	`

	sqlTransaction := transaction.ConcreteTransaction.(sql.Transaction)
	_, err := sqlTransaction.Exec(query, player.UniqueName, player.DisplayName)
	return err
}

func (p *playerRepositorySQL) UpdateRatingByID(transaction transaction.Transaction, ID int64, rating int) error {
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

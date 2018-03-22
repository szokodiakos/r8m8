package player

import (
	"github.com/lib/pq"
	"github.com/szokodiakos/r8m8/player/model"
	"github.com/szokodiakos/r8m8/transaction"
)

type playerRepositorySQL struct {
}

func (p *playerRepositorySQL) GetMultipleByUniqueNames(tr transaction.Transaction, uniqueNames []string) ([]model.Player, error) {
	players := []model.Player{}

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
	err := sqlTransaction.Select(&players, query, pq.Array(uniqueNames))
	return players, err
}

func (p *playerRepositorySQL) Create(tr transaction.Transaction, player model.Player) (model.Player, error) {
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
	if err != nil {
		return model.Player{}, err
	}

	return p.GetByUniqueName(tr, player.UniqueName)
}

func (p *playerRepositorySQL) GetByUniqueName(tr transaction.Transaction, uniqueName string) (model.Player, error) {
	repoPlayer := model.Player{}

	query := `
		SELECT 
			p.id,
			p.unique_name,
			p.display_name
		FROM
			players p
		WHERE
			p.unique_name = $1;
	`

	sqlTransaction := transaction.GetSQLTransaction(tr)
	err := sqlTransaction.Get(&repoPlayer, query, uniqueName)
	return repoPlayer, err
}

// NewRepositorySQL factory
func NewRepositorySQL() Repository {
	return &playerRepositorySQL{}
}

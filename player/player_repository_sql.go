package player

import (
	"database/sql"

	"github.com/lib/pq"
	"github.com/szokodiakos/r8m8/entity"
	"github.com/szokodiakos/r8m8/player/errors"
	"github.com/szokodiakos/r8m8/transaction"
)

type playerRepositorySQL struct {
}

func (p *playerRepositorySQL) GetMultipleByIDs(tr transaction.Transaction, ids []string) ([]entity.Player, error) {
	players := []entity.Player{}

	query := `
		SELECT
			p.id,
			p.display_name
		FROM
			players p
		WHERE
			p.id = ANY($1);
	`

	sqlTransaction := transaction.GetSQLTransaction(tr)
	err := sqlTransaction.Select(&players, query, pq.Array(ids))
	return players, err
}

func (p *playerRepositorySQL) Add(tr transaction.Transaction, player entity.Player) (entity.Player, error) {
	var createdID string

	query := `
		INSERT INTO players
			(id, display_name)
		VALUES
			($1, $2)
		RETURNING id;
	`

	sqlTransaction := transaction.GetSQLTransaction(tr)
	err := sqlTransaction.Get(&createdID, query, player.ID, player.DisplayName)
	if err != nil {
		return entity.Player{}, err
	}

	return p.GetByID(tr, createdID)
}

func (p *playerRepositorySQL) GetByID(tr transaction.Transaction, id string) (entity.Player, error) {
	repoPlayer := entity.Player{}

	query := `
		SELECT 
			p.id,
			p.display_name
		FROM
			players p
		WHERE
			p.id = $1;
	`

	sqlTransaction := transaction.GetSQLTransaction(tr)
	err := sqlTransaction.Get(&repoPlayer, query, id)
	if err == sql.ErrNoRows {
		return repoPlayer, &errors.PlayerNotFoundError{
			ID: id,
		}
	}
	return repoPlayer, err
}

// NewPlayerRepositorySQL factory
func NewPlayerRepositorySQL() Repository {
	return &playerRepositorySQL{}
}

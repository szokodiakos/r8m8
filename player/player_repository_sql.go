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

func (p *playerRepositorySQL) Create(tr transaction.Transaction, player model.Player) (int64, error) {
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

func (p *playerRepositorySQL) GetReporterPlayerByMatchID(tr transaction.Transaction, matchID int64) (model.Player, error) {
	repoPlayer := model.Player{}

	query := `
		SELECT 
			p.id,
			p.unique_name,
			p.display_name
		FROM
			players p,
			matches m
		WHERE
			m.id = $1 AND
			m.reporter_player_id = p.id;
	`

	sqlTransaction := transaction.GetSQLTransaction(tr)
	err := sqlTransaction.Get(&repoPlayer, query, matchID)
	return repoPlayer, err
}

// NewRepositorySQL factory
func NewRepositorySQL() Repository {
	return &playerRepositorySQL{}
}

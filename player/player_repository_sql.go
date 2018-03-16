package player

import (
	"github.com/lib/pq"
	"github.com/szokodiakos/r8m8/transaction"
)

type playerRepositorySQL struct {
}

func (p *playerRepositorySQL) GetMultipleByUniqueNames(tr transaction.Transaction, uniqueNames []string) ([]Player, error) {
	players := []Player{}

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

func (p *playerRepositorySQL) GetReporterPlayerByMatchID(tr transaction.Transaction, matchID int64) (Player, error) {
	repoPlayer := Player{}

	query := `
		SELECT 
			p.id,
			p.unique_name,
			p.display_name
		FROM
			p players,
			m matches
		WHERE
			m.id = $1 AND
			m.reporter_player_id = p.id;
	`

	sqlTransaction := transaction.GetSQLTransaction(tr)
	err := sqlTransaction.Get(&repoPlayer, query, matchID)
	return repoPlayer, err
}

func (p *playerRepositorySQL) GetMultipleByMatchID(tr transaction.Transaction, matchID int64) ([]Player, error) {
	players := []Player{}

	query := `
		SELECT
			p.id,
			p.unique_name,
			p.display_name,
			r.rating AS "rating.rating",
			d.rating_change AS "details.rating_change"
		FROM
			p players,
			r ratings,
			d details,
			l leagues
		WHERE
			m.id = $1 AND
			d.match_id = m.id AND
			d.player_id = p.id AND
			r.player_id = p.id AND
			r.league_id = l.id AND
			m.league_id = l.id;
	`

	sqlTransaction := transaction.GetSQLTransaction(tr)
	err := sqlTransaction.Select(&players, query, matchID)
	return players, err
}

// NewRepositorySQL factory
func NewRepositorySQL() Repository {
	return &playerRepositorySQL{}
}

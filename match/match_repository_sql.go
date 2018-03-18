package match

import (
	"database/sql"

	"github.com/szokodiakos/r8m8/match/errors"
	"github.com/szokodiakos/r8m8/match/model"
	"github.com/szokodiakos/r8m8/transaction"
)

type matchRepositorySQL struct {
}

func (m *matchRepositorySQL) Create(tr transaction.Transaction, leagueID int64, reporterPlayerID int64) (int64, error) {
	var createdID int64

	query := `
		INSERT INTO matches
			(league_id, reporter_player_id, created_at)
		VALUES
			($1, $2, current_timestamp)
		RETURNING id;
	`

	sqlTransaction := transaction.GetSQLTransaction(tr)
	err := sqlTransaction.Get(&createdID, query, leagueID, reporterPlayerID)
	return createdID, err
}

func (m *matchRepositorySQL) GetByID(tr transaction.Transaction, matchID int64) (model.Match, error) {
	match := model.Match{}

	query := `
			SELECT
				m.id,
				p.display_name AS "reporter_player.display_name"
			FROM
				matches m,
				players p
			WHERE
				m.id = $1 AND
				m.reporter_player_id = p.id;
		`

	sqlTransaction := transaction.GetSQLTransaction(tr)
	err := sqlTransaction.Get(&match, query, matchID)
	if err == sql.ErrNoRows {
		return match, &errors.MatchNotFoundError{
			ID: matchID,
		}
	}
	return match, err
}

// NewRepositorySQL factory
func NewRepositorySQL() Repository {
	return &matchRepositorySQL{}
}

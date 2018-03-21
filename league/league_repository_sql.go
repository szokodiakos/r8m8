package league

import (
	"database/sql"

	"github.com/szokodiakos/r8m8/league/errors"
	"github.com/szokodiakos/r8m8/league/model"
	"github.com/szokodiakos/r8m8/transaction"
)

type leagueRepositorySQL struct{}

func (l *leagueRepositorySQL) GetByUniqueName(tr transaction.Transaction, uniqueName string) (model.League, error) {
	league := model.League{}

	query := `
		SELECT
			l.id,
			l.display_name
		FROM
			leagues l
		WHERE
			l.unique_name = $1;
	`

	sqlTransaction := transaction.GetSQLTransaction(tr)
	err := sqlTransaction.Get(&league, query, uniqueName)
	if err == sql.ErrNoRows {
		return league, &errors.LeagueNotFoundError{
			UniqueName: uniqueName,
		}
	}
	return league, err
}

func (l *leagueRepositorySQL) Create(tr transaction.Transaction, league model.League) (model.League, error) {
	query := `
		INSERT INTO leagues
			(unique_name, display_name)
		VALUES
			($1, $2);
	`

	sqlTransaction := transaction.GetSQLTransaction(tr)
	_, err := sqlTransaction.Exec(query, league.UniqueName, league.DisplayName)
	if err != nil {
		return league, err
	}

	return l.GetByUniqueName(tr, league.UniqueName)
}

// NewRepositorySQL factory
func NewRepositorySQL() Repository {
	return &leagueRepositorySQL{}
}

package league

import (
	_sql "database/sql"

	"github.com/szokodiakos/r8m8/league/errors"
	"github.com/szokodiakos/r8m8/transaction"
)

type leagueRepositorySQL struct{}

func (l *leagueRepositorySQL) GetByUniqueName(tr transaction.Transaction, uniqueName string) (League, error) {
	league := League{}

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
	if err == _sql.ErrNoRows {
		return league, &errors.LeagueNotFoundError{
			UniqueName: uniqueName,
		}
	}
	return league, err
}

func (l *leagueRepositorySQL) Create(tr transaction.Transaction, league League) error {
	query := `
		INSERT INTO leagues
			(unique_name, display_name)
		VALUES
			($1, $2);
	`

	sqlTransaction := transaction.GetSQLTransaction(tr)
	_, err := sqlTransaction.Exec(query, league.UniqueName, league.DisplayName)
	return err
}

// NewRepositorySQL factory
func NewRepositorySQL() Repository {
	return &leagueRepositorySQL{}
}

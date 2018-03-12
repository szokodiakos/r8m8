package league

import (
	_sql "database/sql"

	"github.com/szokodiakos/r8m8/sql"
	"github.com/szokodiakos/r8m8/transaction"
)

type leagueRepositorySQL struct{}

func (l *leagueRepositorySQL) GetByUniqueName(transaction transaction.Transaction, uniqueName string) (RepoLeague, error) {
	repoLeague := RepoLeague{}

	query := `
		SELECT
			l.id,
			l.display_name
		FROM
			leagues l
		WHERE
			l.unique_name = $1;
	`

	sqlTransaction := transaction.ConcreteTransaction.(sql.Transaction)
	err := sqlTransaction.Get(&repoLeague, query, uniqueName)
	if err == _sql.ErrNoRows {
		return repoLeague, nil
	}
	return repoLeague, err
}

func (l *leagueRepositorySQL) Create(transaction transaction.Transaction, league League) error {
	query := `
		INSERT INTO leagues
			(unique_name, display_name)
		VALUES
			($1, $2);
	`

	sqlTransaction := transaction.ConcreteTransaction.(sql.Transaction)
	_, err := sqlTransaction.Exec(query, league.UniqueName, league.DisplayName)
	return err
}

// NewRepositorySQL factory
func NewRepositorySQL() Repository {
	return &leagueRepositorySQL{}
}

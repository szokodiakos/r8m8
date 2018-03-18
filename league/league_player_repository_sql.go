package league

import (
	"github.com/szokodiakos/r8m8/league/model"
	"github.com/szokodiakos/r8m8/transaction"
)

type leaguePlayerRepositorySQL struct{}

func (s *leaguePlayerRepositorySQL) GetMultipleByLeagueUniqueName(tr transaction.Transaction, uniqueName string) ([]model.LeaguePlayer, error) {
	leaguePlayers := []model.LeaguePlayer{}

	query := `
		SELECT
			p.display_name AS "player.display_name",
			r.rating AS "rating.rating",
			COUNT(CASE WHEN d.has_won THEN 1 END) AS won_match_count,
			COUNT(*) AS total_match_count
		FROM
			players p,
			ratings r,
			leagues l,
			matches m,
			details d
		WHERE
			l.unique_name = $1 AND
			l.id = r.league_id AND
			r.player_id = p.id AND
			p.id = d.player_id AND
			m.league_id = l.id AND
			d.match_id = m.id
		GROUP BY
			p.display_name,
			r.rating
		ORDER BY
			r.rating DESC;
	`

	sqlTransaction := transaction.GetSQLTransaction(tr)
	err := sqlTransaction.Select(&leaguePlayers, query, uniqueName)

	return leaguePlayers, err
}

// NewPlayerRepositorySQL factory
func NewPlayerRepositorySQL() PlayerRepository {
	return &leaguePlayerRepositorySQL{}
}

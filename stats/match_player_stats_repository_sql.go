package stats

import (
	"github.com/szokodiakos/r8m8/stats/model"
	"github.com/szokodiakos/r8m8/transaction"
)

type matchPlayerStatsRepositorySQL struct {
}

func (d *matchPlayerStatsRepositorySQL) GetMultipleByMatchID(tr transaction.Transaction, matchID int64) ([]model.MatchPlayerStats, error) {
	matchPlayersStats := []model.MatchPlayerStats{}

	query := `
		SELECT
			p.display_name AS "player.display_name",
			r.rating AS "rating.rating",
			d.rating_change AS "details.rating_change",
			d.has_won AS "details.has_won"
		FROM
			players p,
			ratings r,
			details d,
			leagues l,
			matches m
		WHERE
			m.id = $1 AND
			d.match_id = m.id AND
			d.player_id = p.id AND
			r.player_id = p.id AND
			r.league_id = l.id AND
			m.league_id = l.id;
	`

	sqlTransaction := transaction.GetSQLTransaction(tr)
	err := sqlTransaction.Select(&matchPlayersStats, query, matchID)
	return matchPlayersStats, err
}

// NewMatchPlayerStatsRepositorySQL factory
func NewMatchPlayerStatsRepositorySQL() MatchPlayerStatsRepository {
	return &matchPlayerStatsRepositorySQL{}
}

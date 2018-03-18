package match

import (
	"github.com/szokodiakos/r8m8/match/model"
	"github.com/szokodiakos/r8m8/transaction"
)

type matchPlayerRepositorySQL struct {
}

func (d *matchPlayerRepositorySQL) GetMultipleByMatchID(tr transaction.Transaction, matchID int64) ([]model.MatchPlayer, error) {
	matchPlayers := []model.MatchPlayer{}

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
	err := sqlTransaction.Select(&matchPlayers, query, matchID)
	return matchPlayers, err
}

// NewPlayerRepositorySQL factory
func NewPlayerRepositorySQL() PlayerRepository {
	return &matchPlayerRepositorySQL{}
}

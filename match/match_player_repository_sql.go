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
			p.display_name AS "league_player.player.display_name",
			lp.rating AS "league_player.rating",
			mp.rating_change AS "rating_change",
			mp.has_won AS "has_won"
		FROM
			players p,
			league_players lp,
			match_players mp,
			matches m
		WHERE
			m.id = $1 AND
			mp.match_id = m.id AND
			mp.player_id = lp.player_id AND
			mp.league_id = lp.league_id AND
			lp.player_id = p.id;
	`

	sqlTransaction := transaction.GetSQLTransaction(tr)
	err := sqlTransaction.Select(&matchPlayers, query, matchID)
	return matchPlayers, err
}

// NewPlayerRepositorySQL factory
func NewPlayerRepositorySQL() PlayerRepository {
	return &matchPlayerRepositorySQL{}
}

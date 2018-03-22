package league

import (
	"github.com/lib/pq"
	"github.com/szokodiakos/r8m8/league/model"
	"github.com/szokodiakos/r8m8/transaction"
)

type leaguePlayerRepositorySQL struct{}

func (l *leaguePlayerRepositorySQL) GetMultipleByLeagueUniqueNameOrderedByRating(tr transaction.Transaction, uniqueName string) ([]model.LeaguePlayer, error) {
	leaguePlayers := []model.LeaguePlayer{}

	query := `
		SELECT
			p.display_name AS "player.display_name",
			lp.rating,
			COUNT(CASE WHEN mp.has_won THEN 1 END) AS win_count,
			COUNT(*) AS match_count
		FROM
			players p,
			league_players lp,
			leagues l,
			matches m,
			match_players mp
		WHERE
			l.unique_name = $1 AND
			l.id = lp.league_id AND
			lp.player_id = p.id AND
			p.id = mp.player_id AND
			m.league_id = l.id AND
			mp.player_id = lp.player_id AND
			mp.league_id = lp.league_id AND
			mp.match_id = m.id
		GROUP BY
			p.display_name,
			lp.rating
		ORDER BY
			lp.rating DESC
		LIMIT 10;
	`

	sqlTransaction := transaction.GetSQLTransaction(tr)
	err := sqlTransaction.Select(&leaguePlayers, query, uniqueName)

	return leaguePlayers, err
}

func (l *leaguePlayerRepositorySQL) GetMultipleByPlayerUniqueNames(tr transaction.Transaction, uniqueNames []string) ([]model.LeaguePlayer, error) {
	leaguePlayers := []model.LeaguePlayer{}

	query := `
		SELECT
			lp.rating,
			p.id AS "player.id",
			l.id AS "league.id",
			p.unique_name AS "player.unique_name"
		FROM
			league_players lp,
			leagues l,
			players p
		WHERE
			lp.player_id = p.id AND
			lp.league_id = l.id AND
			p.unique_name = ANY($1);
	`

	sqlTransaction := transaction.GetSQLTransaction(tr)
	err := sqlTransaction.Select(&leaguePlayers, query, pq.Array(uniqueNames))
	return leaguePlayers, err
}

func (l *leaguePlayerRepositorySQL) Update(tr transaction.Transaction, leaguePlayer model.LeaguePlayer) error {
	query := `
		UPDATE league_players
		SET
			rating = $1
		WHERE
			player_id = $2 AND
			league_id = $3;
	`

	sqlTransaction := transaction.GetSQLTransaction(tr)
	_, err := sqlTransaction.Exec(query, leaguePlayer.Rating, leaguePlayer.Player.ID, leaguePlayer.League.ID)
	return err
}

func (l *leaguePlayerRepositorySQL) Create(tr transaction.Transaction, leaguePlayer model.LeaguePlayer) error {
	query := `
			INSERT INTO league_players
				(player_id, league_id, rating)
			VALUES
				($1, $2, $3);
		`

	sqlTransaction := transaction.GetSQLTransaction(tr)
	_, err := sqlTransaction.Exec(query, leaguePlayer.Player.ID, leaguePlayer.League.ID, leaguePlayer.Rating)
	return err
}

// NewPlayerRepositorySQL factory
func NewPlayerRepositorySQL() PlayerRepository {
	return &leaguePlayerRepositorySQL{}
}

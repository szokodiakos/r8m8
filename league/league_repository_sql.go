package league

import (
	"database/sql"

	"github.com/szokodiakos/r8m8/entity"
	"github.com/szokodiakos/r8m8/league/errors"
	"github.com/szokodiakos/r8m8/transaction"
)

type leagueRepositorySQL struct {
}

type leaguePlayerSQL struct {
	PlayerID          string `db:"player_id"`
	PlayerDisplayName string `db:"player_display_name"`
	WinCount          int    `db:"win_count"`
	MatchCount        int    `db:"match_count"`
	Rating            int    `db:"rating"`
}

func (l *leagueRepositorySQL) GetByID(tr transaction.Transaction, id string) (entity.League, error) {
	league := entity.League{}

	query := `
		SELECT
			l.id,
			l.display_name
		FROM
			leagues l
		WHERE
			l.id = $1;
	`

	sqlTransaction := transaction.GetSQLTransaction(tr)
	err := sqlTransaction.Get(&league, query, id)
	if err == sql.ErrNoRows {
		return league, &errors.LeagueNotFoundError{
			ID: id,
		}
	}
	if err != nil {
		return league, err
	}

	leaguePlayers, err := getLeaguePlayersByLeagueID(tr, id)
	if err != nil {
		return league, err
	}

	league.LeaguePlayers = leaguePlayers
	return league, nil
}

func getLeaguePlayersByLeagueID(tr transaction.Transaction, leagueID string) ([]entity.LeaguePlayer, error) {
	leaguePlayers := []entity.LeaguePlayer{}
	leaguePlayersSQL := []leaguePlayerSQL{}

	query := `
		SELECT
			p.id AS "player_id",
			p.display_name AS "player_display_name",
			lp.rating AS "rating",
			COUNT(CASE WHEN mp.has_won THEN 1 END) AS "win_count",
			COUNT(*) AS "match_count"
		FROM
			players p,
			league_players lp,
			leagues l,
			matches m,
			match_players mp
		WHERE
			l.id = $1 AND
			l.id = lp.league_id AND
			lp.player_id = p.id AND
			p.id = mp.player_id AND
			m.league_id = l.id AND
			mp.player_id = lp.player_id AND
			mp.league_id = lp.league_id AND
			mp.match_id = m.id
		GROUP BY
			p.id,
			p.display_name,
			lp.rating;
	`

	sqlTransaction := transaction.GetSQLTransaction(tr)
	err := sqlTransaction.Select(&leaguePlayersSQL, query, leagueID)
	if err != nil {
		return leaguePlayers, err
	}

	leaguePlayers = mapToLeaguePlayers(leaguePlayersSQL)
	return leaguePlayers, nil
}

func mapToLeaguePlayers(leaguePlayersSQL []leaguePlayerSQL) []entity.LeaguePlayer {
	leaguePlayers := make([]entity.LeaguePlayer, len(leaguePlayersSQL))
	for i := range leaguePlayersSQL {
		playerID := leaguePlayersSQL[i].PlayerID
		playerDisplayName := leaguePlayersSQL[i].PlayerDisplayName
		rating := leaguePlayersSQL[i].Rating
		matchCount := leaguePlayersSQL[i].MatchCount
		winCount := leaguePlayersSQL[i].WinCount

		player := entity.Player{
			ID:          playerID,
			DisplayName: playerDisplayName,
		}

		leaguePlayer := entity.LeaguePlayer{
			PlayerID: playerID,
			Rating:   rating,
		}

		leaguePlayers[i] = entity.NewLeaguePlayer(leaguePlayer, player, winCount, matchCount)
	}
	return leaguePlayers
}

func (l *leagueRepositorySQL) Add(tr transaction.Transaction, league entity.League) (entity.League, error) {
	query := `
		INSERT INTO leagues
			(id, display_name)
		VALUES
			($1, $2);
	`

	sqlTransaction := transaction.GetSQLTransaction(tr)
	_, err := sqlTransaction.Exec(query, league.ID, league.DisplayName)
	if err != nil {
		return league, err
	}

	leaguePlayers := league.LeaguePlayers
	err = addLeaguePlayers(tr, leaguePlayers, league.ID)
	if err != nil {
		return league, err
	}

	return l.GetByID(tr, league.ID)
}

func addLeaguePlayers(tr transaction.Transaction, leaguePlayers []entity.LeaguePlayer, leagueID string) error {
	for i := range leaguePlayers {
		err := addLeaguePlayer(tr, leaguePlayers[i], leagueID)
		if err != nil {
			return err
		}
	}
	return nil
}

func addLeaguePlayer(tr transaction.Transaction, leaguePlayer entity.LeaguePlayer, leagueID string) error {
	query := `
			INSERT INTO league_players
				(player_id, league_id, rating)
			VALUES
				($1, $2, $3)
			ON CONFLICT
				(player_id, league_id)
			DO UPDATE
				SET rating = $3;
		`

	sqlTransaction := transaction.GetSQLTransaction(tr)
	_, err := sqlTransaction.Exec(query, leaguePlayer.PlayerID, leagueID, leaguePlayer.Rating)
	return err
}

func (l *leagueRepositorySQL) Update(tr transaction.Transaction, league entity.League) error {
	err := addLeaguePlayers(tr, league.LeaguePlayers, league.ID)
	return err
}

// NewLeagueRepositorySQL factory
func NewLeagueRepositorySQL() Repository {
	return &leagueRepositorySQL{}
}

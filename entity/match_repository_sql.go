package entity

import (
	"database/sql"
	"time"

	"github.com/szokodiakos/r8m8/match/errors"
	"github.com/szokodiakos/r8m8/transaction"
)

type matchRepositorySQL struct {
}

type matchSQL struct {
	ID                        int64     `db:"id"`
	LeagueID                  string    `db:"league_id"`
	ReporterPlayerID          string    `db:"reporter_player_id"`
	ReporterPlayerDisplayName string    `db:"reporter_player_display_name"`
	CreatedAt                 time.Time `db:"created_at"`
}

type matchPlayerSQL struct {
	PlayerID           string `db:"player_id"`
	PlayerDisplayName  string `db:"player_display_name"`
	LeaguePlayerRating int    `db:"league_player_rating"`
	RatingChange       int    `db:"rating_change"`
	HasWon             bool   `db:"has_won"`
}

func (m *matchRepositorySQL) Add(tr transaction.Transaction, match Match) (Match, error) {
	var createdID int64

	query := `
		INSERT INTO matches
			(league_id, reporter_player_id, created_at)
		VALUES
			($1, $2, current_timestamp)
		RETURNING id;
	`

	sqlTransaction := transaction.GetSQLTransaction(tr)
	err := sqlTransaction.Get(&createdID, query, match.LeagueID, match.ReporterPlayerID)
	if err != nil {
		return match, err
	}

	err = addMatchPlayers(tr, match.MatchPlayers, createdID, match.LeagueID)
	if err != nil {
		return match, err
	}

	return m.GetByID(tr, createdID)
}

func addMatchPlayers(tr transaction.Transaction, matchPlayers []MatchPlayer, matchID int64, leagueID string) error {
	for i := range matchPlayers {
		err := addMatchPlayer(tr, matchPlayers[i], matchID, leagueID)
		if err != nil {
			return err
		}
	}
	return nil
}

func addMatchPlayer(tr transaction.Transaction, matchPlayer MatchPlayer, matchID int64, leagueID string) error {
	query := `
			INSERT INTO match_players
				(player_id, league_id, match_id, rating_change, has_won)
			VALUES
				($1, $2, $3, $4, $5);
		`

	sqlTransaction := transaction.GetSQLTransaction(tr)
	_, err := sqlTransaction.Exec(query, matchPlayer.PlayerID, leagueID, matchID, matchPlayer.RatingChange, matchPlayer.HasWon)
	return err
}

func (m *matchRepositorySQL) GetByID(tr transaction.Transaction, matchID int64) (Match, error) {
	match := Match{}
	matchSQL := matchSQL{}

	query := `
		SELECT
			m.id,
			p.id AS "reporter_player_id",
			p.display_name AS "reporter_player_display_name"
		FROM
			matches m,
			players p
		WHERE
			m.id = $1 AND
			m.reporter_player_id = p.id
		LIMIT 1;
	`

	sqlTransaction := transaction.GetSQLTransaction(tr)
	err := sqlTransaction.Get(&matchSQL, query, matchID)
	if err == sql.ErrNoRows {
		return match, &errors.MatchNotFoundError{
			ID: matchID,
		}
	}
	if err != nil {
		return match, err
	}

	match = mapToMatch(matchSQL)

	matchPlayers, err := getMatchPlayersByMatchID(tr, matchID)
	if err != nil {
		return match, err
	}

	match.MatchPlayers = matchPlayers
	return match, err
}

func mapToMatch(matchSQL matchSQL) Match {
	return Match{
		ID:               matchSQL.ID,
		LeagueID:         matchSQL.LeagueID,
		CreatedAt:        matchSQL.CreatedAt,
		ReporterPlayerID: matchSQL.ReporterPlayerID,
		reporterPlayer: Player{
			ID:          matchSQL.ReporterPlayerID,
			DisplayName: matchSQL.ReporterPlayerDisplayName,
		},
	}
}

func getMatchPlayersByMatchID(tr transaction.Transaction, matchID int64) ([]MatchPlayer, error) {
	matchPlayers := []MatchPlayer{}
	matchPlayersSQL := []matchPlayerSQL{}

	query := `
		SELECT
			p.id AS "player_id",
			p.display_name AS "player_display_name",
			lp.rating AS "league_player_rating",
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
	err := sqlTransaction.Select(&matchPlayersSQL, query, matchID)
	if err != nil {
		return matchPlayers, err
	}

	matchPlayers = mapToMatchPlayers(matchPlayersSQL)
	return matchPlayers, nil
}

func mapToMatchPlayers(matchPlayersSQL []matchPlayerSQL) []MatchPlayer {
	matchPlayers := make([]MatchPlayer, len(matchPlayersSQL))
	for i := range matchPlayersSQL {
		playerID := matchPlayersSQL[i].PlayerID
		playerDisplayName := matchPlayersSQL[i].PlayerDisplayName
		leaguePlayerRating := matchPlayersSQL[i].LeaguePlayerRating
		ratingChange := matchPlayersSQL[i].RatingChange
		hasWon := matchPlayersSQL[i].HasWon

		matchPlayers[i] = MatchPlayer{
			PlayerID:     playerID,
			HasWon:       hasWon,
			RatingChange: ratingChange,
			leaguePlayer: LeaguePlayer{
				PlayerID: playerID,
				Rating:   leaguePlayerRating,
				player: Player{
					ID:          playerID,
					DisplayName: playerDisplayName,
				},
			},
		}
	}
	return matchPlayers
}

func (m *matchRepositorySQL) GetLatestByReporterPlayerID(tr transaction.Transaction, reporterPlayerID string) (Match, error) {
	match := Match{}
	matchSQL := matchSQL{}

	query := `
		SELECT
			m.id,
			m.league_id,
			p.id AS "reporter_player_id",
			p.display_name AS "reporter_player_display_name"
		FROM
			matches m,
			players p
		WHERE
			p.id = $1 AND
			m.reporter_player_id = p.id
		ORDER BY
			m.created_at DESC
		LIMIT 1;
	`

	sqlTransaction := transaction.GetSQLTransaction(tr)
	err := sqlTransaction.Get(&matchSQL, query, reporterPlayerID)
	if err == sql.ErrNoRows {
		return match, &errors.MatchNotFoundError{
			ReporterPlayerID: reporterPlayerID,
		}
	}
	if err != nil {
		return match, err
	}

	match = mapToMatch(matchSQL)

	matchPlayers, err := getMatchPlayersByMatchID(tr, match.ID)
	if err != nil {
		return match, err
	}

	match.MatchPlayers = matchPlayers
	return match, err
}

func (m *matchRepositorySQL) Remove(tr transaction.Transaction, match Match) error {
	query := `
			DELETE FROM
				matches m
			WHERE
				m.id = $1;
		`

	sqlTransaction := transaction.GetSQLTransaction(tr)
	_, err := sqlTransaction.Exec(query, match.ID)
	return err
}

// NewMatchRepositorySQL factory
func NewMatchRepositorySQL() MatchRepository {
	return &matchRepositorySQL{}
}

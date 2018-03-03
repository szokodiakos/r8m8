package player

import (
	"database/sql"
)

type playerSlackRepositorySQL struct {
	db *sql.DB
}

func (psrs *playerSlackRepositorySQL) GetMultipleByUserIDs(userIDs []string, teamID string) ([]Slack, error) {
	var slackPlayers = make([]Slack, 0, len(userIDs))
	query := `
		SELECT
			p.id,
			p.rating,
			sp.user_id,
			sp.username,
			sp.team_id
		FROM
			players p,
			slack_players sp
		WHERE
			sp.user_id IN (?) AND
			sp.team_id = ?
	`
	rows, err := psrs.db.Query(query, userIDs, teamID)
	if err != nil {
		return slackPlayers, err
	}

	for rows.Next() {
		var id int64
		var rating int
		var userID, username, teamID string

		if err := rows.Scan(&id, &rating, userID, username, teamID); err != nil {
			return slackPlayers, err
		}

		slackPlayer := Slack{
			Player: Player{
				ID:     id,
				Rating: rating,
			},
			UserID:   userID,
			Username: username,
			TeamID:   teamID,
		}
		slackPlayers = append(slackPlayers, slackPlayer)
	}
	return slackPlayers, nil
}

func (psrs *playerSlackRepositorySQL) Create(slackPlayer Slack) (int64, error) {
	var createdID int64
	query := `
		INSERT INTO slack_players
			(player_id, user_id, username, team_id)
		VALUES
			(?, ?, ?, ?);
	`

	res, err := psrs.db.Exec(query, slackPlayer.Player.ID, slackPlayer.UserID, slackPlayer.Username, slackPlayer.TeamID)
	if err != nil {
		return createdID, err
	}

	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return createdID, err
	}

	createdID = lastInsertID
	return createdID, nil
}

// NewSlackRepository factory
func NewSlackRepository(db *sql.DB) SlackRepository {
	return &playerSlackRepositorySQL{
		db: db,
	}
}

package player

import (
	"github.com/szokodiakos/r8m8/sql"

	"github.com/lib/pq"
)

type playerSlackRepositorySQL struct {
	db sql.DB
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
			sp.player_id = p.id AND
			sp.user_id = ANY($1) AND
			sp.team_id = $2;
	`

	rows, err := psrs.db.Query(query, pq.Array(userIDs), teamID)
	if err != nil {
		return slackPlayers, err
	}

	for rows.Next() {
		var id int64
		var rating int
		var userID, username, teamID string

		if err := rows.Scan(&id, &rating, &userID, &username, &teamID); err != nil {
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

func (psrs *playerSlackRepositorySQL) Create(slackPlayer Slack) error {
	query := `
		INSERT INTO slack_players
			(player_id, user_id, username, team_id)
		VALUES
			($1, $2, $3, $4);
	`

	_, err := psrs.db.Exec(query, slackPlayer.Player.ID, slackPlayer.UserID, slackPlayer.Username, slackPlayer.TeamID)
	if err != nil {
		return err
	}

	return nil
}

// NewSlackRepository factory
func NewSlackRepository(db sql.DB) SlackRepository {
	return &playerSlackRepositorySQL{
		db: db,
	}
}

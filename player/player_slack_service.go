package player

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/szokodiakos/r8m8/entity"
	"github.com/szokodiakos/r8m8/player/errors"
)

// SlackService interface
type SlackService interface {
	ToPlayers(text string, teamID string) ([]entity.Player, error)
	ToPlayer(teamID string, userID string, userName string) entity.Player
}

type playerSlackService struct {
}

func (p *playerSlackService) ToPlayers(text string, teamID string) ([]entity.Player, error) {
	slackPlayers := strings.Split(text, " ")
	players := make([]entity.Player, len(slackPlayers))

	for i := range slackPlayers {
		player, err := p.parsePlayer(slackPlayers[i], teamID)

		if err != nil {
			return nil, err
		}

		players[i] = player
	}
	return players, nil
}

func (p *playerSlackService) parsePlayer(slackPlayer string, teamID string) (entity.Player, error) {
	var player entity.Player
	pattern, _ := regexp.Compile(`<@(.*)\|(.*)>`)
	results := pattern.FindStringSubmatch(slackPlayer)

	if isSlackPlayerInvalid(results) {
		return player, &errors.BadSlackPlayerFormatError{SlackPlayer: slackPlayer}
	}

	userID := results[1]
	userName := results[2]

	player = p.ToPlayer(teamID, userID, userName)
	return player, nil
}

func isSlackPlayerInvalid(results []string) bool {
	return (len(results) != 3)
}

func (p *playerSlackService) ToPlayer(teamID string, userID string, userName string) entity.Player {
	displayName := userName
	id := fmt.Sprintf("slack_%v_%v", teamID, userID)
	return entity.Player{
		ID:          id,
		DisplayName: displayName,
	}
}

// NewSlackService factory
func NewSlackService() SlackService {
	return &playerSlackService{}
}

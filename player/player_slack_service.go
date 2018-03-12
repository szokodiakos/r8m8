package player

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/szokodiakos/r8m8/player/errors"
)

// SlackService interface
type SlackService interface {
	ToPlayers(text string, teamID string) ([]Player, error)
	ToPlayer(teamID string, userID string, userName string) Player
}

type playerSlackService struct {
}

func (p *playerSlackService) ToPlayers(text string, teamID string) ([]Player, error) {
	slackPlayers := strings.Split(text, " ")
	players := make([]Player, len(slackPlayers))

	for i := range slackPlayers {
		player, err := p.parsePlayer(slackPlayers[i], teamID)

		if err != nil {
			return nil, err
		}

		players[i] = player
	}
	return players, nil
}

func (p *playerSlackService) parsePlayer(slackPlayer string, teamID string) (Player, error) {
	var player Player
	pattern, _ := regexp.Compile(`<@(.*)\|(.*)>`)
	results := pattern.FindStringSubmatch(slackPlayer)

	if isSlackPlayerInvalid(results) {
		return player, errors.NewBadSlackPlayerFormatError(slackPlayer)
	}

	userID := results[1]
	userName := results[2]

	player = p.ToPlayer(teamID, userID, userName)
	return player, nil
}

func isSlackPlayerInvalid(results []string) bool {
	return (len(results) != 3)
}

func (p *playerSlackService) ToPlayer(teamID string, userID string, userName string) Player {
	displayName := userName
	uniqueName := fmt.Sprintf("slack_%v_%v", teamID, userID)
	return Player{
		DisplayName: displayName,
		UniqueName:  uniqueName,
	}
}

// NewSlackService factory
func NewSlackService() SlackService {
	return &playerSlackService{}
}

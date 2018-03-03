package player

import (
	"regexp"
	"strings"

	"github.com/szokodiakos/r8m8/match/errors"
)

// SlackParserService interface
type SlackParserService interface {
	Parse(text string, teamID string) ([]Slack, error)
}

type playerSlackParserService struct {
}

func (psps *playerSlackParserService) Parse(text string, teamID string) ([]Slack, error) {
	rawSlackPlayers := strings.Split(text, " ")
	slackPlayers := make([]Slack, len(rawSlackPlayers))

	for i := range rawSlackPlayers {
		slackPlayer, err := psps.parseSlackPlayer(rawSlackPlayers[i], teamID)

		if err != nil {
			return nil, err
		}

		slackPlayers[i] = slackPlayer
	}
	return slackPlayers, nil
}

func (psps *playerSlackParserService) parseSlackPlayer(rawSlackPlayer string, teamID string) (Slack, error) {
	var parsedSlackPlayer Slack
	pattern, _ := regexp.Compile(`<@(.*)\|(.*)>`)
	results := pattern.FindStringSubmatch(rawSlackPlayer)

	if psps.isRawSlackPlayerInvalid(results) {
		return parsedSlackPlayer, errors.NewBadRawPlayerFormatError(rawSlackPlayer)
	}

	userID := results[1]
	username := results[2]

	parsedSlackPlayer = Slack{
		UserID:   userID,
		Username: username,
		TeamID:   teamID,
	}
	return parsedSlackPlayer, nil
}

func (psps *playerSlackParserService) isRawSlackPlayerInvalid(results []string) bool {
	return (len(results) != 3)
}

// NewSlackParserService factory
func NewSlackParserService() SlackParserService {
	return &playerSlackParserService{}
}

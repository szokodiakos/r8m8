package match

import (
	"github.com/szokodiakos/r8m8/player"
	"github.com/szokodiakos/r8m8/slack"
)

// SlackService interface
type SlackService interface {
	AddMatch(values string) (slack.MessageResponse, error)
}

type matchSlackService struct {
	matchService       Service
	slackService       slack.Service
	playerSlackService player.SlackService
}

func (sms *matchSlackService) AddMatch(values string) (slack.MessageResponse, error) {
	var messageResponse slack.MessageResponse
	requestValues := sms.slackService.ParseRequestValues(values)
	text := requestValues.Text
	teamID := requestValues.TeamID

	sms.playerSlackService.GetOrAddSlackPlayers(text, teamID)
	// if err != nil {
	// 	return messageResponse, err
	// }
	// if sms.isSlackPlayerCountUneven(slackPlayers) {
	// 	return messageResponse, errors.NewUnevenMatchPlayersError()
	// }
	// // winnerPlayers := sms.getWinnerPlayers(players)
	// // loserPlayers := sms.getLoserPlayers(players)
	// match, err := sms.matchService.AddMatch()
	// if err != nil {
	// 	return messageResponse, err
	// }

	return messageResponse, nil
}

func (sms *matchSlackService) isSlackPlayerCountUneven(players []player.Slack) bool {
	return (len(players) % 2) != 0
}

func (sms *matchSlackService) getWinnerPlayers(players []player.Player) []player.Player {
	lowerhalfPlayers := players[:(len(players) / 2)]
	return lowerhalfPlayers
}

func (sms *matchSlackService) getLoserPlayers(players []player.Player) []player.Player {
	upperhalfPlayers := players[(len(players) / 2):]
	return upperhalfPlayers
}

// NewSlackService factory
func NewSlackService(matchService Service, slackService slack.Service, playerSlackService player.SlackService) SlackService {
	return &matchSlackService{
		matchService:       matchService,
		slackService:       slackService,
		playerSlackService: playerSlackService,
	}
}

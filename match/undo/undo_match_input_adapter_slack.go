package undo

import (
	"github.com/szokodiakos/r8m8/player"
	"github.com/szokodiakos/r8m8/slack"
)

type undoMatchInputAdapterSlack struct {
	slackService       slack.Service
	playerSlackService player.SlackService
}

func (a *undoMatchInputAdapterSlack) Handle(data interface{}) (Input, error) {
	values := data.(string)
	var input Input

	requestValues, err := a.slackService.ParseRequestValues(values)
	if err != nil {
		return input, err
	}

	teamID := requestValues.TeamID
	channelID := requestValues.ChannelID
	userID := requestValues.UserID
	userName := requestValues.UserName
	reporterPlayer := a.playerSlackService.ToPlayer(teamID, channelID, userID, userName)

	input = Input{
		ReporterPlayer: reporterPlayer,
	}
	return input, err
}

// NewUndoMatchInputAdapterSlack factory
func NewUndoMatchInputAdapterSlack(
	slackService slack.Service,
	playerSlackService player.SlackService,
) InputAdapter {
	return &undoMatchInputAdapterSlack{
		slackService:       slackService,
		playerSlackService: playerSlackService,
	}
}

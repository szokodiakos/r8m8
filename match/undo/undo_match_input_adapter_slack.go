package undo

import (
	"github.com/szokodiakos/r8m8/player"
	"github.com/szokodiakos/r8m8/slack"
)

type undoMatchInputAdapterSlack struct {
	slackService       slack.Service
	playerSlackService player.SlackService
}

func (a *undoMatchInputAdapterSlack) Handle(data interface{}) (UndoMatchInput, error) {
	values := data.(string)
	var input UndoMatchInput

	requestValues, err := a.slackService.ParseRequestValues(values)
	if err != nil {
		return input, err
	}

	teamID := requestValues.TeamID
	userID := requestValues.UserID
	userName := requestValues.UserName
	reporterPlayer := a.playerSlackService.ToPlayer(teamID, userID, userName)

	input = UndoMatchInput{
		ReporterPlayer: reporterPlayer,
	}
	return input, err
}

// NewUndoMatchInputAdapterSlack factory
func NewUndoMatchInputAdapterSlack(
	slackService slack.Service,
	playerSlackService player.SlackService,
) UndoMatchInputAdapter {
	return &undoMatchInputAdapterSlack{
		slackService:       slackService,
		playerSlackService: playerSlackService,
	}
}

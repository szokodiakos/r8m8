package match

import (
	"github.com/szokodiakos/r8m8/echo"
	"github.com/szokodiakos/r8m8/match/errors"
	"github.com/szokodiakos/r8m8/slack"
)

type matchSlackErrorHandler struct {
}

func (m *matchSlackErrorHandler) HandleError(err error) interface{} {
	var messageResponse slack.MessageResponse
	switch err.(type) {
	case *errors.ReporterPlayerNotInLeagueError:
		messageResponse = getReporterPlayerNotInLeagueResponse()
		return messageResponse
	case *errors.UnevenMatchPlayersError:
		messageResponse = getUnevenMatchPlayersResponse()
		return messageResponse
	default:
		messageResponse = getDefaultErrorResponse()
		return messageResponse
	}
}

func getReporterPlayerNotInLeagueResponse() slack.MessageResponse {
	text := `
> Darn! You must be the participant of at least one match (including this one). :hushed:
> :exclamation: Please play a match before posting! :exclamation:
	`
	return slack.MessageResponse{
		Text: text,
	}
}

func getUnevenMatchPlayersResponse() slack.MessageResponse {
	text := `
> Darn! Reported players are uneven! :hushed:
> :exclamation: Make sure you report even number of players! :exclamation:
	`
	return slack.MessageResponse{
		Text: text,
	}
}

func getDefaultErrorResponse() slack.MessageResponse {
	text := `
> This is embarrassing!
	`
	return slack.MessageResponse{
		Text: text,
	}
}

// NewSlackErrorHandler factory
func NewSlackErrorHandler() echo.ErrorHandler {
	return &matchSlackErrorHandler{}
}

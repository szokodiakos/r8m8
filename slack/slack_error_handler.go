package slack

import (
	"github.com/szokodiakos/r8m8/echo"
	matchErrors "github.com/szokodiakos/r8m8/match/errors"
)

type slackErrorHandler struct {
}

func (s *slackErrorHandler) HandleError(err error) interface{} {
	var messageResponse MessageResponse
	switch err.(type) {
	case *matchErrors.ReporterPlayerNotInLeagueError:
		messageResponse = getReporterPlayerNotInLeagueResponse()
		return messageResponse
	case *matchErrors.UnevenMatchPlayersError:
		messageResponse = getUnevenMatchPlayersResponse()
		return messageResponse
	default:
		messageResponse = getDefaultErrorResponse()
		return messageResponse
	}
}

func getReporterPlayerNotInLeagueResponse() MessageResponse {
	text := `
> Darn! You must be the participant of at least one match (including this one). :hushed:
> :exclamation: Please play a match before posting! :exclamation:
	`
	return MessageResponse{
		Text: text,
	}
}

func getUnevenMatchPlayersResponse() MessageResponse {
	text := `
> Darn! Reported players are uneven! :hushed:
> :exclamation: Make sure you report even number of players! :exclamation:
	`
	return MessageResponse{
		Text: text,
	}
}

func getDefaultErrorResponse() MessageResponse {
	text := `
> This is embarrassing!
	`
	return MessageResponse{
		Text: text,
	}
}

// NewErrorHandler factory
func NewErrorHandler() echo.ErrorHandler {
	return &slackErrorHandler{}
}

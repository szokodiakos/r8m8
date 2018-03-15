package slack

import (
	"log"
	"net/http"

	"github.com/szokodiakos/r8m8/echo"
	"github.com/szokodiakos/r8m8/match/errors"
)

type slackErrorHandler struct {
}

func (s *slackErrorHandler) HandleError(err error) (int, interface{}) {
	switch err.(type) {
	case *errors.ReporterPlayerNotInLeagueError:
		return getReporterPlayerNotInLeagueResponse()
	case *errors.UnevenMatchPlayersError:
		return getUnevenMatchPlayersResponse()
	default:
		log.Println(err)
		return getDefaultErrorResponse()
	}
}

func getReporterPlayerNotInLeagueResponse() (int, MessageResponse) {
	text := `
> Darn! You must be the participant of at least one match (including this one). :hushed:
> :exclamation: Please play a match before posting! :exclamation:
`
	return getResponse(text)
}

func getUnevenMatchPlayersResponse() (int, MessageResponse) {
	text := `
> Darn! Reported players are uneven! :hushed:
> :exclamation: Make sure you report even number of players! :exclamation:
`
	return getResponse(text)
}

func getDefaultErrorResponse() (int, MessageResponse) {
	text := `
> This is embarrassing!
`
	return getResponse(text)
}

func getResponse(text string) (int, MessageResponse) {
	return http.StatusOK, MessageResponse{
		Text: text,
	}
}

// NewErrorHandler factory
func NewErrorHandler() echo.HTTPErrorHandler {
	return &slackErrorHandler{}
}

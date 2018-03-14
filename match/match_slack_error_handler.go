package match

import (
	"log"
	"net/http"

	"github.com/szokodiakos/r8m8/echo"
	"github.com/szokodiakos/r8m8/match/errors"
	"github.com/szokodiakos/r8m8/slack"
)

type matchSlackErrorHandler struct {
}

func (m *matchSlackErrorHandler) HandleError(err error) (int, interface{}) {
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

func getReporterPlayerNotInLeagueResponse() (int, slack.MessageResponse) {
	text := `
> Darn! You must be the participant of at least one match (including this one). :hushed:
> :exclamation: Please play a match before posting! :exclamation:
`
	return getResponse(text)
}

func getUnevenMatchPlayersResponse() (int, slack.MessageResponse) {
	text := `
> Darn! Reported players are uneven! :hushed:
> :exclamation: Make sure you report even number of players! :exclamation:
`
	return getResponse(text)
}

func getDefaultErrorResponse() (int, slack.MessageResponse) {
	text := `
> This is embarrassing!
`
	return getResponse(text)
}

func getResponse(text string) (int, slack.MessageResponse) {
	return http.StatusOK, slack.MessageResponse{
		Text: text,
	}
}

// NewSlackErrorHandler factory
func NewSlackErrorHandler() echo.ErrorHandler {
	return &matchSlackErrorHandler{}
}

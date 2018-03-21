package slack

import (
	"net/http"

	"github.com/szokodiakos/r8m8/echo"
)

type slackErrorHandler struct {
}

func (s *slackErrorHandler) HandleError(err error) (int, interface{}) {
	return getDefaultErrorResponse()
}

func getDefaultErrorResponse() (int, MessageResponse) {
	text := `
> This is embarrassing!
`
	return getResponse(text)
}

func getResponse(text string) (int, MessageResponse) {
	return http.StatusOK, CreateDirectResponse(text)
}

// NewErrorHandler factory
func NewErrorHandler() echo.HTTPErrorHandler {
	return &slackErrorHandler{}
}

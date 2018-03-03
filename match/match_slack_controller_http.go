package match

import (
	"net/http"

	"github.com/szokodiakos/r8m8/slack"

	"github.com/labstack/echo"
	echoExtensions "github.com/szokodiakos/r8m8/echo"
)

// SlackControllerHTTP struct
type SlackControllerHTTP struct {
	matchSlackService SlackService
}

func (sch *SlackControllerHTTP) postSlackMatch(context echo.Context) error {
	body := context.Get("parsedBody").(string)
	response, err := sch.matchSlackService.AddMatch(body)
	if err != nil {
		return err
	}
	return context.JSON(http.StatusOK, response)
}

// NewSlackControllerHTTP factory
func NewSlackControllerHTTP(e *echo.Echo, matchSlackService SlackService, slackService slack.Service, verificationToken string) *SlackControllerHTTP {
	handler := &SlackControllerHTTP{
		matchSlackService: matchSlackService,
	}
	bodyParser := echoExtensions.BodyParser()
	slackTokenVerifier := echoExtensions.SlackTokenVerifier(slackService, verificationToken)
	slackRoutes := e.Group("/slack", bodyParser, slackTokenVerifier)
	slackRoutes.POST("/match", handler.postSlackMatch)
	return handler
}

package slack

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/spf13/viper"
	echoExtensions "github.com/szokodiakos/r8m8/echo"
)

// HTTPController struct
type HTTPController struct {
	slackMatchService MatchService
}

func (c *HTTPController) postSlackMatch(context echo.Context) error {
	body := context.Get("parsedBody").(string)
	match, err := c.slackMatchService.ParseMatch(body)
	if err != nil {
		return err
	}
	fmt.Println(match)
	slackMessageResponse := MessageResponse{
		Text: "hello",
	}
	return context.JSON(http.StatusOK, slackMessageResponse)
}

// NewHTTPController creates a HTTPController
func NewHTTPController(e *echo.Echo, slackMatchService MatchService) *HTTPController {
	handler := &HTTPController{
		slackMatchService: slackMatchService,
	}
	verificationToken := viper.GetString("slack_verification_token")
	slackRoutes := e.Group("/slack", echoExtensions.BodyParser(), echoExtensions.SlackTokenVerifier(verificationToken))
	slackRoutes.POST("/match", handler.postSlackMatch)
	return handler
}

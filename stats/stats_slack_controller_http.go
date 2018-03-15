package stats

import (
	"net/http"

	"github.com/szokodiakos/r8m8/slack"

	"github.com/labstack/echo"
)

// SlackControllerHTTP struct
type SlackControllerHTTP struct {
	statsSlackService SlackService
}

func (sch *SlackControllerHTTP) postsStatsLeaderboard(context echo.Context) error {
	body := context.Get("parsedBody").(string)
	response, err := sch.statsSlackService.GetLeaderboard(body)
	if err != nil {
		return err
	}
	return context.JSON(http.StatusOK, response)
}

// NewSlackControllerHTTP factory
func NewSlackControllerHTTP(slackGroup *echo.Group, statsSlackService SlackService, slackService slack.Service) *SlackControllerHTTP {
	handler := &SlackControllerHTTP{
		statsSlackService: statsSlackService,
	}
	slackGroup.POST("/stats/leaderboard", handler.postsStatsLeaderboard)
	return handler
}

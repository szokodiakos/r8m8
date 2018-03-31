package leaderboard

import (
	"net/http"

	"github.com/labstack/echo"
)

// LeaderboardControllerHTTP struct
type LeaderboardControllerHTTP struct {
	inputAdapter  LeaderboardInputAdapter
	outputAdapter LeaderboardOutputAdapter
	useCase       UseCase
}

func (g *LeaderboardControllerHTTP) postsStatsLeaderboard(context echo.Context) error {
	body := context.Get("parsedBody").(string)
	input, err := g.inputAdapter.Handle(body)
	if err != nil {
		return err
	}

	output, err := g.useCase.Handle(input)

	response, err := g.outputAdapter.Handle(output, err)
	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, response)
}

// NewLeaderboardControllerHTTP factory
func NewLeaderboardControllerHTTP(
	routeGroup *echo.Group,
	inputAdapter LeaderboardInputAdapter,
	outputAdapter LeaderboardOutputAdapter,
	useCase UseCase,
) *LeaderboardControllerHTTP {
	handler := &LeaderboardControllerHTTP{
		inputAdapter:  inputAdapter,
		outputAdapter: outputAdapter,
		useCase:       useCase,
	}
	routeGroup.POST("/league/leaderboard", handler.postsStatsLeaderboard)
	return handler
}

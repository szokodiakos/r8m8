package leaderboard

import (
	"net/http"

	"github.com/labstack/echo"
)

// GetLeaderboardControllerHTTP struct
type GetLeaderboardControllerHTTP struct {
	inputAdapter  GetLeaderboardInputAdapter
	outputAdapter GetLeaderboardOutputAdapter
	useCase       UseCase
}

func (g *GetLeaderboardControllerHTTP) postsStatsLeaderboard(context echo.Context) error {
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

// NewGetLeaderboardControllerHTTP factory
func NewGetLeaderboardControllerHTTP(
	routeGroup *echo.Group,
	inputAdapter GetLeaderboardInputAdapter,
	outputAdapter GetLeaderboardOutputAdapter,
	useCase UseCase,
) *GetLeaderboardControllerHTTP {
	handler := &GetLeaderboardControllerHTTP{
		inputAdapter:  inputAdapter,
		outputAdapter: outputAdapter,
		useCase:       useCase,
	}
	routeGroup.POST("/league/leaderboard", handler.postsStatsLeaderboard)
	return handler
}

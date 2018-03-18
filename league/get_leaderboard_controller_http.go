package league

import (
	"net/http"

	"github.com/labstack/echo"
)

// GetLeaderboardControllerHTTP struct
type GetLeaderboardControllerHTTP struct {
	inputAdapter  GetLeaderboardInputAdapter
	outputAdapter GetLeaderboardOutputAdapter
	useCase       GetLeaderboardUseCase
}

func (g *GetLeaderboardControllerHTTP) postsStatsLeaderboard(context echo.Context) error {
	body := context.Get("parsedBody").(string)
	input, err := g.inputAdapter.Handle(body)
	if err != nil {
		return err
	}

	output, err := g.useCase.Handle(input)
	if err != nil {
		return err
	}

	response, err := g.outputAdapter.Handle(output)
	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, response)
}

// NewGetLeaderboardControllerHTTP factory
func NewGetLeaderboardControllerHTTP(
	slackGroup *echo.Group,
	inputAdapter GetLeaderboardInputAdapter,
	outputAdapter GetLeaderboardOutputAdapter,
	useCase GetLeaderboardUseCase,
) *GetLeaderboardControllerHTTP {
	handler := &GetLeaderboardControllerHTTP{
		inputAdapter:  inputAdapter,
		outputAdapter: outputAdapter,
		useCase:       useCase,
	}
	slackGroup.POST("/stats/leaderboard", handler.postsStatsLeaderboard)
	return handler
}

package leaderboard

import (
	"net/http"

	"github.com/labstack/echo"
)

// ControllerHTTP struct
type ControllerHTTP struct {
	inputAdapter  InputAdapter
	outputAdapter OutputAdapter
	useCase       UseCase
}

func (c *ControllerHTTP) postsStatsLeaderboard(context echo.Context) error {
	body := context.Get("parsedBody").(string)
	input, err := c.inputAdapter.Handle(body)
	if err != nil {
		return err
	}

	output, err := c.useCase.Handle(input)

	response, err := c.outputAdapter.Handle(output, err)
	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, response)
}

// NewControllerHTTP factory
func NewControllerHTTP(
	routeGroup *echo.Group,
	inputAdapter InputAdapter,
	outputAdapter OutputAdapter,
	useCase UseCase,
) *ControllerHTTP {
	handler := &ControllerHTTP{
		inputAdapter:  inputAdapter,
		outputAdapter: outputAdapter,
		useCase:       useCase,
	}
	routeGroup.POST("/league/leaderboard", handler.postsStatsLeaderboard)
	return handler
}

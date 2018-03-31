package stats

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

func (g *ControllerHTTP) postsStats(context echo.Context) error {
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
	routeGroup.POST("/stats", handler.postsStats)
	return handler
}

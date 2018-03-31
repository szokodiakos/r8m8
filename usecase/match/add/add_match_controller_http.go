package add

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

func (c *ControllerHTTP) postMatch(context echo.Context) error {
	body := context.Get("parsedBody").(string)
	input, err := c.inputAdapter.Handle(body)
	if err != nil {
		return c.handleOutput(context, Output{}, err)
	}

	output, err := c.useCase.Handle(input)

	return c.handleOutput(context, output, err)
}

func (c *ControllerHTTP) handleOutput(context echo.Context, output Output, err error) error {
	response, err := c.outputAdapter.Handle(output, err)
	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, response)
}

// NewAddMatchControllerHTTP factory
func NewAddMatchControllerHTTP(
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
	routeGroup.POST("/match/add", handler.postMatch)
	return handler
}

package match

import (
	"net/http"

	"github.com/labstack/echo"
)

// AddMatchControllerHTTP struct
type AddMatchControllerHTTP struct {
	inputAdapter  AddMatchInputAdapter
	outputAdapter AddMatchOutputAdapter
	useCase       AddMatchUseCase
}

func (a *AddMatchControllerHTTP) postMatch(context echo.Context) error {
	body := context.Get("parsedBody").(string)
	input, err := a.inputAdapter.Handle(body)
	if err != nil {
		return a.handleOutput(context, AddMatchOutput{}, err)
	}

	output, err := a.useCase.Handle(input)

	return a.handleOutput(context, output, err)
}

func (a *AddMatchControllerHTTP) handleOutput(context echo.Context, output AddMatchOutput, err error) error {
	response, err := a.outputAdapter.Handle(output, err)
	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, response)
}

// NewAddMatchControllerHTTP factory
func NewAddMatchControllerHTTP(
	routeGroup *echo.Group,
	inputAdapter AddMatchInputAdapter,
	outputAdapter AddMatchOutputAdapter,
	useCase AddMatchUseCase,
) *AddMatchControllerHTTP {
	handler := &AddMatchControllerHTTP{
		inputAdapter:  inputAdapter,
		outputAdapter: outputAdapter,
		useCase:       useCase,
	}
	routeGroup.POST("/match/add", handler.postMatch)
	return handler
}

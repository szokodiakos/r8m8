package undo

import (
	"net/http"

	"github.com/labstack/echo"
)

// UndoMatchControllerHTTP struct
type UndoMatchControllerHTTP struct {
	inputAdapter  UndoMatchInputAdapter
	outputAdapter UndoMatchOutputAdapter
	useCase       UndoMatchUseCase
}

func (a *UndoMatchControllerHTTP) postMatch(context echo.Context) error {
	body := context.Get("parsedBody").(string)
	input, err := a.inputAdapter.Handle(body)
	if err != nil {
		return a.handleOutput(context, UndoMatchOutput{}, err)
	}

	output, err := a.useCase.Handle(input)

	return a.handleOutput(context, output, err)
}

func (a *UndoMatchControllerHTTP) handleOutput(context echo.Context, output UndoMatchOutput, err error) error {
	response, err := a.outputAdapter.Handle(output, err)
	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, response)
}

// NewUndoMatchControllerHTTP factory
func NewUndoMatchControllerHTTP(
	routeGroup *echo.Group,
	inputAdapter UndoMatchInputAdapter,
	outputAdapter UndoMatchOutputAdapter,
	useCase UndoMatchUseCase,
) *UndoMatchControllerHTTP {
	handler := &UndoMatchControllerHTTP{
		inputAdapter:  inputAdapter,
		outputAdapter: outputAdapter,
		useCase:       useCase,
	}
	routeGroup.POST("/match/undo", handler.postMatch)
	return handler
}

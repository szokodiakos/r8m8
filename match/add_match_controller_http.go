package match

import (
	"net/http"

	"github.com/labstack/echo"
)

// AddMatchControllerHTTP struct
type AddMatchControllerHTTP struct {
	addMatchInputAdapter  AddMatchInputAdapter
	addMatchOutputAdapter AddMatchOutputAdapter
	addMatchUseCase       AddMatchUseCase
}

func (s *AddMatchControllerHTTP) postMatch(context echo.Context) error {
	body := context.Get("parsedBody").(string)
	input, err := s.addMatchInputAdapter.Handle(body)
	if err != nil {
		return err
	}

	output, err := s.addMatchUseCase.Handle(input)
	if err != nil {
		return err
	}

	response, err := s.addMatchOutputAdapter.Handle(output)
	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, response)
}

// NewAddMatchControllerHTTP factory
func NewAddMatchControllerHTTP(
	slackGroup *echo.Group,
	addMatchInputAdapter AddMatchInputAdapter,
	addMatchOutputAdapter AddMatchOutputAdapter,
	addMatchUseCase AddMatchUseCase,
) *AddMatchControllerHTTP {
	handler := &AddMatchControllerHTTP{
		addMatchInputAdapter:  addMatchInputAdapter,
		addMatchOutputAdapter: addMatchOutputAdapter,
		addMatchUseCase:       addMatchUseCase,
	}
	slackGroup.POST("/match", handler.postMatch)
	return handler
}

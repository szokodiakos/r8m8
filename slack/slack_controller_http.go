package slack

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/szokodiakos/r8m8/middleware"
)

// HTTPController struct
type HTTPController struct {
	service Service
}

func (c *HTTPController) postSlackMatch(context echo.Context) error {
	body := context.Get("parsedBody").(string)
	c.service.OpenMatchDialog(body)
	return context.JSON(http.StatusOK, nil)
}

// NewHTTPController creates a HTTPController
func NewHTTPController(e *echo.Echo, slackService Service) *HTTPController {
	handler := &HTTPController{
		service: slackService,
	}

	e.POST("/slack/match", handler.postSlackMatch, middleware.BodyParser())

	return handler
}

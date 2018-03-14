package echo

import (
	"net/http"

	"github.com/labstack/echo"
)

// ErrorHandlerMiddleware handles errors
func ErrorHandlerMiddleware(errorHandler ErrorHandler) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			if err := next(context); err != nil {
				response := errorHandler.HandleError(err)
				context.JSON(http.StatusOK, response)
			}
			return nil
		}
	}
}

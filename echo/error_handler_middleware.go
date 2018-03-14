package echo

import (
	"github.com/labstack/echo"
)

// ErrorHandlerMiddleware handles errors
func ErrorHandlerMiddleware(errorHandler ErrorHandler) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			if err := next(context); err != nil {
				code, response := errorHandler.HandleError(err)
				context.JSON(code, response)
			}
			return nil
		}
	}
}

package echo

import (
	"log"

	"github.com/labstack/echo"
)

// ErrorHandlerMiddleware handles errors
func ErrorHandlerMiddleware(errorHandler HTTPErrorHandler) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			if err := next(context); err != nil {
				log.Println("Unhandled Error", err)
				code, response := errorHandler.HandleError(err)
				context.JSON(code, response)
			}
			return nil
		}
	}
}

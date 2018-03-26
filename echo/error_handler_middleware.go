package echo

import (
	"github.com/labstack/echo"
	"github.com/szokodiakos/r8m8/logger"
)

// ErrorHandlerMiddleware handles errors
func ErrorHandlerMiddleware(errorHandler HTTPErrorHandler) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			if err := next(context); err != nil {
				logger.Get().Error("Unhandled Error", err)
				code, response := errorHandler.HandleError(err)
				context.JSON(code, response)
			}
			return nil
		}
	}
}

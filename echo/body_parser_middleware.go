package echo

import (
	"io/ioutil"

	"github.com/labstack/echo"
)

// BodyParser parses request body into "parsedBody" of Context
func BodyParser() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			body, err := ioutil.ReadAll(context.Request().Body)
			if err != nil {
				context.Logger().Fatal(err)
			}

			context.Set("parsedBody", string(body))
			return next(context)
		}
	}
}

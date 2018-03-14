package slack

import (
	"github.com/labstack/echo"
)

// TokenVerifier checks whether the verification token is right
func TokenVerifier(slackService Service) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			values := context.Get("parsedBody").(string)
			err := slackService.VerifyToken(values)
			if err != nil {
				return err
			}
			return next(context)
		}
	}
}

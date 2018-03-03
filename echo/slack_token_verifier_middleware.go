package echo

import (
	"github.com/labstack/echo"
	"github.com/szokodiakos/r8m8/echo/errors"
	"github.com/szokodiakos/r8m8/slack"
)

// SlackTokenVerifier checks whether the verification token is right
func SlackTokenVerifier(slackService slack.Service, verificationToken string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			values := context.Get("parsedBody").(string)
			requestValues := slackService.ParseRequestValues(values)
			token := requestValues.Token
			if token != verificationToken {
				return errors.NewInvalidVerificationTokenError()
			}
			return next(context)
		}
	}
}

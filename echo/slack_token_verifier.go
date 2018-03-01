package echo

import (
	"net/url"

	"github.com/labstack/echo"
	"github.com/szokodiakos/r8m8/echo/errors"
)

// SlackTokenVerifier checks whether the verification token is right
func SlackTokenVerifier(verificationToken string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			body := context.Get("parsedBody").(string)
			parsedValues, _ := url.ParseQuery(body)
			token := parsedValues.Get("token")
			if token != verificationToken {
				return errors.NewInvalidVerificationTokenError()
			}
			return next(context)
		}
	}
}

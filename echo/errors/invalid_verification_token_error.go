package errors

import (
	"net/http"

	"github.com/labstack/echo"
)

// NewInvalidVerificationTokenError factory
func NewInvalidVerificationTokenError() error {
	return echo.NewHTTPError(http.StatusUnauthorized, "Invalid verification token")
}

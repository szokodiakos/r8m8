package errors

import (
	"net/http"

	"github.com/labstack/echo"
)

// NewInvalidVerificationTokenError creates an error
func NewInvalidVerificationTokenError() error {
	return echo.NewHTTPError(http.StatusUnauthorized, "Invalid verification token")
}

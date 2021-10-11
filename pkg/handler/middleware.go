package controller

import (
	"github.com/labstack/echo/v4"
)

func JWTErrorChecker(err error, c echo.Context) error {
	return echo.ErrUnauthorized
}

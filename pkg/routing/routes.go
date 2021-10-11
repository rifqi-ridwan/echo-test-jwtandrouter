package routing

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterRouting(e *echo.Group) {
	e.GET("test", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})
}

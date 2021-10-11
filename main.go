package main

import (
	h "echo-test-jwtandrouter/pkg/handler"
	"echo-test-jwtandrouter/pkg/routing"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/login", h.Login)

	config := middleware.JWTConfig{
		Claims:                  &jwt.StandardClaims{},
		SigningKey:              []byte("secret"),
		ErrorHandlerWithContext: h.JWTErrorChecker,
	}

	r := e.Group("/", middleware.JWTWithConfig(config))
	routing.RegisterRouting(r)

	e.Logger.Fatal(e.Start(":1323"))
}

package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var IsLoggedIn = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte("secret"),
	Skipper: func(c echo.Context) bool {
		if c.Request().RequestURI == "/signin" {
			return true
		}
		return false
	},
})

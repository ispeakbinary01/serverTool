package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

var routes = map[string][]string {
	"admin": {
		"/inventories/software",
		"/inventories/ssh",
		"/inventories/servers",
		"/users",
	},
	"moderator": {
		"/inventories/software",
		"/inventories/ssh",
		"/inventories/servers",
		"/users",
	},
	"user": {
		"/inventories/software",
		"/inventories/ssh",
		"/inventories/servers",
		"/users",
	},
}



var IsLoggedIn = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte("secret"),
	Skipper: func(c echo.Context) bool {
		if c.Request().RequestURI == "/signin" {
			return true
		}
		return false
	},
})

func AdminRoutes(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		u := c.Get("user").(*jwt.Token)
		claims := u.Claims.(jwt.MapClaims)
		for _, item := range routes[claims["role"].(string)] {
			if c.Request().RequestURI == item && claims["role"] != "admin" {
				return c.JSON(http.StatusUnauthorized, "Role not suitable for function.")
			}
		}
		return next(c)
	}
}


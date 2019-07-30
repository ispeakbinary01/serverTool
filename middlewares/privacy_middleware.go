package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
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

func RoutesPrivileges(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Request().RequestURI == "/signin" {
			return next(c)
		}
		pathMethod := c.Request().Method + " " + c.Path()
		if temp := c.Get("user"); temp != nil {
			u := temp.(*jwt.Token)
			claims := u.Claims.(jwt.MapClaims)
			for _, item := range  Routes[pathMethod] {
				if claims["role"] == item {
					return next(c)
				}
			}
		}
		return c.JSON(http.StatusUnauthorized, "Role not suitable for function.")
	}
}


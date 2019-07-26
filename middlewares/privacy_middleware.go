package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

var routes = []string {
	"/inventories/software",
	"/inventories/ssh",
	"/inventories/servers",
	"/users",
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
		//c.Response().Header().Set(echo.HeaderServer, "Echo/3.0")
		//return next(c)
		u := c.Get("user").(*jwt.Token)
		claims := u.Claims.(jwt.MapClaims)
		for _, item := range routes {
			if c.Request().RequestURI == item && claims["role"] != "Admin" {
				return c.JSON(http.StatusUnauthorized, "Role not suitable for function.")
			}
		}
		return next(c)
	}
}


//var AdminRoutes = middleware.JWTWithConfig(middleware.JWTConfig{
//	SigningKey: []byte("secret"),
//	Skipper: func(c echo.Context) bool {
//		u := c.Get("user").(*jwt.Token)
//		claims := u.Claims.(jwt.MapClaims)
//		for _, item := range routes {
//			if c.Request().RequestURI == item && claims["role"] != "Admin" {
//				fmt.Println("NO!")
//				return false
//			}
//		}
//		return true
//	},
//})
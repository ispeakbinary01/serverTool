package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

var routes = map[string][]string {
	"GET /inventories/software": {
		"admin",
		"moderator",
		"user",
	},
	"GET /inventories/software/:id": {
		"admin",
		"moderator",
		"user",
	},
	"POST /inventories/software": {
		"admin",
		"moderator",
	},
	"PUT /inventories/software/:id": {
		"admin",
		"moderator",
	},
	"DELETE /inventories/software/:id": {
		"admin",
	},
	"GET /inventories/ssh": {
		"admin",
		"moderator",
		"user",
	},
	"GET /inventories/ssh/:id": {
		"admin",
		"moderator",
		"user",
	},
	"POST /inventories/ssh": {
		"admin",
		"moderator",
	},
	"PUT /inventories/ssh/:id": {
		"admin",
		"moderator",
	},
	"DELETE /inventories/ssh/:id": {
		"admin",
	},
	"GET /users": {
		"admin",
		"moderator",
	},
	"GET /users/:id": {
		"admin",
		"moderator",
	},
	"POST /users": {
		"admin",
	},
	"PUT /users/:id": {
		"admin",
	},
	"PATCH /users/:id": {
		"admin",
	},
	"DELETE /users/:id": {
		"admin",
	},
	"GET /inventories/serversSSH/": {
		"admin",
		"moderator",
		"user",
	},
	"GET /inventories/serversSoftware/:id": {
		"admin",
		"moderator",
		"user",
	},
	"GET /serversByUser": {
		"admin",
		"moderator",
		"user",
	},
	"POST /inventories/servers": {
		"admin",
		"moderator",
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

func RoutesPrivileges(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Request().RequestURI == "/signin" {
			return next(c)
		}
		pathMethod := c.Request().Method + " " + c.Path()
		if temp := c.Get("user"); temp != nil {
			u := temp.(*jwt.Token)
			claims := u.Claims.(jwt.MapClaims)
			for _, item := range  routes[pathMethod] {
				if claims["role"] == item {
					return next(c)
				}
			}
		}
		return c.JSON(http.StatusUnauthorized, "Role not suitable for function.")
	}
}


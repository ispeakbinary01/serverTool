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

//var AdminRoutes = middleware.JWTWithConfig(middleware.JWTConfig{
//	SigningKey: []byte("secret"),
//	Skipper: func(c echo.Context) bool {
//		var routes = make([]string, 4)
//		routes = append(routes, "/inventories/software")
//		routes = append(routes, "/inventories/ssh")
//		routes = append(routes, "/inventories/servers")
//		routes = append(routes, "/users")
//		u := c.Get("users").(*jwt.Token)
//		fmt.Println("UUUUUUUUU PRINTED")
//		fmt.Println(u)
//		claims := u.Claims.(jwt.MapClaims)
//		fmt.Println("CLAIMS PRINTED!!!!!")
//		fmt.Println(claims)
//		for _, item := range routes {
//			if c.Request().RequestURI == item && claims["role"] != "Admin" {
//				fmt.Println("NO!")
//				return false
//			}
//		}
//		return true
//	},
//})
package main

import (
	"github.com/ispeakbinary01/serverTool/handlers"
	"github.com/ispeakbinary01/serverTool/middlewares"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)


func main() {
	api := echo.New()
	api.Use(middlewares.IsLoggedIn)
	api.POST("/signin", handlers.Signin)
	api.Use(middlewares.RoutesPrivileges)
	api.Use(middleware.Secure())
	api.GET("/inventories/software", handlers.GetAllSoftware)               // Works
	api.GET("/inventories/software/:id", handlers.GetSoftwareByID)          // Works
	api.POST("/inventories/software", handlers.PostSoftware)                // Works
	api.PUT("/inventories/software/:id", handlers.UpdateSoftware)           // Works
	api.DELETE("/inventories/software/:id", handlers.DeleteSoftware)        // Works but returns Locker{}
	api.GET("/inventories/ssh", handlers.GetSSHs)                           // Works
	api.GET("/inventories/ssh/:id", handlers.GetSSH)                        // Works
	api.POST("/inventories/ssh", handlers.PostSSH)                          // Works
	api.PUT("/inventories/ssh/:id", handlers.UpdateSSH)                     // Works
	api.DELETE("/inventories/ssh/:id", handlers.DeleteSSH)                  // Works but returns null
	api.GET("/users", handlers.GetUsers)                                    // Works
	api.GET("/users/:id", handlers.GetUser)                                 // Works
	api.POST("/users", handlers.PostUser)                                   // Works
	api.PUT("users/:id", handlers.UpdateUser)                               // Works
	api.DELETE("/users/:id", handlers.DeleteUser)                           // Works but returns null
	api.GET("/inventories/serversSSH/:id", handlers.GetServerSSH)           // Works
	api.GET("/inventories/serversSoftware/:id", handlers.GetServerSoftware) // Works
	api.GET("/serversByUser", handlers.GetServersByUser) 					// Works
	api.POST("/inventories/servers", handlers.PostServer)                   // Works

	api.Start(":8080")

}

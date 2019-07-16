package main

import (
	"github.com/ispeakbinary01/serverTool/handlers"
	"github.com/ispeakbinary01/serverTool/middlewares"
	"github.com/labstack/echo/v4"
)

func main() {
	api := echo.New()
	api.POST("/signin", handlers.Signin)
	api.POST("/inventories/software", handlers.PostSoftware, middlewares.IsLoggedIn)                          // Works
	api.GET("/inventories/software", handlers.GetAllSoftware, middlewares.IsLoggedIn) 						// Works
	api.GET("/inventories/software/:id", handlers.GetSoftwareByID, middlewares.IsLoggedIn)                    // Works
	api.DELETE("/inventories/software/:id", handlers.DeleteSoftware, middlewares.IsLoggedIn)                  // Works but returns Locker{}
	api.PUT("/inventories/software/:id", handlers.UpdateSoftware, middlewares.IsLoggedIn)                     // Works
	api.POST("/inventories/ssh", handlers.PostSSH, middlewares.IsLoggedIn)                                    // Works
	api.GET("/inventories/ssh", handlers.GetSSHs, middlewares.IsLoggedIn)                                     // Works
	api.GET("/inventories/ssh/:id", handlers.GetSSH, middlewares.IsLoggedIn)                                  // Works
	api.DELETE("/inventories/ssh/:id", handlers.DeleteSSH, middlewares.IsLoggedIn)                            // Works but returns null
	api.PUT("/inventories/ssh/:id", handlers.UpdateSSH, middlewares.IsLoggedIn)                               // Works
	api.POST("/users", handlers.PostUser)                                             						// Works
	api.GET("/users", handlers.GetUsers, middlewares.IsLoggedIn)                                              // Works
	api.GET("/users/:id", handlers.GetUser, middlewares.IsLoggedIn)                                           // Works
	api.PUT("users/:id", handlers.UpdateUser, middlewares.IsLoggedIn)                                         // Works
	api.DELETE("/users/:id", handlers.DeleteUser, middlewares.IsLoggedIn)                                     // Works but returns null
	api.POST("/inventories/servers", handlers.PostServer, middlewares.IsLoggedIn)                             // Works
	api.GET("/inventories/serversSSH/:id", handlers.GetServerSSH, middlewares.IsLoggedIn)                     // Works
	api.GET("/inventories/serversSoftware/:id", handlers.GetServerSoftware, middlewares.IsLoggedIn)           // Works

	api.Start(":8080")

}

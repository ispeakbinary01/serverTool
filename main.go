package main

import (
	"github.com/ispeakbinary01/serverTool/handlers"
	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	api := echo.New()

	api.POST("/inventories/software", handlers.PostSoftware)
	api.GET("/inventories/software", handlers.GetAllSoftware)
	api.GET("/inventories/software/:id", handlers.GetSoftwareByID)
	api.DELETE("/inventories/software/:id", handlers.DeleteSoftware)
	api.PUT("/inventories/software/:id", handlers.UpdateSoftware)
	api.POST("/inventories/ssh", handlers.PostSSH)
	api.GET("/inventories/ssh", handlers.GetSSHs)
	api.GET("/inventories/ssh/:id", handlers.GetSSH)
	api.DELETE("/inventories/ssh/:id", handlers.DeleteSSH)
	api.PUT("/inventories/ssh/:id", handlers.UpdateSSH)
	api.POST("/users", handlers.PostUser)
	api.GET("/users", handlers.GetUsers)
	api.GET("/users/:id", handlers.GetUser)
	api.PUT("users/:id", handlers.UpdateUser)
	api.DELETE("/users/:id", handlers.DeleteUser)
	api.POST("/inventories/servers", handlers.PostServer)

	api.Start(":8080")

}

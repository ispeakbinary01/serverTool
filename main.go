package main

import (
	"github.com/ispeakbinary01/serverTool/handlers"
	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	api := echo.New()

	api.POST("/inventory/software", handlers.PostSoftware)
	api.GET("/inventory/software", handlers.GetAllSoftwarae)
	api.GET("/inventory/software/:id", handlers.GetSoftwareByID)
	api.DELETE("/inventory/software/:id", handlers.DeleteSoftware)
	//api.PUT("/inventory/software/:id", handlers.UpdateSoftware)
	api.POST("/inventory/ssh", handlers.PostSSH)
	api.GET("/inventory/ssh", handlers.GetSSHs)
	api.GET("/inventory/ssh/:id", handlers.GetSSH)
	api.DELETE("/inventory/ssh/:id", handlers.DeleteSSH)
	api.POST("/user", handlers.PostUser)
	api.GET("/users", handlers.GetUsers)
	api.GET("/user/:id", handlers.GetUser)
	api.DELETE("/user/:id", handlers.DeleteUser)

	api.Start(":8080")

}

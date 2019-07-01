package main

import (
	"github.com/ispeakbinary01/serverTool/handlers"
	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	api := echo.New()

	api.POST("/inventory/server", handlers.PostServer)
	api.POST("/inventory/software", handlers.PostSoftware)
	api.GET("/inventory/software", handlers.GetSoftware)
	api.GET("/inventory/software/:id", handlers.GetSoftwareID)
	api.DELETE("/inventory/software/:id", handlers.DeleteSoftware)
	api.PUT("/inventory/software/:id", handlers.UpdateSoftware)

	api.Start(":8080")

}

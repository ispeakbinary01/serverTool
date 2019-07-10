package handlers

import (
	"fmt"
	"github.com/ispeakbinary01/serverTool/pkg/server"
	"github.com/labstack/echo/v4"
)

// PostServer ...
func PostServer(c echo.Context) error {
	s := server.NewServer()
	if err := c.Bind(s); err != nil {
		fmt.Println(err)
		return err
	}
	seID, err := s.CreateServer()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return c.JSON(201, seID)
}

// GetServerSSH ...
func GetServerSSH(c echo.Context) error {
	requestID := c.Param("id")
	response, err := server.GetServerSSH(requestID)
	if err != nil {
		return err
	}

	return c.JSON(200, response)
}

// GetServerSoftware ...
func GetServerSoftware(c echo.Context) error {
	requestID := c.Param("id")
	response, err := server.GetserverSoftware(requestID)
	if err != nil {
		return err
	}

	return c.JSON(200, response)
}



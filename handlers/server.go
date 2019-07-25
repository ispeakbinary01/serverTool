package handlers

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/ispeakbinary01/serverTool/pkg/server"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)


// PostServer ...
func PostServer(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	s := server.NewServer()
	if err := c.Bind(s); err != nil {
		log.Printf("%s", err.Error())
	}
	seID, err := s.CreateServer(claims["id"])
	valErr := s.Validate()
	if valErr != nil {
		return c.JSON(http.StatusBadRequest, valErr)
	}
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
		log.Printf("%s", err.Error())
	}

	return c.JSON(200, response)
}

// GetServerSoftware ...
func GetServerSoftware(c echo.Context) error {
	requestID := c.Param("id")
	response, err := server.GetserverSoftware(requestID)
	if err != nil {
		log.Printf("%s", err.Error())
	}

	return c.JSON(200, response)
}





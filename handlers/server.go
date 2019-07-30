package handlers

import (
	"encoding/json"
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
		log.Printf("%s", err)
		return err
	}
	valErr := s.Validate()
	if valErr != nil {
		return c.JSON(http.StatusBadRequest, valErr)
	}
	seID, err := s.CreateServer(claims["id"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	return c.JSON(201, seID)
}

// GetServers ...
func GetServers(c echo.Context) error {
	response, err := server.GetAllServers()
	if err != nil {
		log.Printf("%s", err)
	}

	return c.JSON(200, response)
}


// GetServer ...
func GetServer (c echo.Context) error {
	requestID := c.Param("id")
	s, err := server.GetServerByID(requestID)
	if err != nil {
		log.Printf("%s", err)
		return err
	}
	return c.JSON(200, s)
}

// GetServerSSH ...
func GetServerSSH(c echo.Context) error {
	requestID := c.Param("id")
	response, err := server.GetServerSSH(requestID)
	if err != nil {
		log.Printf("%s", err)
	}

	return c.JSON(200, response)
}

// GetServerSoftware ...
func GetServerSoftware(c echo.Context) error {
	requestID := c.Param("id")
	response, err := server.GetServerSoftware(requestID)
	if err != nil {
		log.Printf("%s", err)
	}

	return c.JSON(200, response)
}

// AddServerToUser
func AddServerToUser(c echo.Context) error {
	var data map[string]int
	err := json.NewDecoder(c.Request().Body).Decode(&data)
	if err != nil {
		log.Printf("%s", err)
		return err
	}
	err2 := server.AddServerToUser(data["user_id"], data["server_id"])
	if err2 != nil {
		log.Printf("%s", err2)
		return err2
	}
	return c.JSON(200, "Relation created")
}





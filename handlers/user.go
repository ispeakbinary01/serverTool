package handlers

import (
	"fmt"
	"github.com/ispeakbinary01/serverTool/db"
	"github.com/ispeakbinary01/serverTool/pkg/user"
	"github.com/labstack/echo/v4"
	"net/http"
)

// PostUser ...
func PostUser(c echo.Context) error {
	u := user.NewUser()
	if err := c.Bind(u); err != nil {
		return err
	}
	uid, err := u.CreateUser()
	if err != nil {
		return err
	}
	u.ID = uid
	return c.JSON(201, u)
}

// GetUsers ...
func GetUsers(c echo.Context) error {
	response := user.GetAllUsers()

	return c.JSON(200, response)
}

// GetUser ...
func GetUser(c echo.Context) error {
	requestID := c.Param("id")
	u := user.GetUserByID(requestID)

	return c.JSON(200, u)
}

// DeleteUser ...
func DeleteUser(c echo.Context) error {
	requestID := c.Param("id")
	u := user.DeleteUser(requestID)

	return c.JSON(200, "Deleted" + u)
}
 
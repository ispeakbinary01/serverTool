package handlers

import (
	"github.com/ispeakbinary01/serverTool/pkg/user"
	"github.com/labstack/echo/v4"
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
	response, err := user.GetAllUsers()
	if err != nil {
		return err
	}

	return c.JSON(200, response)
}

// GetUser ...
func GetUser(c echo.Context) error {
	requestID := c.Param("id")
	u, err := user.GetUserByID(requestID)
	if err != nil {
		return err
	}

	return c.JSON(200, u)
}

// DeleteUser ...
func DeleteUser(c echo.Context) error {
	requestID := c.Param("id")
	u := user.DeleteUser(requestID)

	return c.JSON(200, u)
}

// UpdateUser ...
func UpdateUser(c echo.Context) error {
	requestID := c.Param("id")
	u := user.NewUser()
	if err := c.Bind(u); err != nil {
		return err
	}
	uid, err := u.UpdateUser(requestID)
	if err != nil {
		return err
	}
	return c.JSON(201, uid)

}
 
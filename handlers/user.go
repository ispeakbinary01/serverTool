package handlers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/ispeakbinary01/serverTool/pkg/user"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)



// PostUser ...
func PostUser(c echo.Context) error {
	u := user.NewUser()
	if err := c.Bind(u); err != nil {
		log.Printf("%s", err)
		return err
	}
	valErr := u.Validate()
	if valErr != nil {
		return c.JSON(http.StatusBadRequest, valErr)
	}
	uid, err := u.CreateUser()
	if err != nil {
		log.Printf("%s", err)
		return err
	}
	u.ID = uid

	return c.JSON(201, u)
}

// GetUsers ...
func GetUsers(c echo.Context) error {
	response, err := user.GetAllUsers()
	if err != nil {
		log.Printf("%s", err)
		return err
	}

	return c.JSON(200, response)
}

// GetUser ...
func GetUser(c echo.Context) error {
	requestID := c.Param("id")
	u, err := user.GetUserByID(requestID)
	if err != nil {
		log.Printf("%s", err)
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
		log.Printf("%s", err)
		return err
	}
	uid, err := u.UpdateUser(requestID)
	if err != nil {
		log.Printf("%s", err)
		return err
	}
	return c.JSON(201, uid)

}

// GetServersByUser ...
func GetServersByUser(c echo.Context) error {
	u := c.Get("user").(*jwt.Token)
	claims := u.Claims.(jwt.MapClaims)
	response, err := user.GetServersByUser(claims["id"])
	if err != nil {
		return c.JSON(402, response)
	}
	return c.JSON(200, response)
}

// PatchRole ...
func PatchRole(c echo.Context) error {
	requestID := c.Param("id")
	requestEmail := c.Param("email")
	u := user.NewUser()
	if err := c.Bind(u); err != nil {
		log.Printf("%s", err)
		return err
	}
	_, err2 := u.PatchRole(requestID)
	if err2 != nil {
		log.Printf("%s", err2)
		return err2
	}
	return c.JSON(201, "Role changed on user " + requestEmail)

}
 
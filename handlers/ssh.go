package handlers

import (
	"fmt"
	"github.com/ispeakbinary01/serverTool/pkg/user"
	"net/http"
	"strconv"
	"github.com/ispeakbinary01/serverTool/pkg/server/ssh"
	"github.com/ispeakbinary01/serverTool/db"
	"github.com/labstack/echo/v4"
)

// PostSSH ...
func PostSSH(c echo.Context) error {
	ssh := ssh.NewSSH()
	if err := c.Bind(ssh); err != nil {
		return err
	}
	sshId, err := ssh.CreateSSH()
	if err != nil {
		return err
	}
	ssh.ID = sshId
	return c.JSON(201, ssh)
}

// GetSSHs ...
func GetSSHs(c echo.Context) error {
	response, err := ssh.GetAllSSHs()
	if err != nil {
		return err
	}

	return c.JSON(200, response)
}

// GetSSH ...
func GetSSH(c echo.Context) error {
	requestID := c.Param("id")
	ssh := ssh.GetSShByID(requestID)

	return c.JSON(200, ssh)
}

// DeleteSSH ...
func DeleteSSH(c echo.Context) error {
	requestID := c.Param("id")
	ssh := ssh.DeleteSSH(requestID)

	return c.JSON(200, ssh)
}

// UpdateSSH ...
// func UpdateSSH(c echo.Context) error {
// 	updatedSSH := new(structs.SSH)
// 	requestID, _ := strconv.Atoi(c.Param("id"))
// 	if err := c.Bind(updatedSSH); err != nil {
// 		return err
// 	}
// 	if PutHelper(requestID) {
// 		sql := "UPDATE ssh SET "
// 	}
// }

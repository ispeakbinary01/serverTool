package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"github.com/ispeakbinary01/serverTool/pkg/ssh"
	"github.com/ispeakbinary01/serverTool/db"
	"github.com/labstack/echo/v4"
)

// PostSSH ...
func PostSSH(c echo.Context) error {
	ssh := ssh.NewSSH()
	if err := c.Bind(ssh); err != nil {
		return err
	}
	sshid, err := ssh.CreateSSH()
	if err != nil {
		return err
	}
	ssh.ID = sshid
	return c.JSON(201, ssh)
}

// GetSSHs ...
func GetSSHs(c echo.Context) error {
	ssh := ssh.SSHsQuery()

	return c.JSON(http.StatusOK, ssh)

}

// GetSSH ...
func GetSSH(c echo.Context) error {
	requestedID := c.Param("id")
	ssh := ssh.SSHQuery(requestedID)
	return c.JSON(http.StatusOK, ssh)
}

// DeleteSSH ...
func DeleteSSH(c echo.Context) error {
	requestID := c.Param("id")
	ssh := ssh.DeleteSSHQuery(requestID)

	return c.JSON(http.StatusOK, ssh)
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

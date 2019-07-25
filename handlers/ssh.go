package handlers

import (
	"github.com/ispeakbinary01/serverTool/pkg/server/ssh"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

// PostSSH ...
func PostSSH(c echo.Context) error {
	ssh := ssh.NewSSH()
	if err := c.Bind(ssh); err != nil {
		log.Printf("%s", err)
	}
	sshId, err := ssh.CreateSSH()
	valErr := ssh.Validate()
	if valErr != nil {
		return c.JSON(http.StatusBadRequest, valErr)
	}
	if err != nil {
		log.Printf("%s", err)
		return err
	}
	ssh.ID = sshId
	return c.JSON(201, ssh)
}

// GetSSHs ...
func GetSSHs(c echo.Context) error {
	response, err := ssh.GetAllSSHs()
	if err != nil {
		log.Printf("%s", err)
		return err
	}

	return c.JSON(200, response)
}

// GetSSH ...
func GetSSH(c echo.Context) error {
	requestID := c.Param("id")
	ssh, err := ssh.GetSShByID(requestID)
	if err != nil {
		log.Printf("%s", err)
		return err
	}
	return c.JSON(200, ssh)
}

// DeleteSSH ...
func DeleteSSH(c echo.Context) error {
	requestID := c.Param("id")
	delSsh := ssh.DeleteSSH(requestID)

	return c.JSON(200, delSsh)
}

// UpdateSSH ...
func UpdateSSH(c echo.Context) error {
	requestID := c.Param("id")
	ssh := ssh.NewSSH()
	if err := c.Bind(ssh); err != nil {
		log.Printf("%s", err)
		return err
	}
	sshid, err := ssh.UpdateSSH(requestID)
	if err != nil {
		log.Printf("%s", err)
		return err
	}
	return c.JSON(201, sshid)

}

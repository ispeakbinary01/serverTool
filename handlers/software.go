package handlers

import (
	"github.com/ispeakbinary01/serverTool/pkg/server/software"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

// PostSoftware ...
func PostSoftware(c echo.Context) error {
	s := software.NewSoftware()
	if err := c.Bind(s); err != nil {
		log.Printf("%s", err)
	}
	swID, err := s.CreateSoftware()
	valErr := s.Validate()
	if valErr != nil {
		return c.JSON(http.StatusBadRequest, valErr)
	}
	if err != nil {
		log.Printf("%s", err)
	}
	s.ID = swID
	return c.JSON(201, swID)
}

// GetAllSoftware ...
func GetAllSoftware(c echo.Context) error {
	response, err := software.GetAllSoftware()
	if err != nil {
		log.Printf("%s", err)
	}

	return c.JSON(200, response)
}

// GetSoftwareByID ...
func GetSoftwareByID(c echo.Context) error {
	requestID := c.Param("id")
	s, err := software.GetSoftwareByID(requestID)
	if err != nil {
		log.Printf("%s", err)
	}
	return c.JSON(200, s)
}

// DeleteSoftware ...
func DeleteSoftware(c echo.Context) error {
	requestID := c.Param("id")
	s, err := software.DeleteSoftware(requestID)
	if err != nil {
		log.Printf("%s", err)
	}

	return c.JSON(200, s)
}

// UpdateSoftware ...
func UpdateSoftware(c echo.Context) error {
	requestID := c.Param("id")
	sw := software.NewSoftware()
	if err := c.Bind(sw); err != nil {
		log.Printf("%s", err)
	}
	swid, err := sw.UpdateSoftware(requestID)
	if err != nil {
		log.Printf("%s", err)
	}
	return c.JSON(201, swid)

}

package handlers

import (
	"fmt"
	"github.com/ispeakbinary01/serverTool/pkg/server/software"
	"github.com/labstack/echo/v4"
)

// PutHelper ...
//func PutHelper(id int) bool {
//	var count int
//
//	row := db.Get().QueryRow("SELECT COUNT(*) FROM software")
//	err := row.Scan(&count)
//	if err != nil {
//		panic(err)
//	}
//	if id < count {
//		return true
//	}
//
//	return false
//}

// PostSoftware ...
func PostSoftware(c echo.Context) error {
	s := software.NewSoftware()
	if err := c.Bind(s); err != nil {
		return err
	}
	swID, err := s.CreateSoftware()
	if err != nil {
		return err
	}
	s.ID = swID
	return c.JSON(201, swID)
}

// GetAllSoftware ...
func GetAllSoftware(c echo.Context) error {
	response, err := software.GetAllSoftware()
	if err != nil {
		return err
	}

	return c.JSON(200, response)
}

// GetSoftwareByID ...
func GetSoftwareByID(c echo.Context) error {
	requestID := c.Param("id")
	s, err := software.GetSoftwareByID(requestID)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return c.JSON(200, s)
}

// DeleteSoftware ...
func DeleteSoftware(c echo.Context) error {
	requestID := c.Param("id")
	s, err := software.DeleteSoftware(requestID)
	if err != nil {
		return err
	}

	return c.JSON(200, s)
}

// UpdateSoftware ...
func UpdateSoftware(c echo.Context) error {
	requestID := c.Param("id")
	sw := software.NewSoftware()
	if err := c.Bind(sw); err != nil {
		return err
	}
	swid, err := sw.UpdateSoftware(requestID)
	if err != nil {
		return err
	}
	return c.JSON(201, swid)

}

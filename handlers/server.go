package handlers

import (
	"fmt"
	"github.com/ispeakbinary01/serverTool/pkg/server"
	"github.com/labstack/echo/v4"
)

// PostServer ...
func PostServer(c echo.Context) error {
	s := server.NewServer()
	if err := c.Bind(s); err != nil {
		fmt.Println(err)
		return err
	}
	seID, err := s.CreateServer()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return c.JSON(201, seID)
}

//// GetAllSoftware ...
//func GetAllSoftware(c echo.Context) error {
//	response, err := software.GetAllSoftware()
//	if err != nil {
//		return err
//	}
//
//	return c.JSON(200, response)
//}
//
//// GetSoftwareByID ...
//func GetSoftwareByID(c echo.Context) error {
//	requestID := c.Param("id")
//	s, err := software.GetSoftwareByID(requestID)
//	if err != nil {
//		fmt.Println(err)
//		return err
//	}
//	return c.JSON(200, s)
//}
//
//// DeleteSoftware ...
//func DeleteSoftware(c echo.Context) error {
//	requestID := c.Param("id")
//	s := software.DeleteSoftware(requestID)
//
//	return c.JSON(200, s)
//}
//
//// UpdateSoftware ...
//func UpdateSoftware(c echo.Context) error {
//	requestID := c.Param("id")
//	sw := software.NewSoftware()
//	if err := c.Bind(sw); err != nil {
//		return err
//	}
//	swid, err := sw.UpdateSoftware(requestID)
//	if err != nil {
//		return err
//	}
//	return c.JSON(201, swid)
//
//}



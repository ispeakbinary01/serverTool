package handlers

import (
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
func GetAllSoftwarae(c echo.Context) error {
	response, err := software.GetAllSoftware()
	if err != nil {
		return err
	}

	return c.JSON(200, response)
}

// GetSoftwareByID ...
func GetSoftwareByID(c echo.Context) error {
	requestID := c.Param("id")
	s := software.GetSoftwareByID(requestID)

	return c.JSON(200, s)
}

// DeleteSoftware ...
func DeleteSoftware(c echo.Context) error {
	requestID := c.Param("id")
	s := software.DeleteSoftware(requestID)

	return c.JSON(200, s)
}

// UpdateSoftware ...
//func UpdateSoftware(c echo.Context) error {
//	updatedSoftware := new(structs.Software)
//	requestID, _ := strconv.Atoi(c.Param("id"))
//	if err := c.Bind(updatedSoftware); err != nil {
//		return err
//	}
//	if PutHelper(requestID) {
//		sql := "UPDATE software SET name = ?, version = ? WHERE id = ?"
//		stmt, err := db.Get().Prepare(sql)
//		if err != nil {
//			panic(err)
//		}
//		_, err2 := stmt.Exec(updatedSoftware.Name, updatedSoftware.Version, requestID)
//		if err2 != nil {
//			panic(err2)
//		}
//
//		return c.JSON(http.StatusAccepted, "Updated!")
//	}
//	sql := "INSERT INTO software(name, version) VALUES(?, ?)"
//	stmt, err := db.Get().Prepare(sql)
//
//	if err != nil {
//		fmt.Print(err.Error())
//	}
//
//	defer stmt.Close()
//
//	result, err2 := stmt.Exec(updatedSoftware.Name, updatedSoftware.Version)
//
//	if err2 != nil {
//		panic(err2)
//	}
//
//	fmt.Println(result.LastInsertId())
//
//	return c.JSON(http.StatusCreated, updatedSoftware.Name)
//}

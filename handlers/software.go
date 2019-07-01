package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ispeakbinary01/serverTool/db"
	"github.com/ispeakbinary01/serverTool/structs"
	"github.com/labstack/echo/v4"
)

// PostSoftware ...
func PostSoftware(c echo.Context) error {
	software := new(structs.Software)
	if err := c.Bind(software); err != nil {
		return err
	}

	sql := "INSERT INTO software(name, version) VALUES(?, ?)"
	stmt, err := db.Get().Prepare(sql)

	if err != nil {
		fmt.Println("=======ERROR 1===========")
		fmt.Println(err)
	}

	defer stmt.Close()

	res, err2 := stmt.Exec(software.Name, software.Version)

	if err2 != nil {
		fmt.Println("=======ERROR 2===========")
		fmt.Println(err2)
	}

	fmt.Println(res.LastInsertId())

	return c.JSON(http.StatusCreated, software.Name)
}

// GetSoftware ...
func GetSoftware(c echo.Context) error {
	var id int
	var name string
	var version int

	err, _ := db.Get().Query("SELECT id, name, version FROM software")
	for err.Next() {
		err.Scan(&id, &name, &version)
		fmt.Println(strconv.Itoa(id) + ": " + name + " " + strconv.Itoa(version))
	}

	if err != nil {
		panic(err)
	}

	response := structs.Software{ID: id, Name: name, Version: version}

	return c.JSON(http.StatusOK, response)
}

// GetSoftwareID ...
func GetSoftwareID(c echo.Context) error {
	requestID := c.Param("id")
	var id int
	var name string
	var version int

	err := db.Get().QueryRow("SELECT id, name, version FROM software WHERE id = ?", requestID)
	err.Scan(&id, &name, &version)

	response := structs.Software{ID: id, Name: name, Version: version}

	return c.JSON(http.StatusOK, response)

}

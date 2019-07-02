package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ispeakbinary01/serverTool/db"
	"github.com/ispeakbinary01/serverTool/structs"
	"github.com/labstack/echo/v4"
)

// PutHelper ...
func PutHelper(id int) bool {
	var count int

	row := db.Get().QueryRow("SELECT COUNT(*) FROM software")
	err := row.Scan(&count)
	if err != nil {
		panic(err)
	}
	if id < count {
		return true
	}

	return false
}

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

// DeleteSoftware ...
func DeleteSoftware(c echo.Context) error {
	requestID := c.Param("id")
	requestName := c.Param("name")
	sql := "DELETE FROM software WHERE id = ?"
	stmt, err := db.Get().Prepare(sql)
	if err != nil {
		panic(err)
	}

	_, err2 := stmt.Exec(requestID)
	if err2 != nil {
		panic(err2)
	}

	return c.JSON(http.StatusOK, "Deleted software with name "+requestName)
}

// UpdateSoftware ...
func UpdateSoftware(c echo.Context) error {
	updatedSoftware := new(structs.Software)
	requestID, _ := strconv.Atoi(c.Param("id"))
	if err := c.Bind(updatedSoftware); err != nil {
		return err
	}
	if PutHelper(requestID) {
		sql := "UPDATE software SET name = ?, version = ? WHERE id = ?"
		stmt, err := db.Get().Prepare(sql)
		if err != nil {
			panic(err)
		}
		_, err2 := stmt.Exec(updatedSoftware.Name, updatedSoftware.Version, requestID)
		if err2 != nil {
			panic(err2)
		}

		return c.JSON(http.StatusAccepted, "Updated!")
	}
	sql := "INSERT INTO software(name, version) VALUES(?, ?)"
	stmt, err := db.Get().Prepare(sql)

	if err != nil {
		fmt.Print(err.Error())
	}

	defer stmt.Close()

	result, err2 := stmt.Exec(updatedSoftware.Name, updatedSoftware.Version)

	if err2 != nil {
		panic(err2)
	}

	fmt.Println(result.LastInsertId())

	return c.JSON(http.StatusCreated, updatedSoftware.Name)
}

package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ispeakbinary01/serverTool/db"
	"github.com/ispeakbinary01/serverTool/structs"
	"github.com/labstack/echo/v4"
)

// PostSSH ...
func PostSSH(c echo.Context) error {
	ssh := new(structs.SSH)
	if err := c.Bind(ssh); err != nil {
		return err
	}

	sql := "INSERT INTO ssh(username, password, key) VALUES(?, ?, ?)"
	stmt, err := db.Get().Prepare(sql)

	if err != nil {
		return err
	}

	defer stmt.Close()

	res, err2 := stmt.Exec(ssh.Username, ssh.Password, ssh.Key)

	if err2 != nil {
		return err2
	}

	fmt.Println(res.LastInsertId())

	return c.JSON(http.StatusCreated, "SSH account with username "+ssh.Username+" has been created")
}

// GetSSHs ...
func GetSSHs(c echo.Context) error {
	var id int
	var key int
	var username string
	var password string

	err, _ := db.Get().Query("SELECT id, key FROM ssh")
	for err.Next() {
		err.Scan(&id, &key)
		fmt.Println(strconv.Itoa(id) + ": " + strconv.Itoa(key))
	}

	if err != nil {
		fmt.Println(err)
	}

	response := structs.SSH{ID: id, Username: username, Password: password, Key: key}

	return c.JSON(http.StatusOK, response)

}

// GetSSH ...
func GetSSH(c echo.Context) error {
	requestedID := c.Param("id")
	var id int
	var username string
	var password string
	var key int

	err := db.Get().QueryRow("SELECT id, key FROM ssh WHERE id = ?", requestedID)
	err.Scan(&id, &key)

	response := structs.SSH{ID: id, Username: username, Password: password, Key: key}

	return c.JSON(http.StatusOK, response)
}

// DeleteSSH ...
func DeleteSSH(c echo.Context) error {
	requestID := c.Param("id")
	sql := "DELETE FROM ssh WHERE id = ?"
	stmt, err := db.Get().Prepare(sql)
	if err != nil {
		panic(err)
	}

	_, err2 := stmt.Exec(requestID)
	if err2 != nil {
		panic(err2)
	}

	return c.JSON(http.StatusOK, "Deleted SSH with ID "+requestID)
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

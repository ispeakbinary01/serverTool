package handlers

import (
	"fmt"
	"net/http"

	"github.com/ispeakbinary01/serverTool/db"
	"github.com/ispeakbinary01/serverTool/structs"
	"github.com/labstack/echo/v4"
)

// PostUser ...
func PostUser(c echo.Context) error {
	user := new(structs.User)
	if err := c.Bind(user); err != nil {
		return err
	}

	sql := "INSERT INTO user(username, email, password) VALUES(?, ?, ?)"
	stmt, err := db.Get().Prepare(sql)

	if err != nil {
		fmt.Println(err)
	}

	defer stmt.Close()

	res, err2 := stmt.Exec(user.Username, user.Email, user.Password)

	if err2 != nil {
		panic(err2)
	}

	fmt.Println(res.LastInsertId())

	return c.JSON(http.StatusCreated, "User with username "+user.Username+" was created")

}

// GetUsers ...
func GetUsers(c echo.Context) error {
	var id int
	var username string
	var email string
	var password string

	err, _ := db.Get().Query("SELECT username, email FROM user")
	for err.Next() {
		err.Scan(&username, &email)
		fmt.Println(username + ": " + email)
	}

	if err != nil {
		fmt.Println(err)
	}

	response := structs.User{ID: id, Username: username, Email: email, Password: password}

	return c.JSON(http.StatusOK, response)
}

// GetUser ...
func GetUser(c echo.Context) error {
	requestedID := c.Param("id")
	var id int
	var username string
	var email string
	var password string

	err := db.Get().QueryRow("SELECT username, email FROM user WHERE id = ?", requestedID)
	err.Scan(&username, &email)

	response := structs.User{ID: id, Username: username, Email: email, Password: password}

	return c.JSON(http.StatusOK, response)
}

// DeleteUser ...
func DeleteUser(c echo.Context) error {
	requestID := c.Param("id")
	sql := "DELETE FROM user WHERE id = ?"
	stmt, err := db.Get().Prepare(sql)
	if err != nil {
		panic(err)
	}
	_, err2 := stmt.Exec(requestID)
	if err2 != nil {
		panic(err2)
	}

	return c.JSON(http.StatusOK, "Deleted")
}

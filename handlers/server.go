package handlers

import (
	"fmt"
	"net/http"

	"github.com/ispeakbinary01/serverTool/db"
	"github.com/ispeakbinary01/serverTool/structs"
	"github.com/labstack/echo/v4"
)

// PostServer ...
func PostServer(c echo.Context) error {
	server := new(structs.Server)
	if err := c.Bind(server); err != nil {
		return err
	}

	sql := "INSERT INTO server(ip, os, software, ssh) VALUES(?, ?, ?, ?)"
	stmt, err := db.Get().Prepare(sql)

	if err != nil {
		fmt.Print(err.Error())
	}

	defer stmt.Close()

	res, err2 := stmt.Exec(server.IP, server.Os, server.Software, server.SSH)

	if err2 != nil {
		panic(err2)
	}

	fmt.Println(res.LastInsertId())

	return c.JSON(http.StatusCreated, server.IP)
}

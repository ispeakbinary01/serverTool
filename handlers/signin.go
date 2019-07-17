package handlers

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/ispeakbinary01/serverTool/db"
	"github.com/ispeakbinary01/serverTool/pkg/user"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"time"
)

// checkHash

func checkHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Signin ...
func Signin(c echo.Context) error {
	u := user.User{}
	scanned := user.User{}
	err := c.Bind(&u)
	if err != nil {
		log.Fatal(err)
	}
	stmt := db.Get().QueryRow(userCheck, u.Username)
	if stmt == nil {
		log.Printf("User with username %s was not found. \n", u.Username)
	}
	err2 := stmt.Scan(&scanned.Username, &scanned.Password, &scanned.Position)
	if err2 != nil {
		log.Fatal(err)
	}
	if checkHash(u.Password, scanned.Password) {
		fmt.Printf("Welcome %s \n", scanned.Username)
		token := jwt.New(jwt.SigningMethodHS256)

		claims := token.Claims.(jwt.MapClaims)
		claims["username"] = scanned.Username
		claims["position"] = scanned.Position
		claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, map[string]string{
			"token": t,
		})
	} else {
		fmt.Printf("Wrong password for username %s \n", scanned.Username)
		return echo.ErrUnauthorized
	}
}



const userCheck = `
SELECT username, password, position FROM user WHERE username = ?
`
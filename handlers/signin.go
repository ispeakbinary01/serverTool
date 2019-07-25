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
		log.Printf("%s", err.Error())
	}
	stmt := db.Get().QueryRow(userCheck, u.Email)
	if stmt == nil {
		log.Printf("User with email %s was not found. \n", u.Email)
	}
	err2 := stmt.Scan(&scanned.ID, &scanned.Email, &scanned.Password, &scanned.Position)
	if err2 != nil {
		log.Printf("%s", err)
	}
	if checkHash(u.Password, scanned.Password) {
		fmt.Printf("Welcome %s \n", scanned.Email)

		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["email"] = scanned.Email
		claims["id"] = scanned.ID
		claims["exp"] = time.Now().Add(time.Minute * 60).Unix()

		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			log.Printf("%s", err)
		}

		return c.JSON(http.StatusOK, map[string]string{
			"token": t,
		})
	} else {
		fmt.Printf("Wrong password for email %s \n", scanned.Email)
		return echo.ErrUnauthorized
	}
}


const userCheck = `
SELECT id, email, password, position FROM users WHERE email = ?
`
package user

import (
	"fmt"
	"github.com/ispeakbinary01/serverTool/db"
	"github.com/ispeakbinary01/serverTool/pkg/server"
	"golang.org/x/crypto/bcrypt"
	"log"
)

// hashPassword
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err


}

// CreateUser ...
func (u *User) CreateUser() (int, error) {
	stmt, err := db.Get().Prepare(createUser)
	if err != nil {
		return 0, err
	}
	hash, err := hashPassword(u.Password)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	res, err := stmt.Exec(u.Username, u.Email, hash, u.Position)
	if err != nil {
		return 0, err
	}
	r, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(r), nil
}

// GetAllUsers ...
func GetAllUsers() ([]User, error) {
	us := []User{}
	res, err := db.Get().Query(getUsers)
	for res.Next() {
		u := User{}
		res.Scan(&u.Username, &u.Email, &u.Position)
		us = append(us, u)
		// fmt.Printf("%v+\n")
	}
	if err != nil {
		return nil, err
	}
	return us, nil
}

// GetUserByID ...
func GetUserByID(id string) (*User, error) {
	u := User{}
	res := db.Get().QueryRow(getUser, id)
	if res == nil {
		return nil, nil
	}
	err:= res.Scan(&u.Username, &u.Email, &u.Position)
	if err != nil {
		return nil, err
	}

	return &u, nil
	//if err := db.Get().QueryRow(getUser, id); err != nil {
	//	return nil, err
	//}
	//err.Scan(&u.Username, &u.Email)
	//return &u, nil
}

// DeleteUser ...
func DeleteUser(id string) error {
	stmt, err := db.Get().Prepare(deleteUser)
	if err != nil { 
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {

		return err
	}
	return nil
}

// GetServersByUser
func GetServersByUser(uid interface{}) ([]server.Server,error) {
	var servers []server.Server
	stmt, err := db.Get().Prepare(serversByUser)
	if err != nil {
		return nil, err
	}
	res, err := stmt.Query(uid)
	if err != nil {
		return nil, err
	}
	for res.Next() {
		sr := server.Server{}
		res.Scan(&sr.ID, &sr.IP, &sr.OS)
		servers = append(servers, sr)
	}
	return servers, nil
}

// UpdateUser ...
func (u *User) UpdateUser(id string) (*User, error) {
	//requestID, _ := strconv.Atoi(c.Param("id"))
	//if err := c.Bind(updatedSoftware); err != nil {
	//	return err
	//}
		stmt, err := db.Get().Prepare(updateUser)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		res, err2 := stmt.Exec(&u.Username, &u.Email, &u.Position, id)
		if err2 != nil {
			fmt.Println(err2)
			return nil, err2
		}
		res.LastInsertId()

		return u, nil
}

const serversByUser = `
SELECT server.id, server.ip, server.os FROM server INNER JOIN server_user_rel sur ON server.id = sur.server_id INNER JOIN user ON sur.user_id = user.id WHERE user.id = ?
`

const deleteUser = `
DELETE FROM user WHERE id = ?
`

const getUser = `
SELECT username, email, position FROM user WHERE id = ?
`

const getUsers = `
SELECT username, email, position FROM user
`


const createUser = `
INSERT INTO user (username, email, password, position) VALUES (?, ?, ?, ?)
`

const updateUser = `
UPDATE user SET username = ?, email = ?, position = ? WHERE id = ?
`
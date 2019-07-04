package user

import (
	"fmt"
	"github.com/ispeakbinary01/serverTool/db"
)

// CreateUser ...
func (u *User) CreateUser() (int, error) {
	stmt, err := db.Get().Prepare(createUser)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(u.Username, u.Email, u.Password)
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
		res.Scan(&u.Username, &u.Email)
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
	err:= res.Scan(&u.Username, &u.Email)
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
		res, err2 := stmt.Exec(&u.Username, &u.Email, id)
		if err2 != nil {
			fmt.Println(err2)
			return nil, err2
		}
		res.LastInsertId()

		return u, nil
}

const deleteUser = `
DELETE FROM user WHERE id = ?
`

const getUser = `
SELECT username, email FROM user WHERE id = ?
`

const getUsers = `
SELECT username, email FROM user
`


const createUser = `
INSERT INTO user (username, email, password) VALUES (?, ?, ?)
`

const updateUser = `
UPDATE user SET username = ?, email = ? WHERE id = ?
`
package user

import (

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
func GetAllUsers() ([]User{}, error) {
	us := []User{}
	res, err := db.Get().Query(getUsers)
	for res.Next() {
		u := User{}
		res.Scan(&u.Username, &u.Emai)
		us = append(us, u)
		// fmt.Printf("%v+\n")
	}
	if err != nil {
		return nil, err
	}
	return us
}

// GetUserByID ...
func GetUserByID(id int) (*User, error) {
	u := User{}
	if err := db.Get().QueryRow(getUser, id); err != nil {
		return nil, err
	}
	err.Scan(&u.Username, &u.Email)
	return &u, nil
}

// DeleteUser ...
func DeleteUser(id int) error {
	stmt, err := db.Get().Prepare(deleteUser)
	if err != nil { 
		return err
	}
	if err :s= stmt.Exec(id); err != nil {
		return err
	}
	return nil
} 

const deleteUser = `
DELETE FROM user WHERE id = ?
`

const getUser = `
SELECT name, email FROM user WHERE id = ?
`

const getUsers = `
SELECT username, email FROM user
`

const createUser = `
INSERT INTO user(username, email, password) VALUES(?, ?, ?)
`
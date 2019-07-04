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
func GetUserByID(id string) *User {
	u := User{}
	res := db.Get().QueryRow(getUser, id)
	if res != nil {
		return nil
	}
	err:= res.Scan(u.Username, u.Email)
	if err != nil {
		panic(err)
	}

	return &u
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
INSERT INTO user(username, email, password) VALUES(?, ?, ?)
`
package user

import (
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
		log.Printf("%s", err.Error())
		return 0, err
	}
	hash, err := hashPassword(u.Password)
	if err != nil {
		log.Printf("%s", err.Error())
		return 0, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(u.Email, hash, u.Role)
	if err != nil {
		log.Printf("%s", err.Error())
		return 0, err
	}
	r, err := res.LastInsertId()
	if err != nil {
		log.Printf("%s", err.Error())
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
		res.Scan(&u.Email, &u.Role)
		us = append(us, u)
		// fmt.Printf("%v+\n")
	}
	if err != nil {
		log.Printf("%s", err.Error())
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
	err:= res.Scan(&u.Email, &u.Role)
	if err != nil {
		log.Printf("%s", err.Error())
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
		log.Printf("%s", err.Error())
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		log.Printf("%s", err.Error())
		return err
	}
	return nil
}

// GetServersByUser
func GetServersByUser(uid interface{}) ([]server.Server,error) {
	var servers []server.Server
	stmt, err := db.Get().Prepare(serversByUser)
	if err != nil {
		log.Printf("%s", err.Error())
		return nil, err
	}
	res, err := stmt.Query(uid)
	if err != nil {
		log.Printf("%s", err.Error())
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
			log.Printf("%s", err.Error())
			return nil, err
		}
		res, err2 := stmt.Exec(&u.Email, &u.Role, id)
		if err2 != nil {
			log.Printf("%s", err2.Error())
			return nil, err2
		}
		res.LastInsertId()

		return u, nil
}

const serversByUser = `
SELECT servers.id, servers.ip, servers.os FROM servers INNER JOIN server_user_rel sur ON servers.id = sur.server_id INNER JOIN users ON sur.user_id = users.id WHERE users.id = ?
`

const deleteUser = `
DELETE FROM users WHERE id = ?
`

const getUser = `
SELECT email, role FROM users WHERE id = ?
`

const getUsers = `
SELECT email, role FROM users
`


const createUser = `
INSERT INTO users (email, password, role) VALUES (?, ?, ?)
`

const updateUser = `
UPDATE users SET email = ?, role = ? WHERE id = ?
`
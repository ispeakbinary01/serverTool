package ssh

import (
	"fmt"
	"strconv"

	"github.com/ispeakbinary01/serverTool/db"
)

// CreateSSH ...
func (ssh *SSH) CreateSSH() (int64, error) {
	stmt, err := db.Get().Prepare(createSSH)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(ssh.Username, ssh.Password, ssh.Key)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

// SSHsQuery ...
func SSHsQuery() SSH {
	var id int
	var username string
	var key int
	var password string
	err, _ := db.Get().Prepare(getSSHs)

	for err.Next() {
		err.Scan(&username, &key)
		fmt.Println(username + ": " + strconv.Itoa(key))
	}

	if err != nil {
		fmt.Println(err)
	}
	res := SSH{ID: id, Username: username, Password: password, Key: key}

	return res
}

// SSHQuery ...
func SSHQuery(id int) SSH {
	var username string
	var password string
	var key int
	err := db.Get().QueryRow(getSSH, id)
	err.Scan(&username, &key)

	res := SSH{ID: id, Username: username, Password: password, Key: key}

	return res
}

// DeleteSSHQuery ...
func DeleteSSHQuery(id int) {
	var username string
	var password string
	var key int
	stmt, err := db.Get().Prepare(deleteSSH)
	if err != nil {
		return err
	}
	_, err2 := stmt.Exec(id)
	if err2 != nil {
		return err2
	}

	deleteSSH := SSH{
		ID:       id,
		Username: username,
		Password: password,
		Key:      key,
	}
	return deleteSSH
}

const deleteSSH = `
DELETE FROM ssh WHERE id = ?
`

const getSSH = `
SELECT username, key FROM ssh WHERE id = ?
`

const getSSHs = `
SELECT username, key FROM ssh
`

const createSSH = `
INSERT INTO ssh(username, password, key) VALUES(?, ?, ?)
`

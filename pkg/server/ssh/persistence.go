package ssh

import (
	"fmt"
	"github.com/ispeakbinary01/serverTool/db"
)

// CreateSSH ...
func (ssh *SSH) CreateSSH() (int, error) {
	stmt, err := db.Get().Prepare(createSSH)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(ssh.Username, ssh.Key)
	if err != nil {
		return 0, err
	}
	r, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(r), nil
}

// GetAllSSHs ...
func GetAllSSHs() ([]SSH, error) {
	sshs := []SSH{}
	res, err := db.Get().Query(getSSHs)
	for res.Next() {
		ssh := SSH{}
		res.Scan(&ssh.Username, &ssh.Key)
		sshs = append(sshs, ssh)
		// fmt.Printf("%v+\n")
	}
	if err != nil {
		return nil, err
	}
	return sshs, nil
}

/// GetSShByID...
func GetSShByID(id string) (*SSH, error) {
	ssh := SSH{}
	res := db.Get().QueryRow(getSSH, id)
	if res == nil {
		return nil, nil
	}
	err:= res.Scan(&ssh.Username, &ssh.Key)
	if err != nil {
		return nil, err
	}

	return &ssh, nil
}

// DeleteSSH ...
func DeleteSSH(id string) error {
	stmt, err := db.Get().Prepare(deleteSSH)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

// UpdateSSH ...
func (ssh *SSH) UpdateSSH(id string) (*SSH, error) {
	//requestID, _ := strconv.Atoi(c.Param("id"))
	//if err := c.Bind(updatedSoftware); err != nil {
	//	return err
	//}
	stmt, err := db.Get().Prepare(updateSSH)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	res, err2 := stmt.Exec(ssh.Username, ssh.Key, id)
	if err2 != nil {
		fmt.Println(err2)
		return nil, err2
	}
	res.LastInsertId()

	return ssh, nil
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

const updateSSH = `
UPDATE ssh SET username = ?, key = ? WHERE id = ?
`
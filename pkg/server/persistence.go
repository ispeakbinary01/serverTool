package server

import (
	"fmt"
	"github.com/ispeakbinary01/serverTool/db"
)

// CreateServer ...
func (s *Server) CreateServer() (int, error) {
	stmt, err := db.Get().Prepare(createServer)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(s.IP, s.OS, s.Software, s.SSH)
	if err != nil {
		return 0, err
	}
	r, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(r), nil
}

// GetAllServers ...
func GetAllServers() ([]Server, error) {
	se := []Server{}
	res, err := db.Get().Query(getAllServers)
	if err != nil {
		return nil, err
	}
	for res.Next() {
		s := Server{}
		res.Scan(&s.IP, &s.OS, &s.Software, s.SSH)
		se = append(se, s)
	}
	return se, nil
}

// GetServerByID ...
func GetServerByID(id string) (*Server, error) {
	s := Server{}
	res := db.Get().QueryRow(getServer, id)
	if res == nil {
		return nil, nil
	}
	err := res.Scan(&s.IP, &s.OS, &s.Software, &s.SSH)
	if err != nil {
		return nil, err
	}

	return &s, nil
}

// DeleteServer ...
func DeleteServer(id string) error {
	stmt, err := db.Get().Prepare(deleteServer)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

// UpdateServer ...
func (se *Server) UpdateServer(id string) (*Server, error) {
	stmt, err := db.Get().Prepare(updateServer)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	res, err2 := stmt.Exec(&se.IP, &se.OS, &se.Software, &se.SSH, id)
	if err2 != nil {
		fmt.Println(err2)
		return nil, err2
	}
	res.LastInsertId()

	return se, nil
}

const deleteServer = `
DELETE FROM server WHERE id = ?
`

const getServer = `
SELECT ip, os, software, ssh FROM server WHERE id = ?
`

const getAllServers = `
SELECT id, ip, os, software, ssh FROM server
`


const createServer = `
INSERT INTO server(ip, os, software, ssh) VALUES(?, ?, ?, ?)
`

const updateServer = `
UPDATE server SET ip = ?, os = ?, software = ?, ssh = ? WHERE id = ?
`

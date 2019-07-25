package server

import (
	"github.com/ispeakbinary01/serverTool/db"
	"github.com/ispeakbinary01/serverTool/pkg/server/software"
	"github.com/ispeakbinary01/serverTool/pkg/server/ssh"
	"log"
	"strconv"
)

// CreateServer ...
func (s *Server) CreateServer(uid interface{}) (int, error) {
	stmt, err := db.Get().Prepare(createServer)
	if err != nil {
		log.Printf("%s", err)
		return 0, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(s.IP, s.OS)
	if err != nil {
		log.Printf("%s", err)
		return 0, err
	}
	r, err := res.LastInsertId()
	if err != nil {
		log.Printf("%s", err)
		return 0, err
	}

	stmt2, err := db.Get().Prepare(userServerRel)
	if err != nil {
		log.Printf("%s", err)
		return 0, err
	}
	defer stmt2.Close()
	_, err2 := stmt2.Exec(uid, r)
	if err2 != nil {
		log.Printf("%s", err2)
		return 0, err2
	}
	return int(r), nil
}

// GetAllServers ...
func GetAllServers() ([]Server, error) {
	se := []Server{}
	res, err := db.Get().Query(getAllServers)
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}
	for res.Next() {
		s := Server{}
		err := res.Scan(&s.IP, &s.OS)
		if err != nil {
			log.Printf("%s", err)
			return nil, err
		}
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
	err := res.Scan(&s.IP, &s.OS)
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}

	return &s, nil
}

// DeleteServer ...
func DeleteServer(id string) error {
	stmt, err := db.Get().Prepare(deleteServer)
	if err != nil {
		log.Printf("%s", err)
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		log.Printf("%s", err)
		return err
	}
	return nil
}

// UpdateServer ...
func (s *Server) UpdateServer(id string) (*Server, error) {
	stmt, err := db.Get().Prepare(updateServer)
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}
	res, err2 := stmt.Exec(&s.IP, &s.OS, id)
	if err2 != nil {
		log.Printf("%s", err2)
		return nil, err2
	}
	res.LastInsertId()

	return s, nil
}

// GetServerSSH
func GetServerSSH(serverId string) ([]ssh.SSH, error) {
	s := []ssh.SSH{}
	id, err := strconv.Atoi(serverId)
	if err != nil {
		log.Printf("String conversion failed:  %s\n", err.Error())
		return nil, err
	}
	stmt, err := db.Get().Query(getServerSSH, id)
	if err != nil {
		log.Printf("Query failed:  %s\n", err.Error())
		return nil, err
	}
	for stmt.Next() {
		ssh := ssh.SSH{}
		err := stmt.Scan(&ssh.ID, &ssh.Key)
		if err != nil {
			log.Printf("%s", err)
			return nil, err
		}
		s = append(s, ssh)
	}

	return s, nil
}

// GetServerSoftware
func GetserverSoftware(serverId string) ([]software.Software, error) {
	var sw []software.Software
	id, err := strconv.Atoi(serverId)
	if err != nil {
		log.Printf("String conversion failed %s \n", err.Error())
		return nil, err
	}
	stmt, err := db.Get().Query(getServerSoftware, id)
	if err != nil {
		log.Printf("Query failed: %s \n", err.Error())
		return nil, err
	}
	for stmt.Next() {
		s := software.Software{}
		err := stmt.Scan(&s.Name, &s.Version)
		if err != nil {
			log.Printf("Filling up array failed %s \n", err.Error())
			return nil, err
		}
		sw = append(sw, s)
	}

	return sw, nil
}

const userServerRel = `
INSERT INTO server_user_rel (user_id, server_id) VALUES (?, ?)
`

const deleteServer = `
DELETE FROM servers WHERE id = ?
`

const getServer = `
SELECT ip, os FROM servers WHERE id = ?
`

const getAllServers = `
SELECT id, ip, os FROM servers
`

const createServer = `
INSERT INTO servers(ip, os) VALUES(?, ?)
`

const updateServer = `
UPDATE servers SET ip = ?, os = ? WHERE id = ?
`

const getServerSSH = `
SELECT key FROM ssh WHERE server_id = ?
`

const getServerSoftware = `
SELECT name, version FROM software WHERE server_id = ?
`



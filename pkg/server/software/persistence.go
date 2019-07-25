package software

import (
	"database/sql"
	"github.com/ispeakbinary01/serverTool/db"
	"log"
)

// CreateSoftware ...
func (s *Software) CreateSoftware() (int, error) {
	stmt, err := db.Get().Prepare(createSoftware)
	if err != nil {
		log.Printf("%s", err.Error())
		return 0, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(s.Name, s.Version, s.ServerID)
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
// GetAllSoftware ...
func GetAllSoftware() ([]Software, error) {
	sw := []Software{}
	res, err := db.Get().Query(getAllSoftware)
	for res.Next() {
		s := Software{}
		res.Scan(&s.ID, &s.Name, &s.Version, &s.ServerID)
		sw = append(sw, s)
		// fmt.Printf("%v+\n")
	}
	if err != nil {
		log.Printf("%s", err.Error())
		return nil, err
	}
	return sw, nil
}

// GetSoftwareByID ...
func GetSoftwareByID(id string) (*Software, error) {
	s := Software{}
	res := db.Get().QueryRow(getSoftware, id)
	if res == nil {
		return nil, nil
	}
	err := res.Scan(&s.ID, &s.Name, &s.Version, &s.ServerID)
	if err != nil {
		log.Printf("%s", err.Error())
		return nil, err
	}

	return &s, nil
}

// DeleteSoftware ...
func DeleteSoftware(id string) (sql.Result, error) {
	stmt, err := db.Get().Prepare(deleteSoftware)
	if err != nil {
		log.Printf("%s", err.Error())
		return nil, err
	}
	result, err := stmt.Exec(id)
	if err != nil {
		log.Printf("%s", err.Error())
		return nil, err
	}
	return result, nil
}

// UpdateSoftware ...
func (sw *Software) UpdateSoftware(id string) (*Software, error) {
	//requestID, _ := strconv.Atoi(c.Param("id"))
	//if err := c.Bind(updatedSoftware); err != nil {
	//	return err
	//}
	stmt, err := db.Get().Prepare(updateSoftware)
	if err != nil {
		log.Printf("%s", err.Error())
		return nil, err
	}
	res, err2 := stmt.Exec(&sw.Name, &sw.Version, &sw.ServerID, id)
	if err2 != nil {
		log.Printf("%s", err2.Error())
		return nil, err2
	}
	res.LastInsertId()

	return sw, nil
}

const deleteSoftware = `
DELETE FROM software WHERE id = ?
`

const getSoftware = `
SELECT id, name, version, server_id FROM software WHERE id = ?
`

const getAllSoftware = `
SELECT id, name, version, server_id FROM software
`


const createSoftware = `
INSERT INTO software(name, version, server_id) VALUES(?, ?, ?)
`

const updateSoftware = `
UPDATE software SET name = ?, version = ?, server_id = ? WHERE id = ?
`


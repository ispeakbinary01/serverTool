package software

import (
	"fmt"
	"github.com/ispeakbinary01/serverTool/db"
)

// CreateSoftware ...
func (s *Software) CreateSoftware() (int, error) {
	stmt, err := db.Get().Prepare(createSoftware)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(s.Name, s.Version)
	if err != nil {
		return 0, err
	}
	r, err := res.LastInsertId()
	if err != nil {
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
		res.Scan(&s.Name, &s.Version)
		sw = append(sw, s)
		// fmt.Printf("%v+\n")
	}
	if err != nil {
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
	err := res.Scan(&s.Name, &s.Version)
	if err != nil {
		return nil, err
	}

	return &s, nil
}

// DeleteSoftware ...
func DeleteSoftware(id string) error {
	stmt, err := db.Get().Prepare(deleteSoftware)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

// UpdateSoftware ...
func (sw *Software) UpdateSoftware(id string) (*Software, error) {
	//requestID, _ := strconv.Atoi(c.Param("id"))
	//if err := c.Bind(updatedSoftware); err != nil {
	//	return err
	//}
	stmt, err := db.Get().Prepare(updateSoftware)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	res, err2 := stmt.Exec(&sw.Name, &sw.Version, id)
	if err2 != nil {
		fmt.Println(err2)
		return nil, err2
	}
	res.LastInsertId()

	return sw, nil
}

const deleteSoftware = `
DELETE FROM software WHERE id = ?
`

const getSoftware = `
SELECT name, version FROM software WHERE id = ?
`

const getAllSoftware = `
SELECT id, name, version FROM software
`


const createSoftware = `
INSERT INTO software(name, version) VALUES(?, ?)
`

const updateSoftware = `
UPDATE software SET name = ?, version = ? WHERE id = ?
`

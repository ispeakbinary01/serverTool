package software

import "github.com/ispeakbinary01/serverTool/db"

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
func GetSoftwareByID(id string) *Software {
	s := Software{}
	res := db.Get().QueryRow(getSoftware, id)
	if res != nil {
		return nil
	}
	err:= res.Scan(s.Name, s.Version)
	if err != nil {
		return nil
	}

	return &s
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

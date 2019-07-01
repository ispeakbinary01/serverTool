package db

import "database/sql"

var db *sql.DB

// Get ...
func Get() *sql.DB {
	if db == nil {
		var err error
		db, err = sql.Open("sqlite3", "./serverInventory.db")
		if err != nil {
			panic(err)
		}
	}

	return db
}

// Init ...
func Init() {
	dbc := Get()
	query, err := dbc.Prepare("CREATE TABLE IF NOT EXISTS server(id INTEGER PRIMARY KEY, ip INTEGER, os TEXT, software TEXT, ssh TEXT)")

	if _, err = query.Exec(); err != nil {
		panic(err)
	}
}
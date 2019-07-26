package db

import (
	_ "github.com/mattn/go-sqlite3"
	"database/sql"
	"log"
)

var db *sql.DB

// Get ...
func Get() *sql.DB {
	if db == nil {
		var err error
		db, err = sql.Open("sqlite3", "./serverInventory.db")
		if err != nil {
			log.Printf("%s", err)
		}
	}

	return db
}

// init ...
func init() {
	dbc := Get()
	for i, _ := range tablesCreated {
		query, err := dbc.Prepare(tablesCreated[i])
		if _, err = query.Exec(); err != nil {
			log.Printf("%s", err)
		}
	}
}

var tablesCreated = []string {
"CREATE TABLE IF NOT EXISTS servers (id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE, ip TEXT NOT NULL UNIQUE, os TEXT NOT NULL);",
"CREATE TABLE IF NOT EXISTS server_user_rel (id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE, user_id INTEGER NOT NULL, server_id INTEGER NOT NULL);",
"CREATE TABLE IF NOT EXISTS software (id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE, name TEXT NOT NULL, server_id INTEGER NOT NULL);",
"CREATE TABLE IF NOT EXISTS ssh (id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE, key INTEGER NOT NULL UNIQUE, server_id INTEGER NOT NULL);",
"CREATE TABLE IF NOT EXISTS users (id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE, email TEXT NOT NULL UNIQUE, password TEXT NOT NULL, role TEXT NOT NULL);",
}
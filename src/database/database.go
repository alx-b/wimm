package database

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

const (
	ErrCouldNotConnect     = DatabaseError("Could not connect to specified database file")
	ErrCouldNotCreateTable = DatabaseError("Could not create table")
	//ErrNotFound                   = DatabaseError("Could not get value from database, no such index")
	ErrCouldNotQueryDatabase = DatabaseError("Could not query database")
	//ErrNoValueModified            = DatabaseError("No value was modified")
)

type DatabaseError string

func (e DatabaseError) Error() string {
	return string(e)
}

type Database struct {
	conn *sql.DB
}

// Close database connection, call the sql.DB Close function
func (d *Database) CloseConnection() {
	d.conn.Close()
}

// Create file, tables and connection to database
func CreateDB(filePath string) Database {
	db, err := sql.Open("sqlite", filePath)

	if err != nil {
		panic(fmt.Errorf("%s: %s", ErrCouldNotConnect, err.Error()))
	}

	createIfNotExist(db)

	return Database{db}
}

// Create tables if they do not exist
func createIfNotExist(db *sql.DB) {
	_, err := db.Exec(
		"CREATE TABLE IF NOT EXISTS purchase (id INTEGER PRIMARY KEY, name TEXT, seller TEXT, tag TEXT, cost REAL, date TEXT)",
	)

	if err != nil {
		panic(fmt.Errorf("%s: %s", ErrCouldNotCreateTable, err.Error()))
	}
}

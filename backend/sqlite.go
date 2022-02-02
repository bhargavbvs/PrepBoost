package main

import (
	"database/sql"
)

// SQLite holds the DB connection
type SQLite struct {
	DB *sql.DB
}

func FromSQLite(conn *sql.DB) *SQLite {

	statement, _ := conn.Prepare(`
	CREATE TABLE IF NOT EXISTS 
		questions (
			ID INTEGER PRIMARY KEY, 
			question TEXT, 
			optionA TEXT, 
			optionB TEXT, 
			optionC TEXT, 
			optionD TEXT
		);
	`)

	statement.Exec()
	return &SQLite{
		DB: conn,
	}
}

package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", "./prepboost.db")

	statement, _ := db.Prepare("CREATE TABLE IF NOT EXISTS questions (id INTEGER PRIMARY KEY, question TEXT, optionA TEXT, optionB TEXT, optionC TEXT, optionD TEXT)")
	statement.Exec()

}

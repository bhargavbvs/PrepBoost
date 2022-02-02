package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", "./prepboost.db")
	feed := prepboost.FromSQLite(db)

	r := mux.NewRouter()
	// r.Use(mux.)

	r.HandleFunc("/users/{id}", GetUsers).Methods("GET")
	r.HandleFunc("/users", CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")

	fmt.Println("Watching on port: 3000")
	http.ListenAndServe(":3000", r)

	// statement.Exec()
}

package main

import (
	"log"

	"prepboost.com/web/routes"

	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	router := mux.NewRouter()

	routes.RegisterUserRoutes(router)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe("localhost:9010", router))

	fmt.Println("Watching on port: 9010")
	// http.ListenAndServe(":8080", router)

}

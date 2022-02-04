package main

import (
	controllers "prepboost.com/web/controllers"
	"prepboost.com/web/database"

	"fmt"
	"net/http"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db := database.DatabaseConnection()

	router := controllers.InitializeHandlers(db)
	defer db.Close()

	fmt.Println("Watching on port: 8080")
	http.ListenAndServe(":8080", router)

}

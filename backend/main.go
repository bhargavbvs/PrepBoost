package main

import (
	controllers "prepboost.com/web/controllers/users"
	"prepboost.com/web/database"

	"fmt"
	"net/http"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db := database.DatabaseConnection()

	router := controllers.InitializeUserHandlers(db)
	defer db.Close()

	fmt.Println("Watching on port: 3000")
	http.ListenAndServe(":3000", router)

}

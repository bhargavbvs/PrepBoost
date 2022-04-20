package main

import (
	models "prepboost.com/web/models"
	"prepboost.com/web/routes"

	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/cors"
)

func main() {

	router := mux.NewRouter()

	// Establishing Database Connection
	models.Init()

	corsWrapper := cors.New(cors.Options{
		AllowedMethods:   []string{"GET", "POST"},
		AllowedHeaders:   []string{"Content-Type", "Origin", "Accept", "*"},
		AllowCredentials: true,
		AllowedOrigins:   []string{"http://localhost:3000"},
	})

	routes.RegisterUserRoutes(router)
	routes.RegisterQuestionsRoutes(router)
	routes.RegisterBookmarkRoutes(router)
	routes.RegisterLeaderboardRoutes(router)
	http.Handle("/", router)
	fmt.Println("Starting server....")
	// log.Fatal(http.ListenAndServe("localhost:9010", router))
	http.ListenAndServe(":9010", corsWrapper.Handler(router))
	fmt.Println("Watching on port: 9010")
	// http.ListenAndServe(":8080", router)

}

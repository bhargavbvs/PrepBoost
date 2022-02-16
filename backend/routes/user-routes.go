package routes

import (
	"github.com/gorilla/mux"
	controllers "prepboost.com/web/controllers/users"
)

var RegisterUserRoutes = func(router *mux.Router) {
	router.HandleFunc("/users/login/", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/users/{userId}", controllers.GetUser).Methods("GET")
	router.HandleFunc("/users/{userId}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{userId}", controllers.DeleteUser).Methods("DELETE")
}

// func InitializeHandlers(db *gorm.DB) *mux.Router {

// 	router := mux.NewRouter()

// 	//Handlers for the Users API
// 	router.HandleFunc("/users/login", func(w http.ResponseWriter, r *http.Request) {
// 		controllers.CreateUser(w, r, db)
// 	}).Methods("POST")
// 	router.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
// 		controllers.GetUser(w, r, db)
// 	}).Methods("GET")
// 	router.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
// 		controllers.UpdateUser(w, r, db)
// 	}).Methods("PUT")
// 	router.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
// 		controllers.DeleteUser(w, r, db)
// 	}).Methods("DELETE")

// 	return router
// }

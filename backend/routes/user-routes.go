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

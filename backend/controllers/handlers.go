package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	controllers_que "prepboost.com/web/controllers/questions"
	controllers "prepboost.com/web/controllers/users"
)

func InitializeHandlers(db *gorm.DB) *mux.Router {

	router := mux.NewRouter()

	//Handlers for the Users API
	router.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		controllers.CreateUser(w, r, db)
	}).Methods("POST")
	router.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		controllers.GetUser(w, r, db)
	}).Methods("GET")
	router.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		controllers.UpdateUser(w, r, db)
	}).Methods("PUT")
	router.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		controllers.DeleteUser(w, r, db)
	}).Methods("DELETE")

	//Handlers for the Questions API
	router.HandleFunc("/questions/{year}", func(w http.ResponseWriter, r *http.Request) {
		controllers_que.GetYearwiseQuestions(w, r, db)
	}).Methods("GET")

	return router
}

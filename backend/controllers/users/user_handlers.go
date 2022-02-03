package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func InitializeUserHandlers(db *gorm.DB) *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		CreateUser(w, r, db)
	}).Methods("POST")
	router.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		GetUser(w, r, db)
	}).Methods("GET")
	router.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		UpdateUser(w, r, db)
	}).Methods("PUT")
	router.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		DeleteUser(w, r, db)
	}).Methods("DELETE")
	return router
}

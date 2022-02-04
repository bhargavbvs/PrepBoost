package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"prepboost.com/web/database"
)

func CreateUser(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	var p database.UserJson

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	parsingerr := HandleUserCreationUpdationErrors(p, db, -1)
	if parsingerr != "" {
		HandleErrorResponse(w, parsingerr)
		return
	}
	user := database.User{Username: p.Username, Mobile: p.Mobile, Paid: p.Paid, Search_left: p.Search_left, Session_id: p.Session_id}
	result := db.Create(&user)
	if result.Error != nil {
		fmt.Println(result.Error.Error())
	}
	body, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)

}

func UpdateUser(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		HandleErrorResponse(w, "Id must be a valid integer")
		return
	}
	var user database.User
	db.First(&user, "id = ?", id)
	var p database.UserJson
	err = json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		fmt.Println(err)
		return
	}
	parsingerr := HandleUserCreationUpdationErrors(p, db, id)
	if parsingerr != "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(parsingerr))
		return
	}
	user.Update(p.Username, p.Mobile, p.Paid, p.Search_left, p.Session_id)
	db.Save(&user)
	body, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)

}

func DeleteUser(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	vars := mux.Vars(r)
	id := vars["id"]
	var user database.User
	db.First(&user, "id = ?", id)
	body, err := json.Marshal(user)
	db.Delete(&user)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)

}

func GetUser(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	vars := mux.Vars(r)
	id := vars["id"]
	var user database.User
	db.First(&user, "id = ?", id)
	body, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

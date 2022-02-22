package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	models "prepboost.com/web/models"
	"prepboost.com/web/utils"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	CreateUser := &models.User{}
	utils.ParseJsonBody(r, CreateUser)

	u := CreateUser.CreateUser()
	res, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var updateUser = &models.User{}
	utils.ParseJsonBody(r, updateUser)

	vars := mux.Vars(r)
	userId := vars["userId"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	userDetails, db := models.GetUserById(ID)
	body, err := json.Marshal(userDetails)
	if err != nil {
		fmt.Println("error while updating user details")
		fmt.Println(err)
		fmt.Println(body)
		return
	}

	if updateUser.Username != "" {
		userDetails.Username = updateUser.Username
	}
	if updateUser.Email != "" {
		userDetails.Email = updateUser.Email
	}
	if updateUser.Mobile != "" {
		userDetails.Mobile = updateUser.Mobile
	}

	db.Save(&userDetails)
	res, _ := json.Marshal(userDetails)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	userDetails := models.DeleteUser(ID)
	body, err := json.Marshal(userDetails)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	userDetails, _ := models.GetUserById(ID)
	body, err := json.Marshal(userDetails)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

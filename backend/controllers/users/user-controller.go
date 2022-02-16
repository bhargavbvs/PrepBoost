package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	// config "prepboost.com/web/config"
	models "prepboost.com/web/models"
	"prepboost.com/web/utils"
)

var db *gorm.DB

// func init() {
// 	db = config.DatabaseConnection()
// 	db.AutoMigrate(&Book{})
// }

func CreateUser(w http.ResponseWriter, r *http.Request) {

	CreateUser := &models.User{}
	utils.ParseJsonBody(r, CreateUser)

	u := CreateUser.CreateUser()
	res, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
		return
	}
	// parsingerr := HandleUserCreationUpdationErrors(p, db, -1)
	// if parsingerr != "" {
	// 	HandleErrorResponse(w, parsingerr)
	// 	return
	// }

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		HandleErrorResponse(w, "Id must be a valid integer")
		return
	}
	var user models.User
	db.First(&user, "id = ?", id)
	var p models.UserJson
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

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var user models.User
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

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	userDetails := models.GetUserById(ID)
	body, err := json.Marshal(userDetails)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

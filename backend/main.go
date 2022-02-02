package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID          uint       `gorm:"primaryKey;autoIncrement"`
	Username    string     `gorm:"not null"`
	Mobile      string     `gorm:"not null;unique"`
	Paid        int        `gorm:"not null;default:0"`
	Search_left int        `gorm:"not null;default:100"`
	Session_id  string     `gorm:"not null;unique"`
	Created_at  *time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP" faker:"-"`
	Updated_at  *time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP" faker:"-"`
}

type UserJson struct {
	Username    string
	Mobile      string
	Paid        int
	Search_left int
	Session_id  string
}

func (r *User) update(
	Username string,
	Mobile string,
	Paid int,
	Search_left int,
	Session_id string) int {
	r.Username = Username
	r.Mobile = Mobile
	r.Paid = Paid
	r.Search_left = Search_left
	r.Session_id = Session_id
	return 1
}

func HandleUserCreationUpdationErrors(user UserJson, db *gorm.DB, id int) string {
	var dbUser User
	err := ""
	result := db.Where("session_id = ? and id != ?", user.Session_id, id).First(&dbUser)
	if result.RowsAffected != 0 {
		err = "User with session id Already exists " + user.Session_id
		return err
	}
	result = db.Where("mobile = ? and id != ?", user.Mobile, id).First(&dbUser)
	if result.RowsAffected != 0 {
		err = "User with Mobile Number Already exists " + user.Mobile
		return err
	}
	return ""
}

func HandleErrorResponse(w http.ResponseWriter, error_msg string) {
	error_resp := `{"error": "` + error_msg + `"}`
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(error_resp))
	return
}

func CreateUser(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	var p UserJson

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
	user := User{Username: p.Username, Mobile: p.Mobile, Paid: p.Paid, Search_left: p.Search_left, Session_id: p.Session_id}
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
	id, _ := strconv.Atoi(vars["id"])
	var user User
	db.First(&user, "id = ?", id)
	var p UserJson
	err := json.NewDecoder(r.Body).Decode(&p)
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
	user.update(p.Username, p.Mobile, p.Paid, p.Search_left, p.Session_id)
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
	var user User
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
	var user User
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

func DBMigrate(db *gorm.DB) {

	db.AutoMigrate(&User{})
}

func main() {

	db, _ := gorm.Open("sqlite3", "./prepboost.db")
	defer db.Close()
	DBMigrate(db)
	r := mux.NewRouter()

	r.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		CreateUser(w, r, db)
	}).Methods("POST")
	r.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		GetUser(w, r, db)
	}).Methods("GET")
	r.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		UpdateUser(w, r, db)
	}).Methods("PUT")
	r.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		DeleteUser(w, r, db)
	}).Methods("DELETE")
	fmt.Println("Watching on port: 3000")
	http.ListenAndServe(":3000", r)

}

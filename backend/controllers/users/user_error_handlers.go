package controllers

import (
	"net/http"

	"github.com/jinzhu/gorm"
	models "prepboost.com/web/models"
)

func HandleUserCreationUpdationErrors(user models.User, db *gorm.DB, id int) string {
	var dbUser models.User
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

	result = db.Where(" id = ? ", id).First(&dbUser)
	if result.RowsAffected == 0 && id > 0 {
		err = "User with id not exists" + string(dbUser.ID)
		return err
	}
	return ""
}

func HandleErrorResponse(w http.ResponseWriter, error_msg string) {
	error_resp := `{"error": "` + error_msg + `"}`
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(error_resp))
}

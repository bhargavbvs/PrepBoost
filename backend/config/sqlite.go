package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

// func DBMigrate(db *gorm.DB) {
// 	db.AutoMigrate(&User{}, &Questions{}, &Exam{})
// }

func checkErr(err error, sucess_message string) {
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(sucess_message)
	}
}

func DatabaseConnection() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./prepboost.db")
	checkErr(err, "Database Created")

	// DBMigrate(db)

	return db

}

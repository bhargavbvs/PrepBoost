package models

import (
	"time"

	"github.com/jinzhu/gorm"
	config "prepboost.com/web/config"
)

var db *gorm.DB

// User table declarations

type User struct {
	ID          uint       `gorm:"primaryKey;autoIncrement"`
	Username    string     `gorm:"not null"`
	Mobile      string     `gorm:"not null;unique"`
	Email       string     `gorm:"null"`
	Paid        int        `gorm:"not null;default:0"`
	Search_left int        `gorm:"not null;default:100"`
	Session_id  string     `gorm:"not null;unique"`
	Created_at  *time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP" faker:"-"`
	Updated_at  *time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP" faker:"-"`
}

type Exam struct {
	ID         uint       `gorm:"primaryKey;autoIncrement"`
	Exam       string     `gorm:"not null;unique"`
	Type       string     `gorm:"not null"`
	Created_at *time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP" faker:"-"`
	Updated_at *time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP" faker:"-"`
}

type Questions struct {
	ID          uint `gorm:"primaryKey;autoIncrement"`
	ExamID      uint
	Exam        Exam       `gorm:"foreignkey : ExamID"`
	Year        uint       `gorm:"not null"`
	Question    string     `gorm:"not null"`
	Answer      string     `gorm:"not null"`
	Option1     string     `gorm:"not null"`
	Option2     string     `gorm:"not null"`
	Option3     string     `gorm:"not null"`
	Option4     string     `gorm:"not null"`
	Explanation string     `gorm:"null"`
	Created_at  *time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP" faker:"-"`
	Updated_at  *time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP" faker:"-"`
}

func init() {
	db = config.DatabaseConnection()
	// db = config.GetDB()
	db.AutoMigrate(&User{}, &Exam{}, &Questions{})
}

func (u *User) CreateUser() *User {
	db.NewRecord(u)
	db.Create(&u)
	return u
}

func GetUserById(Id int64) (*User, *gorm.DB) {
	var user User
	db.Where("ID=?", Id).Find(&user)
	return &user, db
}

func DeleteUser(ID int64) User {
	var user User
	db.Where("ID=?", ID).Delete(user)
	return user
}

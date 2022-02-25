package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	ID          uint       `gorm:"primaryKey;autoIncrement"`
	Username    string     `gorm:"not null"`
	Mobile      string     `gorm:"not null;unique"`
	Email       string     `gorm:"null"`
	Password    string     `gorm:"not null"`
	Paid        int        `gorm:"not null;default:0"`
	Search_left int        `gorm:"not null;default:100"`
	Session_id  string     `gorm:"not null;unique"`
	Created_at  *time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP" faker:"-"`
	Updated_at  *time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP" faker:"-"`
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

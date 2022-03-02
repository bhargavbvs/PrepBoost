package models

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
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

type LoginUser struct {
	Username string `gorm:"not null"`
	Password string `gorm:"not null"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func (u *User) CreateUser() *User {
	db.NewRecord(u)
	db.Create(&u)
	return u
}

func UserPasswordCheck(loginUser LoginUser) string {
	var user User
	db.Where("username=?", loginUser.Username).Find(&user)
	return user.Password
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

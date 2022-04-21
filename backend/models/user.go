package models

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

type User struct {
	ID          uint      `gorm:"primaryKey;autoIncrement"`
	Username    string    `gorm:"not null"`
	Mobile      string    `gorm:"not null;unique"`
	Email       string    `gorm:"null"`
	Password    string    `gorm:"not null"`
	Paid        int       `gorm:"not null;default:0"`
	Search_left int       `gorm:"not null;default:100"`
	Session_id  string    `gorm:"not null;unique"`
	Created_at  time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP" faker:"-"`
	Updated_at  time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP" faker:"-"`
	Token       string    `gorm:"null"`
}

type LoginUser struct {
	Username string `gorm:"not null"`
	Password string `gorm:"not null"`
	Token    string `gorm:"null"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func (u *User) CreateUser() *User {
	DB.NewRecord(u)
	DB.Create(&u)
	return u
}

func UserPasswordCheck(loginUser LoginUser) string {
	var user User
	DB.Where("username=?", loginUser.Username).Find(&user)
	return user.Password
}

func FindUserId(loginUser LoginUser) uint {
	var user User
	DB.Where("username=?", loginUser.Username).Find(&user)
	return user.ID
}

func GetUserById(Id int64) (*User, *gorm.DB) {
	var user User
	fmt.Println("testing request received", Id)
	DB.Where("ID=?", Id).Find(&user)
	return &user, DB
}

func GetUserFromUserId(Id uint) (*User, *gorm.DB) {
	var user User
	DB.Where("ID=?", Id).Find(&user)
	return &user, DB
}

func DeleteUser(ID int64) User {
	var user User
	DB.Where("ID=?", ID).Delete(user)
	return user
}

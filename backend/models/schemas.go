package models

import (
	"time"

	"github.com/jinzhu/gorm"
	config "prepboost.com/web/config"
)

var db *gorm.DB

func init() {
	db = config.DatabaseConnection()

	db.AutoMigrate(&User{}, &Exams{}, &Questions{})
}

type Exams struct {
	ID         uint       `gorm:"primaryKey;autoIncrement"`
	Exam       string     `gorm:"not null;unique"`
	Type       string     `gorm:"not null"`
	Created_at *time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP" faker:"-"`
	Updated_at *time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP" faker:"-"`
}

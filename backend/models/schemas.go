package models

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	config "prepboost.com/web/config"
)

var DB *gorm.DB

func Init() {
	fmt.Print("database connection ......")
	DB = config.DatabaseConnection()

	DB.AutoMigrate(&User{}, &Exams{}, &Questions{})
}

type Exams struct {
	ID         uint       `gorm:"primaryKey;autoIncrement"`
	Exam       string     `gorm:"not null;unique"`
	Type       string     `gorm:"not null"`
	Created_at *time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP" faker:"-"`
	Updated_at *time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP" faker:"-"`
}

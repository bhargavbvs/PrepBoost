package models

import (
	"time"
)

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

func GetQuestionsByYear(Year int64) []Questions {
	var Questions []Questions
	db.Where("Year=?", Year).Find(&Questions)
	return Questions
}

package models

import (
	"gorm.io/gorm"
)

type Speaker struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Company     string
	Phone       string
	Email       string `gorm:"uniqueIndex"`
	TgNickname  string
	Description string `gorm:"type:text"`
	ImageURL    string
	Events      []Event `gorm:"many2many:event_speakers;"`
}

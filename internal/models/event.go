package models

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Name        string    `gorm:"not null"`
	Description string    `gorm:"type:text"`
	StartTime   time.Time `gorm:"not null"`
	EndTime     time.Time `gorm:"not null"`
	Tags        []Tag     `gorm:"many2many:event_tags;"`
	Speakers    []Speaker `gorm:"many2many:event_speakers;"`
}

type Tag struct {
	gorm.Model
	Name   string  `gorm:"uniqueIndex;not null"`
	Events []Event `gorm:"many2many:event_tags;"`
}

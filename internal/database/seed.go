package database

import (
	"time"

	"github.com/hokamsingh/go-backend-template/internal/models"
	"gorm.io/gorm"
)

func SeedDatabase(db *gorm.DB) error {
	// Проверяем, есть ли уже данные в базе
	var count int64
	db.Model(&models.Tag{}).Count(&count)
	if count > 0 {
		return nil // Данные уже существуют, пропускаем заполнение
	}

	// Create tags
	tags := []models.Tag{
		{Name: "Технологии"},
		{Name: "Бизнес"},
		{Name: "Наука"},
		{Name: "Искусство"},
	}
	for _, tag := range tags {
		if err := db.Create(&tag).Error; err != nil {
			return err
		}
	}

	// Create speakers
	speakers := []models.Speaker{
		{
			Name:        "Иван Петров",
			Company:     "Tech Solutions",
			Phone:       "+7 999 123 45 67",
			Email:       "ivan@example.com",
			TgNickname:  "@ipetrov",
			Description: "Эксперт по искусственному интеллекту",
			ImageURL:    "https://example.com/speakers/ivan.jpg",
		},
		{
			Name:        "Анна Сидорова",
			Company:     "Data Insights",
			Phone:       "+7 999 765 43 21",
			Email:       "anna@example.com",
			TgNickname:  "@asidorova",
			Description: "Специалист по анализу данных",
			ImageURL:    "https://example.com/speakers/anna.jpg",
		},
	}
	for _, speaker := range speakers {
		if err := db.Create(&speaker).Error; err != nil {
			return err
		}
	}

	// Create events
	events := []models.Event{
		{
			Name:        "Конференция по AI",
			Description: "Обсуждение последних трендов в ИИ",
			StartTime:   time.Now().Add(24 * time.Hour),
			EndTime:     time.Now().Add(48 * time.Hour),
		},
		{
			Name:        "Воркшоп по Data Science",
			Description: "Практический семинар по анализу данных",
			StartTime:   time.Now().Add(72 * time.Hour),
			EndTime:     time.Now().Add(96 * time.Hour),
		},
	}
	for _, event := range events {
		if err := db.Create(&event).Error; err != nil {
			return err
		}
	}

	// Связываем события с тегами и спикерами
	if err := db.Model(&events[0]).Association("Tags").Append(&tags[0], &tags[2]); err != nil {
		return err
	}
	if err := db.Model(&events[0]).Association("Speakers").Append(&speakers[0]); err != nil {
		return err
	}

	if err := db.Model(&events[1]).Association("Tags").Append(&tags[0], &tags[1]); err != nil {
		return err
	}
	if err := db.Model(&events[1]).Association("Speakers").Append(&speakers[1]); err != nil {
		return err
	}

	return nil
}

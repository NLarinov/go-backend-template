package repository

import (
	"context"

	"github.com/hokamsingh/go-backend-template/internal/models"
	"gorm.io/gorm"
)

type SpeakerRepository struct {
	db *gorm.DB
}

func NewSpeakerRepository(db *gorm.DB) *SpeakerRepository {
	return &SpeakerRepository{db: db}
}

func (r *SpeakerRepository) Create(ctx context.Context, speaker *models.Speaker) error {
	return r.db.WithContext(ctx).Create(speaker).Error
}

func (r *SpeakerRepository) GetByID(ctx context.Context, id uint) (*models.Speaker, error) {
	var speaker models.Speaker
	err := r.db.WithContext(ctx).Preload("Events").First(&speaker, id).Error
	if err != nil {
		return nil, err
	}
	return &speaker, nil
}

func (r *SpeakerRepository) GetAll(ctx context.Context) ([]*models.Speaker, error) {
	var speakers []*models.Speaker
	err := r.db.WithContext(ctx).Preload("Events").Find(&speakers).Error
	return speakers, err
}

func (r *SpeakerRepository) Update(ctx context.Context, speaker *models.Speaker) error {
	return r.db.WithContext(ctx).Save(speaker).Error
}

func (r *SpeakerRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.Speaker{}, id).Error
}

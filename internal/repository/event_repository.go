package repository

import (
	"context"

	"github.com/hokamsingh/go-backend-template/internal/models"
	"gorm.io/gorm"
)

type EventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) *EventRepository {
	return &EventRepository{db: db}
}

func (r *EventRepository) Create(ctx context.Context, event *models.Event) error {
	return r.db.WithContext(ctx).Create(event).Error
}

func (r *EventRepository) GetByID(ctx context.Context, id uint) (*models.Event, error) {
	var event models.Event
	err := r.db.WithContext(ctx).Preload("Tags").Preload("Speakers").First(&event, id).Error
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (r *EventRepository) GetAll(ctx context.Context) ([]*models.Event, error) {
	var events []*models.Event
	err := r.db.WithContext(ctx).Preload("Tags").Preload("Speakers").Find(&events).Error
	return events, err
}

func (r *EventRepository) Update(ctx context.Context, event *models.Event) error {
	return r.db.WithContext(ctx).Save(event).Error
}

func (r *EventRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.Event{}, id).Error
}

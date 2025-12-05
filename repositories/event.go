package repositories

import (
	"context"

	"github.com/gucchunchun/ticket-booking-system-v1/models"
	"gorm.io/gorm"
)

type EventRepository struct {
	db *gorm.DB
}

func (r *EventRepository) GetMany(ctx context.Context) ([]*models.Event, error) {
	events := []*models.Event{}

	res := r.db.Model(&models.Event{}).Find(&events)

	if res.Error != nil {
		return nil, res.Error
	}

	return events, nil
}

func (r *EventRepository) GetOne(ctx context.Context, id uint) (*models.Event, error) {
	event := &models.Event{}

	res := r.db.Model(&models.Event{}).First(event, id)

	if res.Error != nil {
		return nil, res.Error
	}

	return event, nil
}

func (r *EventRepository) CreateOne(ctx context.Context, event *models.Event) error {
	// Implementation goes here
	return nil
}

func NewEventRepository(db *gorm.DB) models.EventRepository {
	return &EventRepository{db: db}
}

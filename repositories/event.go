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

	res := r.db.Model(&models.Event{}).Order("updated_at desc").Find(&events)

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

func (r *EventRepository) CreateOne(ctx context.Context, event *models.Event) (*models.Event, error) {
	res := r.db.Model(&models.Event{}).Create(event)

	if res.Error != nil {
		return nil, res.Error
	}

	return event, nil
}

func (r *EventRepository) UpdateOne(ctx context.Context, id uint, updateData map[string]interface{}) (*models.Event, error) {
	res := r.db.Model(&models.Event{}).Where("id = ?", id).Updates(updateData)

	if res.Error != nil {
		return nil, res.Error
	}

	return r.GetOne(ctx, id)
}

func (r *EventRepository) DeleteOne(ctx context.Context, id uint) error {
	res := r.db.Delete(&models.Event{}, id)

	if res.Error != nil {
		return res.Error
	}
	return nil
}

func NewEventRepository(db *gorm.DB) models.EventRepository {
	return &EventRepository{db: db}
}

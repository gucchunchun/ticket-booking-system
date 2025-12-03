package repositories

import (
	"context"
	"time"

	"github.com/gucchunchun/ticket-booking-system-v1/models"
)

type EventRepository struct {
	db any
}

func (r *EventRepository) GetMany(ctx context.Context) ([]*models.Event, error) {
	events := []*models.Event{}

	events = append(events, &models.Event{
		ID:        1,
		Name:      "Concert A",
		Location:  "Venue A",
		Date:      time.Now(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	return events, nil
}

func (r *EventRepository) GetOne(ctx context.Context, id uint) (*models.Event, error) {
	// Implementation goes here
	return nil, nil
}

func (r *EventRepository) CreateOne(ctx context.Context, event *models.Event) error {
	// Implementation goes here
	return nil
}

func NewEventRepository(db any) models.EventRepository {
	return &EventRepository{db: db}
}

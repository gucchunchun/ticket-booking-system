package models

import (
	"context"
	"time"
)

type Event struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	Location  string    `json:"location"`
	Date      time.Time `json:"date"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type EventRepository interface {
	GetMany(ctx context.Context) ([]*Event, error)
	GetOne(ctx context.Context, id uint) (*Event, error)
	CreateOne(ctx context.Context, event *Event) error
	// Update(ctx context.Context, event *Event) error
	// Delete(ctx context.Context, id uint) error
}

package models

import (
	"context"
	"time"
)

type Ticket struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	EventID   uint      `json:"event_id"`
	Event     Event     `gorm:"foreignKey:EventID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"event"`
	Entered   bool      `gorm:"default:false" json:"entered"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TicketRepository interface {
	GetMany(ctx context.Context) ([]*Ticket, error)
	GetOne(ctx context.Context, id uint) (*Ticket, error)
	CreateOne(ctx context.Context, ticket *Ticket) (*Ticket, error)
	UpdateOne(ctx context.Context, id uint, updateData map[string]interface{}) (*Ticket, error)
	// DeleteOne(ctx context.Context, id uint) error
}

type ValidTicket struct {
	TicketID uint `json:"ticket_id"`
}

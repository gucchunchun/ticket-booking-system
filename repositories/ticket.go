package repositories

import (
	"context"

	"github.com/gucchunchun/ticket-booking-system-v1/models"
	"gorm.io/gorm"
)

func NewTicketRepository(db *gorm.DB) models.TicketRepository {
	return &TicketRepository{db: db}
}

type TicketRepository struct {
	db *gorm.DB
}

func (r *TicketRepository) GetMany(ctx context.Context) ([]*models.Ticket, error) {
	tickets := []*models.Ticket{}

	res := r.db.Model(&models.Ticket{}).Preload("Event").Order("updated_at desc").Find(&tickets)

	if res.Error != nil {
		return nil, res.Error
	}

	return tickets, nil
}

func (r *TicketRepository) GetOne(ctx context.Context, id uint) (*models.Ticket, error) {
	ticket := &models.Ticket{}

	res := r.db.Model(&models.Ticket{}).Preload("Event").First(&ticket, id)

	if res.Error != nil {
		return nil, res.Error
	}

	return ticket, nil
}
func (r *TicketRepository) CreateOne(ctx context.Context, ticket *models.Ticket) (*models.Ticket, error) {
	res := r.db.Model(&models.Ticket{}).Create(&ticket)

	if res.Error != nil {
		return nil, res.Error
	}

	return ticket, nil
}

func (r *TicketRepository) UpdateOne(ctx context.Context, id uint, updateData map[string]interface{}) (*models.Ticket, error) {
	res := r.db.Model(&models.Ticket{}).Where("id = ?", id).Updates(updateData)

	if res.Error != nil {
		return nil, res.Error
	}
	return r.GetOne(ctx, id)
}

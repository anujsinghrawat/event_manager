package repositories

import (
	"context"

	"github.com/anujsinghrawat/event-manager/models"
	"gorm.io/gorm"
)

type TicketRepository struct {
	db *gorm.DB
}

func (r *TicketRepository) GetAll(ctx context.Context, userId uint) ([]*models.Ticket, error) {
	var tickets []*models.Ticket
	res := r.db.Model(&models.Ticket{}).Where("user_id = ?", userId).Preload("Event").Order("updated_at desc").Find(&tickets)
	if res.Error != nil {
		return nil, res.Error
	}
	return tickets, nil
}

func (r *TicketRepository) GetByID(ctx context.Context, ticketId uint, userId uint) (*models.Ticket, error) {
	ticket := &models.Ticket{}
	res := r.db.Model(ticket).Preload("Event").Where("id = ?", ticketId).Where("user_id = ?", userId).First(ticket)
	if res.Error != nil {
		return nil, res.Error
	}
	return ticket, nil
}

func (r *TicketRepository) Create(ctx context.Context, ticket *models.Ticket, userId uint) (*models.Ticket, error) {
	ticket.UserID = userId
	if err := r.db.Create(ticket).Error; err != nil {
		return nil, err
	}
	return r.GetByID(ctx, ticket.ID, userId)
}

func (r *TicketRepository) UpdateTicket(ctx context.Context, ticketId uint, updateData map[string]interface{}, userId uint) (*models.Ticket, error) {
	// ticket := &models.Ticket{}
	// if err := r.db.First(&ticket, ticketId).Error; err != nil {
	// 	return nil, err
	// }
	// if err := r.db.Model(ticket).Updates(updateData).Error; err != nil {
	// 	return nil, err
	// }
	// return ticket, nil

	ticket := &models.Ticket{}
	updateRes := r.db.Model(ticket).Where("id = ?", ticketId).Updates(updateData)
	if updateRes.Error != nil {
		return nil, updateRes.Error
	}
	return r.GetByID(ctx, ticketId, userId)
}

func (r *TicketRepository) DeleteTicket(ctx context.Context, ticketId uint, userId uint) error {
	if err := r.db.Delete(&models.Ticket{}, ticketId).Error; err != nil {
		return err
	}
	return nil
}

func NewTicketRepository(db *gorm.DB) models.TicketRepository {
	return &TicketRepository{
		db: db,
	}
}

package models

import (
	"context"

	"gorm.io/gorm"
)

type Ticket struct {
	gorm.Model
	EventID uint `json:"event_id"`
	UserID uint `json:"user_id" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Event   Event `json:"event" gorm:"foreignKey:EventID;constraints:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Entered bool `json:"entered" gorm:"default:false"`
}


type TicketRepository interface {
	GetAll(ctx context.Context, userId uint) ([]*Ticket, error)
	GetByID(ctx context.Context, ticketId uint, userId uint) (*Ticket, error)
	Create(ctx context.Context, ticket *Ticket, userId uint) (*Ticket, error)
	UpdateTicket(ctx context.Context, ticketId uint, updateData map[string]interface{}, userId uint) (*Ticket, error)
	DeleteTicket(ctx context.Context, ticketId uint, userId uint) error
}

type ValidateTicket struct {
	TicketID uint `json:"ticket_id"`
	OwnerId uint `json:"owner_id"`
}
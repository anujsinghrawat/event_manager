package models

import (
	"context"
	"time"

	"gorm.io/gorm"
)

// Using gorm.Model embeds ID, CreatedAt, UpdatedAt, and DeletedAt fields.
type Event struct {
	gorm.Model
	Name                  string    `json:"name"`
	Location              string    `json:"location"`
	Date                  time.Time `json:"date"`
	TotalTicketsPurchased int64     `json:"total_tickets_purchased" gorm:"-"`
	TotalTicketsEntered   int64     `json:"total_tickets_entered" gorm:"-"`
}

type EventRepository interface {
	GetAll(ctx context.Context) ([]*Event, error)
	GetByID(ctx context.Context, eventId uint) (*Event, error)
	Create(ctx context.Context, event *Event) (*Event, error)
	UpdateEvent(ctx context.Context, eventId uint, updateData map[string]interface{}) (*Event, error)
	DeleteEvent(ctx context.Context, eventId uint) error
}

func (e *Event) AfterFind(db *gorm.DB) (err error) {
	baseQuery := db.Model(&Ticket{}).Where(&Ticket{EventID: e.ID})

	if res := baseQuery.Count(&e.TotalTicketsPurchased); res.Error != nil {
		return res.Error
	}

	if res := baseQuery.Where("entered = ?", true).Count(&e.TotalTicketsEntered); res.Error != nil {
		return res.Error
	}
	return nil

}

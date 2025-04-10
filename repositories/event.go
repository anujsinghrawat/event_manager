package repositories

import (
	"context"

	"github.com/anujsinghrawat/event-manager/models"
	"gorm.io/gorm"
)

type EventRepository struct {
	db *gorm.DB
}

func (r *EventRepository) GetAll(ctx context.Context) ([]*models.Event, error) {
	var events []*models.Event
	// Order before calling Find.
	if err := r.db.Order("updated_at desc").Find(&events).Error; err != nil {
		return nil, err
	}
	return events, nil
}

func (r *EventRepository) GetByID(ctx context.Context, eventId uint) (*models.Event, error) {
	event := &models.Event{}
	if err := r.db.First(&event, eventId).Error; err != nil {
		return nil, err
	}
	return event, nil
}

func (r *EventRepository) Create(ctx context.Context, event *models.Event) (*models.Event, error) {
	if err := r.db.Create(event).Error; err != nil {
		return nil, err
	}
	return event, nil
}

func (r *EventRepository) UpdateEvent(ctx context.Context, eventId uint, updateData map[string]interface{}) (*models.Event, error) {
	event := &models.Event{}
	if err := r.db.First(&event, eventId).Error; err != nil {
		return nil, err
	}
	if err := r.db.Model(event).Updates(updateData).Error; err != nil {
		return nil, err
	}
	return event, nil
}

func (r *EventRepository) DeleteEvent(ctx context.Context, eventId uint) error {
	if err := r.db.Delete(&models.Event{}, eventId).Error; err != nil {
		return err
	}
	return nil
}

func NewEventRepository(db *gorm.DB) *EventRepository {
	return &EventRepository{db: db}
}

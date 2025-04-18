package repositories

import (
	"context"

	"github.com/anujsinghrawat/event-manager/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func (r *UserRepository) GetAll(ctx context.Context) ([]*models.User, error) {
	var users []*models.User
	res := r.db.Model(&models.User{}).Order("created_at desc").Find(&users)
	if res.Error != nil {
		return nil, res.Error
	}
	return users, nil
}

func NewUserRepository(db *gorm.DB) models.UserRepository {
	return &UserRepository{
		db: db,
	}
}

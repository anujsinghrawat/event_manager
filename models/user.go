package models

import (
	"context"

	"gorm.io/gorm"
)

type UserRole string

const (
	Manager  UserRole = "manager"
	Attendee UserRole = "attendee"
)

type User struct {
	gorm.Model
	Email    string   `json:"email" gorm:"text;not null"`
	Role     UserRole `json:"role" gorm:"text;default:attendee"`
	Password string   `json:"-"` // Exclude from JSON response
}


type UserRepository interface {
	GetAll(ctx context.Context) ([]*User, error)
}

// this was just fun
// func (u *User) AfterCreate(db *gorm.DB)(err error){
// 	if u.ID == 1 {
// 		db.Model(u).Update("role", Manager)
// 	}
// 	return
// }

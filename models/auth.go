package models

import (
	"context"
	"net/mail"

	"golang.org/x/crypto/bcrypt"
)

type RegisterUser struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	Role     UserRole `json:"role" validate:"required"`
}

type AuthCredential struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type AuthRepository interface {
	RegisterUser(ctx context.Context, registerData *RegisterUser) (*User, error)
	GetUser(ctx context.Context, query interface{}, args ...interface{}) (*User, error)
}

type AuthService interface {
	Login(ctx context.Context, loginData *AuthCredential) (string, *User, error)
	Register(ctx context.Context, registerData *RegisterUser) (string, *User, error)
}

func MatchesHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

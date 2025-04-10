package services

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/anujsinghrawat/event-manager/models"
	"github.com/anujsinghrawat/event-manager/utils"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	repository models.AuthRepository
}

func (s *AuthService) Login(ctx context.Context, loginData *models.AuthCredential) (string, *models.User, error) {
	user, err := s.repository.GetUser(ctx, "email = ?", loginData.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", nil, fmt.Errorf("Invalid credentials")
		}
		return "", nil, err
	}

	if !models.MatchesHash(loginData.Password, user.Password) {
		return "", nil, fmt.Errorf("Invalid Password")
	}

	claims := jwt.MapClaims{
		"id":   user.ID,
		"role": user.Role,
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
	}

	token, err := utils.GenerateJWT(claims, jwt.SigningMethodHS256, os.Getenv("JWT_SECRET"))

	if err != nil {
		return "", nil, err
	}
	return token, user, nil

}

func (s *AuthService) Register(c context.Context, registerData *models.AuthCredential) (string, *models.User, error) {
	if !models.IsValidEmail(registerData.Email) {
		return "", nil, fmt.Errorf("Invalid email, please enter a valid email")
	}

	if _, err := s.repository.GetUser(c, "email = ?", registerData.Email); err == nil {
		return "", nil, fmt.Errorf("The email is already registered")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerData.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", nil, err
	}
	registerData.Password = string(hashedPassword)

	user, err := s.repository.RegisterUser(c, registerData)
	if err != nil {
		return "", nil, err
	}

	claims := jwt.MapClaims{
		"id":   user.ID,
		"role": user.Role,
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
	}

	token, err := utils.GenerateJWT(claims, jwt.SigningMethodHS256, os.Getenv("JWT_SECRET"))
	if err != nil {
		return "", nil, err
	}
	return token, user, nil
}

func NewAuthService(repository models.AuthRepository) models.AuthService {
	return &AuthService{repository: repository}
}

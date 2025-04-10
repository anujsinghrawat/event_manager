package handlers

import (
	"context"
	"fmt"
	"time"

	"github.com/anujsinghrawat/event-manager/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

type AuthHandler struct {
	service models.AuthService
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	creds := &models.AuthCredential{}
	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	if err := c.BodyParser(&creds); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  -1,
			"message": err.Error(),
			"data":    nil,
		})
	}

	if err := validate.Struct(creds); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  -1,
			"message": err.Error(),
			"data":    nil,
		})
	}

	token, user, err := h.service.Login(context, creds)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"status":  -2,
			"message": err.Error(),
			"data":    nil,
		})
	}
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": 0,
		"data": &fiber.Map{
			"token": token,
			"user":  user,
		},
		"message": "Login successful",
	})
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	creds := &models.AuthCredential{}
	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	if err := c.BodyParser(&creds); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  -1,
			"message": err.Error(),
			"data":    nil,
		})
	}

	if err := validate.Struct(creds); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  -2,
			"message": fmt.Errorf("Please provide a valid email and password").Error(),
			"data":    nil,
		})
	}

	token, user, err := h.service.Register(context, creds)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"status":  -3,
			"message": err.Error(),
			"data":    nil,
		})
	}
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": 0,
		"data": &fiber.Map{
			"token": token,
			"user":  user,
		},
		"message": "Register successful",
	})
}

func NewAuthHandler(router fiber.Router, service models.AuthService) {
	handler := &AuthHandler{service: service}

	router.Post("/login", handler.Login)
	router.Post("/register", handler.Register)
}

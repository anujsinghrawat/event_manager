package handlers

import (
	"context"
	"time"

	"github.com/anujsinghrawat/event-manager/models"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	repository models.UserRepository
}

func (h *UserHandler) GetAll(c *fiber.Ctx) error {
	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	users, err := h.repository.GetAll(context)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(&fiber.Map{
			"status":  -1,
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  0,
		"data":    users,
		"message": "Users fetched successfully",
	})
}

func NewUserHandler(router fiber.Router, repository models.UserRepository) {
	handler := &UserHandler{repository: repository}

	router.Get("/", handler.GetAll)

}

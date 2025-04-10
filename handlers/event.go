package handlers

import (
	"context"
	"strconv"
	"time"

	"github.com/anujsinghrawat/event-manager/models"
	"github.com/gofiber/fiber/v2"
)

type EventHandler struct {
	repository models.EventRepository
}

func (h *EventHandler) GetAll(c *fiber.Ctx) error {
	context, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	events, err := h.repository.GetAll(context)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(&fiber.Map{
			"status":  -1,
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  0,
		"data":    events,
		"message": "Events List Fetched Successfully",
	})
}

func (h *EventHandler) GetByID(c *fiber.Ctx) error {
	eventId, _ := strconv.Atoi(c.Params("eventId"))
	context, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	event, err := h.repository.GetByID(context, uint(eventId))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  -1,
			"message": "Error Fetching Event By ID " + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  0,
		"data":    event,
		"message": "Event Fetched Successfully",
	})
}

func (h *EventHandler) Create(c *fiber.Ctx) error {
	event := &models.Event{}
	context, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := c.BodyParser(event); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status":  -1,
			"message": "Error Parsing Event " + err.Error(),
			"data":    nil,
		})
	}

	event, err := h.repository.Create(context, event)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  -1,
			"message": "Error Creating Event " + err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"status":  0,
		"data":    event,
		"message": "Event Created Successfully",
	})
}

func (h *EventHandler) UpdateEvent(c *fiber.Ctx) error {
	eventId, _ := strconv.Atoi(c.Params("eventId"))
	updateData :=  make(map[string]interface{})
	context, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status":  -1,
			"message": "Error Parsing Event " + err.Error(),
			"data":    nil,
		})
	}

	event, err := h.repository.UpdateEvent(context, uint(eventId), updateData)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  -1,
			"message": "Error Updating Event " + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  0,
		"data":    event,
		"message": "Event Updated Successfully",
	})
}

func (h *EventHandler) DeleteEvent(c *fiber.Ctx) error {
	eventId, _ := strconv.Atoi(c.Params("eventId"))

	context, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	err := h.repository.DeleteEvent(context, uint(eventId))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  -1,
			"message": "Error Deleting Event :" + err.Error(),
		})
	}
	return c.Status(fiber.StatusNoContent).JSON(&fiber.Map{
		"status":  0,
		"message": "Event Deleted Successfully",
	})
}

func NewEventHandler(router fiber.Router, repository models.EventRepository) {
	handler := &EventHandler{repository: repository}

	router.Get("/", handler.GetAll)
	router.Get("/:eventId", handler.GetByID)
	router.Post("/", handler.Create)
	router.Put("/:eventId", handler.UpdateEvent)
	router.Delete("/:eventId", handler.DeleteEvent)
}

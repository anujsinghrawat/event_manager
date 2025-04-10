package handlers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/anujsinghrawat/event-manager/models"
	"github.com/gofiber/fiber/v2"
	"github.com/skip2/go-qrcode"
)

type TicketHandler struct {
	repository models.TicketRepository
}

func (h *TicketHandler) GetAll(c *fiber.Ctx) error {
	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()
	userId := c.Locals("userId").(uint)
	tickets, err := h.repository.GetAll(context, userId)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(&fiber.Map{
			"status":  -1,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  0,
		"data":    tickets,
		"message": "Tickets fetched successfully",
	})
}

func (h *TicketHandler) GetByID(c *fiber.Ctx) error {
	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	// ticketId := c.Params("ticket_id")
	ticketId, _ := strconv.Atoi(c.Params("ticket_id"))
	userId := c.Locals("userId").(uint)
	ticket, err := h.repository.GetByID(context, uint(ticketId), userId)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(&fiber.Map{
			"status":  -1,
			"message": err.Error(),
		})
	}
	var QRCode []byte
	QRCode, err = qrcode.Encode(
		fmt.Sprintf("ticketId: %d, ownerId: %d", ticket.ID, userId),
		qrcode.Medium,
		250,
	)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  -2,
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  0,
		"data":    &fiber.Map{
			"ticket": ticket,
			"qrCode": QRCode,
		},
		"message": "Ticket fetched successfully",
	})
}

func (h *TicketHandler) Create(c *fiber.Ctx) error {
	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	ticket := &models.Ticket{}
	if err := c.BodyParser(ticket); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status":  -1,
			"message": err.Error(),
		})
	}
	userId := c.Locals("userId").(uint)

	ticket, err := h.repository.Create(context, ticket, userId)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(&fiber.Map{
			"status":  -1,
			"message": err.Error(),
		})
	}
	return c.Status((fiber.StatusCreated)).JSON(&fiber.Map{
		"status":  0,
		"data":    ticket,
		"message": "Ticket created successfully",
	})
}

func (h *TicketHandler) Validate(c *fiber.Ctx) error {
	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	validateBody := &models.ValidateTicket{}

	if err := c.BodyParser(validateBody); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status":  -1,
			"message": err.Error(),
		})
	}
	validateData := make(map[string]interface{})
	validateData["entered"] = true
	userId := c.Locals("userId").(uint)
	ticket, err := h.repository.UpdateTicket(context, validateBody.TicketID, validateData, userId)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(&fiber.Map{
			"status":  -1,
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  0,
		"data":    ticket,
		"message": "Welcome to the show",
	})
}

func NewTicketHandler(router fiber.Router, repository models.TicketRepository) {
	handler := &TicketHandler{repository: repository}

	router.Get("/", handler.GetAll)
	router.Get("/:ticket_id", handler.GetByID)
	router.Post("/", handler.Create)
	router.Post("/validate", handler.Validate)
	// router.Put("/:ticket_id", handler.Update)
}

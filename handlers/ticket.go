package handlers

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gucchunchun/ticket-booking-system-v1/models"
)

type TicketHandler struct {
	repository models.TicketRepository
}

func (h *TicketHandler) GetMany(ctx *fiber.Ctx) error {
	ctx2, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	tickets, err := h.repository.GetMany(ctx2)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Tickets fetched successfully",
		"data":    tickets,
	})
}
func (h *TicketHandler) GetOne(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("ticketId")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "fail",
			"message": "Invalid ticket ID",
		})
	}

	ctx2, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	ticket, err := h.repository.GetOne(ctx2, uint(id))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Ticket fetched successfully",
		"data":    ticket,
	})
}
func (h *TicketHandler) CreateOne(ctx *fiber.Ctx) error {
	ticket := &models.Ticket{}
	if err := ctx.BodyParser(ticket); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "fail",
			"message": "Invalid request body",
		})
	}

	ctx2, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	createdTicket, err := h.repository.CreateOne(ctx2, ticket)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "Ticket created successfully",
		"data":    createdTicket,
	})
}
func (h *TicketHandler) UpdateOne(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("ticketId")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "fail",
			"message": "Invalid ticket ID",
		})
	}

	updateData := make(map[string]interface{})
	if err := ctx.BodyParser(&updateData); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "fail",
			"message": "Invalid request body",
		})
	}

	ctx2, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	updatedTicket, err := h.repository.UpdateOne(ctx2, uint(id), updateData)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Ticket updated successfully",
		"data":    updatedTicket,
	})
}

func NewTicketHandler(router fiber.Router, repository models.TicketRepository) *TicketHandler {
	handler := &TicketHandler{repository: repository}

	router.Get("/", handler.GetMany)
	router.Get("/:ticketId", handler.GetOne)
	router.Post("/", handler.CreateOne)
	router.Put("/:ticketId", handler.UpdateOne)

	return handler
}

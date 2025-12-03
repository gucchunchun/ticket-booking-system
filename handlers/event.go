package handlers

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gucchunchun/ticket-booking-system-v1/models"
)

type EventHandler struct {
	repository models.EventRepository
}

func (h *EventHandler) GetMany(ctx *fiber.Ctx) error {
	ctx2, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	events, err := h.repository.GetMany(ctx2)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Events fetched successfully",
		"data":    events,
	})
}
func (h *EventHandler) GetOne(c *fiber.Ctx) error {
	return nil
}
func (h *EventHandler) CreateOne(c *fiber.Ctx) error {
	return nil
}

func NewEventHandler(router fiber.Router, repository models.EventRepository) {
	handler := &EventHandler{
		repository: repository,
	}

	router.Get("/", handler.GetMany)
	router.Get("/:eventId", handler.GetOne)
	router.Post("/", handler.CreateOne)
}

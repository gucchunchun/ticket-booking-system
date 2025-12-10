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
func (h *EventHandler) GetOne(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("eventId")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "fail",
			"message": "Invalid event ID",
		})
	}

	ctx2, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	event, err := h.repository.GetOne(ctx2, uint(id))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Event fetched successfully",
		"data":    event,
	})
}
func (h *EventHandler) CreateOne(ctx *fiber.Ctx) error {
	event := &models.Event{}

	ctx2, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	if err := ctx.BodyParser(event); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "fail",
			"message": "Invalid request body",
		})
	}

	createdEvent, err := h.repository.CreateOne(ctx2, event)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "fail",
			"message": err.Error(),
			"data":    createdEvent,
		})
	}
	return nil
}
func (h *EventHandler) UpdateOne(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("eventId")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "fail",
			"message": "Invalid event ID",
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

	updatedEvent, err := h.repository.UpdateOne(ctx2, uint(id), updateData)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "fail",
			"message": err.Error(),
			"data":    updatedEvent,
		})
	}
	return nil
}

func (h *EventHandler) DeleteOne(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("eventId")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "fail",
			"message": "Invalid event ID",
		})
	}

	ctx2, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	err = h.repository.DeleteOne(ctx2, uint(id))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}
	return nil
}

func NewEventHandler(router fiber.Router, repository models.EventRepository) {
	handler := &EventHandler{
		repository: repository,
	}

	router.Get("/", handler.GetMany)
	router.Get("/:eventId", handler.GetOne)
	router.Post("/", handler.CreateOne)
	router.Put("/:eventId", handler.UpdateOne)
	router.Delete("/:eventId", handler.DeleteOne)
}

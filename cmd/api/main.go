package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gucchunchun/ticket-booking-system-v1/handlers"
	"github.com/gucchunchun/ticket-booking-system-v1/repositories"
)

func main() {
	// Application entry point
	app := fiber.New(fiber.Config{
		AppName:      "Ticket Booking System",
		ServerHeader: "fiber",
	})

	// Initialize database connection (db)
	var db any // Replace with actual DB connection

	eventRepository := repositories.NewEventRepository(db)

	server := app.Group("/api/v1")

	handlers.NewEventHandler(server.Group("/event"), eventRepository)

	app.Listen(":3000")
}

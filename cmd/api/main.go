package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gucchunchun/ticket-booking-system-v1/config"
	"github.com/gucchunchun/ticket-booking-system-v1/db"
	"github.com/gucchunchun/ticket-booking-system-v1/handlers"
	"github.com/gucchunchun/ticket-booking-system-v1/repositories"
)

func main() {
	envConfig := config.NewEnvConfig()

	db := db.Init(envConfig, db.DBMigrator)

	// Application entry point
	app := fiber.New(fiber.Config{
		AppName:      "Ticket Booking System",
		ServerHeader: "fiber",
	})

	eventRepository := repositories.NewEventRepository(db)

	server := app.Group("/api/v1")

	handlers.NewEventHandler(server.Group("/event"), eventRepository)

	app.Listen(fmt.Sprintf(":%s", envConfig.ServerPort))
}

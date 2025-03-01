package main

import (
	"log"

	"task-service/internal/handlers"
	"task-service/pkg/db"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database := db.Connect()
	defer db.Close(database)

	app := fiber.New()

	handlers.SetupRoutes(app, database)

	log.Fatal(app.Listen(":8080"))
}

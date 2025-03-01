package handlers

import (
	"task-service/internal/repository"
	"task-service/internal/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// SetupRoutes registers all routes
func SetupRoutes(app *fiber.App, db *gorm.DB) {
	// Initialize repository and service
	taskRepo := repository.NewTaskRepository(db)
	taskService := services.NewTaskService(taskRepo)

	// Initialize handler
	taskHandler := NewTaskHandler(taskService)

	// Define routes
	api := app.Group("/api") // API prefix

	api.Post("/tasks", taskHandler.CreateTask)
	api.Get("/tasks", taskHandler.GetTasks)
	api.Get("/tasks/:id", taskHandler.GetTask)
	api.Put("/tasks/:id", taskHandler.UpdateTask)
	api.Delete("/tasks/:id", taskHandler.DeleteTask)
}

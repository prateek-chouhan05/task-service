package handlers

import (
	"strconv"
	"task-service/internal/models"
	"task-service/internal/services"

	"github.com/gofiber/fiber/v2"
)

type TaskHandler struct {
	service *services.TaskService
}

func NewTaskHandler(service *services.TaskService) *TaskHandler {
	return &TaskHandler{service}
}


func (h *TaskHandler) CreateTask(c *fiber.Ctx) error {
	var task models.Task


	if err := c.BodyParser(&task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Invalid JSON format",
			"message": err.Error(),
		})
	}

	if task.Title == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Validation error",
			"message": "Title is required",
		})
	}
	validStatuses := map[string]bool{
		"Pending":    true,
		"InProgress": true,
		"Completed":  true,
	}
	if _, valid := validStatuses[string(task.Status)]; !valid {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Validation error",
			"message": "Invalid status. Allowed values: 'Pending', 'InProgress', 'Completed'",
		})
	}

	// 🔹 Call service to create task
	if err := h.service.CreateTask(&task); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Database error",
			"message": "Failed to create task",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Task created successfully",
		"task":    task,
	})
}



func (h *TaskHandler) GetTasks(c *fiber.Ctx) error {
	// Parse query parameters
	limit, err := strconv.Atoi(c.Query("limit", "10"))
	if err != nil || limit <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Validation error",
			"message": "Limit must be a positive integer",
		})
	}

	offset, err := strconv.Atoi(c.Query("offset", "0"))
	if err != nil || offset < 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Validation error",
			"message": "Offset must be a non-negative integer",
		})
	}

	status := c.Query("status", "")

	// Fetch tasks and total count
	tasks, totalCount, err := h.service.GetTasks(limit, offset, status)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Database error",
			"message": "Failed to fetch tasks",
		})
	}

	// Calculate pagination values
	currentPage := (offset / limit) + 1
	totalPages := (totalCount + limit - 1) / limit // Equivalent to ceil(totalCount / limit)
	hasNext := offset+limit < totalCount

	// Return paginated response
	return c.JSON(fiber.Map{
		"tasks":       tasks,
		"total_pages": totalPages,
		"current_page": currentPage,
		"has_next":    hasNext,
	})
}
func (h *TaskHandler) GetTask(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	task, err := h.service.GetTaskByID(uint(id))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Task not found"})
	}
	return c.JSON(task)
}

func (h *TaskHandler) UpdateTask(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	task := new(models.Task)
	if err := c.BodyParser(task); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	task.ID = uint(id)
	if err := h.service.UpdateTask(task); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update task"})
	}
	return c.JSON(task)
}

func (h *TaskHandler) DeleteTask(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	if err := h.service.DeleteTask(uint(id)); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete task"})
	}
	return c.JSON(fiber.Map{"message": "Task deleted successfully"})
}

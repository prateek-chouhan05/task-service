package main

import (
	"log"
	"task-service/internal/models"
	"task-service/pkg/db"
)

func main() {
	// Connect to the database
	database := db.Connect()
	defer db.Close(database)

	// Run migrations
	err := database.AutoMigrate(&models.Task{})
	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	log.Println("Migrations applied successfully!")
}

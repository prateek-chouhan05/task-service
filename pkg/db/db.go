package db

import (
	"fmt"
	"log"
	"task-service/internal/models"
	"task-service/pkg/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connect initializes the database connection
func Connect() *gorm.DB {
	cfg := config.LoadConfig()

	// Database connection string
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort,
	)

	// Connect to PostgreSQL
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// 🚀 Run AutoMigrate for models (This runs migrations!)
	err = db.AutoMigrate(&models.Task{})
	if err != nil {
		log.Fatal("Failed to run migrations:", err)
	}

	log.Println("Database connected!")
	return db
}

// Close terminates the database connection
func Close(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Println("Error getting database connection:", err)
		return
	}
	sqlDB.Close()
}

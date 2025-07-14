package main

import (
	"log"

	"api-go/config"
	"api-go/models"
	"api-go/routes"
)

func main() {
	// Initialize database
	config.InitDB()

	// Auto migrate the schema
	err := config.DB.AutoMigrate(&models.Product{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	log.Println("Database migration completed")

	// Setup routes
	r := routes.SetupRoutes()

	// Start server
	log.Println("Server starting on :8000")
	err = r.Run(":8000")
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
} 
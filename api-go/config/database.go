package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, using default values")
	}

	var dsn string
	if os.Getenv("DATABASE_URL") != "" {
		dsn = os.Getenv("DATABASE_URL")
	} else {
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_DATABASE"),
		)
	}

	log.Println("Connecting to database:", dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = db
	log.Println("Database connected successfully")

	// Instrument GORM with OpenTelemetry
	if err := DB.Use(otelgorm.NewPlugin()); err != nil {
		log.Println("Failed to register otelgorm plugin:", err)
	} else {
		log.Println("otelgorm plugin registered for OpenTelemetry tracing")
	}
} 
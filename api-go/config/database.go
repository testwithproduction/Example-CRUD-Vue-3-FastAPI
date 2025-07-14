package config

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	slog.Debug("Loading .env file")
	err := godotenv.Load()
	if err != nil {
		slog.Warn("Error loading .env file, using default values", "err", err)
	}

	var dsn string
	if os.Getenv("DATABASE_URL") != "" {
		dsn = os.Getenv("DATABASE_URL")
		slog.Debug("Using DATABASE_URL from environment", "dsn", dsn)
	} else {
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_DATABASE"),
		)
		slog.Debug("Built DSN from environment variables", "dsn", dsn)
	}

	slog.Info("Connecting to database", "dsn", dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		slog.Error("Failed to connect to database", "err", err)
		os.Exit(1)
	}

	DB = db
	slog.Info("Database connected successfully")

	// Instrument GORM with OpenTelemetry
	slog.Debug("Registering otelgorm plugin for OpenTelemetry tracing")
	if err := DB.Use(otelgorm.NewPlugin()); err != nil {
		slog.Warn("Failed to register otelgorm plugin", "err", err)
	} else {
		slog.Info("otelgorm plugin registered for OpenTelemetry tracing")
	}
} 
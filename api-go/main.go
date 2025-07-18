package main

import (
	"log/slog"

	"api-go/config"
	"api-go/models"
	"api-go/routes"
	"context"
	"os"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
)

func main() {
	// --- OpenTelemetry + OTLP setup ---
	ctx := context.Background()
	otlpEndpoint := os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT")
	slog.Debug("Loaded OTEL_EXPORTER_OTLP_ENDPOINT", "otlpEndpoint", otlpEndpoint)
	if otlpEndpoint == "" {
		otlpEndpoint = "localhost:4317" // default for Jaeger/OTEL collector gRPC
		slog.Debug("OTLP endpoint not set, using default", "default", otlpEndpoint)
	}
	// Initialize slog logger with JSON output and debug level
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	slog.SetDefault(logger)

	slog.Debug("Initializing OTLP exporter", "endpoint", otlpEndpoint)
	exp, err := otlptracegrpc.New(ctx,
		otlptracegrpc.WithEndpoint(otlpEndpoint),
		otlptracegrpc.WithInsecure(),
	)
	if err != nil {
		slog.Error("failed to create OTLP exporter", "err", err)
		os.Exit(1)
	}
	resource := resource.NewWithAttributes(
		"",
		attribute.String("service.name", "api-go"),
		attribute.String("environment", os.Getenv("ENVIRONMENT")),
	)
	tracerProvider := trace.NewTracerProvider(
		trace.WithBatcher(exp),
		trace.WithResource(resource),
	)
	otel.SetTracerProvider(tracerProvider)
	defer func() {
		if err = tracerProvider.Shutdown(ctx); err != nil {
			slog.Error("failed to shutdown tracer provider", "err", err)
			os.Exit(1)
		}
	}()
	// --- End OpenTelemetry setup ---
	// Initialize database
	slog.Debug("Initializing database connection")
	config.InitDB()

	// Auto migrate the schema
	slog.Debug("Running database migrations")
	err = config.DB.AutoMigrate(&models.Product{})
	if err != nil {
		slog.Error("Failed to migrate database", "err", err)
		os.Exit(1)
	}
	slog.Info("Database migration completed")

	// Setup routes
	slog.Debug("Setting up HTTP routes")
	r := routes.SetupRoutes()

	// Start server
	slog.Info("Server starting", "addr", ":8000")
	err = r.Run(":8000")
	if err != nil {
		slog.Error("Failed to start server", "err", err)
		os.Exit(1)
	}
} 
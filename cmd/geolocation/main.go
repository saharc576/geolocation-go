package main

import (
	"geolocation/pkg/api"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	// Initialize endpoint
	endpoint := api.NewGetEnrichGeolocation()

	// Register the endpoint with Fiber
	app.Add(endpoint.Method(), endpoint.Path(), endpoint.Handler())

	// Start the server on port 8080
	if err := app.Listen(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

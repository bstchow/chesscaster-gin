package main

import (
	"chesscaster-gin/models"
	"chesscaster-gin/routes"
)

// Entrypoint for app.
func main() {
	// Load the routes
	r := routes.SetupRouter()

	// Initialize database
	models.SetupDatabase()

	// Start the HTTP API
	r.Run()
}

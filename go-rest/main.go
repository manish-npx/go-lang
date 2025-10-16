package main

import (
	"log"
	"net/http"

	"github.com/manish-npx/go-lang/go-rest/routes"
)

func main() {
	// Register all routes defined in routes.go
	routes.RegisterRoutes()

	// Start the server on port 8080
	log.Println("✅ Server running on http://localhost:8080")

	// ListenAndServe keeps the server running.
	// If it fails, log.Fatal will print the error and stop the program.
	log.Fatal(http.ListenAndServe(":8080", nil))
}

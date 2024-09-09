package main

import (
	"log"
	"net/http"

	"automated-seo-analyzer/backend/internal/api"
	"automated-seo-analyzer/backend/internal/database"
)

func main() {
	log.Println("Starting SEO Analyzer server")

	// Initialize the database
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()
	log.Println("Database connection established")

	mux := http.NewServeMux()

	// Set up routes
	api.SetupRoutes(mux)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

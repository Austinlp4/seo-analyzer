package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Austinlp4/seo-analyzer/backend/internal/api"
)

func main() {
	mux := http.NewServeMux()

	// Register API routes
	api.RegisterHandlers(mux)

	// Serve static files and handle client-side routing
	mux.HandleFunc("/", api.HandleStaticFiles)

	fmt.Println("Server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

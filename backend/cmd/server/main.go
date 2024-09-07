package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/yourusername/seo-analyzer/backend/internal/api"
	"github.com/yourusername/seo-analyzer/backend/internal/auth"
	"github.com/yourusername/seo-analyzer/backend/internal/database"
)

func main() {
	// Initialize the database
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	router := gin.Default()

	// Apply middlewares
	router.Use(api.SecurityHeadersMiddleware())
	router.Use(api.RateLimitMiddleware(redisClient))

	// Set up routes
	api.SetupRoutes(router, db)
	auth.SetupRoutes(router, db)

	// Serve frontend
	router.Static("/", "./frontend/build")

	log.Fatal(router.Run(":8080"))
}

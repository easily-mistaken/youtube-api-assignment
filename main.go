package main

import (
	"os"
	"time"

	"github.com/easily-mistaken/youtube-api-assignment/pkg/api"
	"github.com/easily-mistaken/youtube-api-assignment/pkg/database"
	"github.com/easily-mistaken/youtube-api-assignment/pkg/youtube"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database
	db := database.Connect()
	
	// Create YouTube client with multiple API keys
	ytClient := youtube.NewClient([]string{
		os.Getenv("YOUTUBE_API_KEY_1"),
		os.Getenv("YOUTUBE_API_KEY_2"),
	})

	// Start background fetcher with configurable interval
	fetchInterval := 10 * time.Second
	query := os.Getenv("YOUTUBE_SEARCH_QUERY")
	
	go youtube.StartPeriodicFetch(db, ytClient, query, fetchInterval)

	// Setup HTTP server
	router := gin.Default()
	handler := api.NewHandler(db)
	
	// Register routes
	api.RegisterRoutes(router, handler)

	// Start server
	router.Run(":8080")
}
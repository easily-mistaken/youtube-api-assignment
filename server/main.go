package main

import (
	"os"
	"time"

	"github.com/easily-mistaken/youtube-api-assignment/pkg/api"
	"github.com/easily-mistaken/youtube-api-assignment/pkg/database"
	"github.com/easily-mistaken/youtube-api-assignment/pkg/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db := database.Connect()
	
	ytService := services.NewYoutubeService(db)

	fetchInterval := 10 * time.Second
	query := os.Getenv("YOUTUBE_SEARCH_QUERY")
	
	// Start periodic fetch in a goroutine
	go func() {
		for {
			_, err := ytService.FetchVideos(query)
			if err != nil {
				// Log error but continue
				println("Error fetching videos:", err.Error())
			}
			time.Sleep(fetchInterval)
		}
	}()

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	handler := api.NewHandler(db)
	api.RegisterRoutes(router, handler)

	router.Run(":8080")
}
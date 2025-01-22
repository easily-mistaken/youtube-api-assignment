package main

import (
	"os"
	"time"

	"github.com/easily-mistaken/youtube-api-assignment/pkg/api"
	"github.com/easily-mistaken/youtube-api-assignment/pkg/database"
	"github.com/easily-mistaken/youtube-api-assignment/pkg/youtube"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db := database.Connect()
	
	ytClient := youtube.NewClient([]string{
		os.Getenv("YOUTUBE_API_KEY_1"),
		os.Getenv("YOUTUBE_API_KEY_2"),
	})

	fetchInterval := 10 * time.Second
	query := os.Getenv("YOUTUBE_SEARCH_QUERY")
	
	go youtube.StartPeriodicFetch(db, ytClient, query, fetchInterval)

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
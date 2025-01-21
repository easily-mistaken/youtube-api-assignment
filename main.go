package main

import (
    "log"
    "os"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/easily-mistaken/youtube-api-assignment/internal/youtube"
    "github.com/easily-mistaken/youtube-api-assignment/internal/worker"
)

func main() {
    // Initialize Gin
    r := gin.Default()

    // Basic health check
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })

    // Initialize YouTube client
    apiKeys := []string{os.Getenv("YOUTUBE_API_KEY")}
    youtubeClient := youtube.NewClient(apiKeys)

    // Initialize fetcher
    fetcher := worker.NewVideoFetcher(youtubeClient, 10*time.Second, "golang")
    fetcher.Start()

    // Start server
    log.Fatal(r.Run(":8080"))
}
package worker

import (
    "log"
    "time"
    
    "github.com/easily-mistaken/youtube-api-assignment/internal/youtube"
)

type VideoFetcher struct {
    client      *youtube.Client
    interval    time.Duration
    searchQuery string
}

func NewVideoFetcher(client *youtube.Client, interval time.Duration, query string) *VideoFetcher {
    return &VideoFetcher{
        client:      client,
        interval:    interval,
        searchQuery: query,
    }
}

func (vf *VideoFetcher) Start() {
    go vf.run()
}

func (vf *VideoFetcher) run() {
    ticker := time.NewTicker(vf.interval)
    defer ticker.Stop()

    // Initial fetch
    vf.fetchVideos()

    for range ticker.C {
        vf.fetchVideos()
    }
}

func (vf *VideoFetcher) fetchVideos() {
    resp, err := vf.client.FetchVideos(vf.searchQuery)
    if err != nil {
        log.Printf("Error fetching videos: %v", err)
        return
    }

    // Log the fetched videos count
    log.Printf("Fetched %d videos for query: %s", len(resp.Items), vf.searchQuery)
}
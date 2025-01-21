package worker

import (
    "time"
)

type VideoFetcher struct {
    client      interface{}
    store       interface{}
    interval    time.Duration
    searchQuery string
}

func NewVideoFetcher(interval time.Duration, query string) *VideoFetcher {
    return &VideoFetcher{
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

    for range ticker.C {
        // TODO: implement the actual fetching logic
    }
}
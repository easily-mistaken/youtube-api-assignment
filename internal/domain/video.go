package domain

import (
    "time"
)

type Video struct {
    VideoID      string    `gorm:"primaryKey" json:"video_id"`
    Title        string    `json:"title"`
    Description  string    `json:"description"`
    PublishedAt  time.Time `json:"published_at"`
    ThumbnailURL string    `json:"thumbnail_url"`
    ChannelTitle string    `json:"channel_title"`
    FetchedAt    time.Time `json:"fetched_at"`
}
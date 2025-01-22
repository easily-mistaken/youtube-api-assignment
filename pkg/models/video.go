package models

import (
	"time"

	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	VideoID      string    `gorm:"uniqueIndex" json:"videoId"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	PublishedAt  time.Time `gorm:"index" json:"publishedAt"`
	ThumbnailURL string    `json:"thumbnailUrl"`
	ViewCount    int64     `json:"viewCount"`
	Channel      string    `json:"channel"`
} 
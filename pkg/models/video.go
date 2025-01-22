package models

import (
	"time"

)

type Video struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	PublishedAt  time.Time `gorm:"index" json:"publishedAt"`
	ThumbnailURL string    `json:"thumbnailUrl"`
	Channel      string    `json:"channel"`
} 
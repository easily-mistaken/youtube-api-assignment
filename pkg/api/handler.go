package api

import (
	"github.com/easily-mistaken/youtube-api-assignment/pkg/models"

	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	db *gorm.DB
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{db: db}
}

func (h *Handler) GetVideos(c *gin.Context) {
	var videos []models.Video
	
	query := h.db.Model(&models.Video{}).
		Order("published_at DESC").
		Scopes(paginate(c))
		
	if channel := c.Query("channel"); channel != "" {
		query = query.Where("channel = ?", channel)
	}
	
	if err := query.Find(&videos).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(200, videos)
}

func paginate(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
		pageSize, _ := strconv.Atoi(c.DefaultQuery("size", "10"))
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
} 
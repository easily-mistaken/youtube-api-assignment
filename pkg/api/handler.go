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
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	
	// Validate pagination parameters
	if limit <= 0 {
		limit = 10
	}
	if page <= 0 {
		page = 1
	}

	var videos []models.Video
	var total int64
	
	query := h.db.Model(&models.Video{}).Order("published_at DESC")
		
	// Apply channel filter if provided
	if channel := c.Query("channel"); channel != "" {
		query = query.Where("channel = ?", channel)
	}

	query.Count(&total)
	
	// Apply pagination
	offset := (page - 1) * limit
	if err := query.Offset(offset).Limit(limit).Find(&videos).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(200, gin.H{
		"page":   page,
		"limit":  limit,
		"total":  total,
		"videos": videos,
	})
}
package youtube

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/easily-mistaken/youtube-api-assignment/pkg/models"
	"gorm.io/gorm"
)

type Client struct {
	apiKeys []string
	currentKeyIndex int
	httpClient *http.Client
}

func NewClient(apiKeys []string) *Client {
	return &Client{
		apiKeys: apiKeys,
		httpClient: &http.Client{Timeout: 10 * time.Second},
	}
}

func StartPeriodicFetch(db *gorm.DB, client *Client, query string, interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		videos, err := client.FetchVideos(query)
		if err != nil {
			fmt.Printf("Error fetching videos: %v\n", err)
			continue
		}
		
		// Store videos in database
		for _, video := range videos {
			if err := db.Create(&video).Error; err != nil {
				fmt.Printf("Error storing video: %v\n", err)
			}
		}
	}
}

func (c *Client) FetchVideos(query string) ([]models.Video, error) {
	apiKey := c.apiKeys[c.currentKeyIndex]
	url := fmt.Sprintf("https://www.googleapis.com/youtube/v3/search?part=snippet&q=%s&type=video&order=date&key=%s&maxResults=50", query, apiKey)
	
	resp, err := c.httpClient.Get(url)
	if err != nil {
		c.rotateKey()
		return nil, err
	}
	defer resp.Body.Close()
	
	if resp.StatusCode == 403 {
		c.rotateKey()
		return nil, fmt.Errorf("API quota exceeded")
	}
	
	var result struct {
		Items []struct {
			Snippet struct {
				Title       string    `json:"title"`
				Description string    `json:"description"`
				PublishedAt string    `json:"publishedAt"`
				Thumbnails  struct {
					High struct {
						URL string `json:"url"`
					} `json:"high"`
				} `json:"thumbnails"`
				ChannelTitle string `json:"channelTitle"`
			} `json:"snippet"`
		} `json:"items"`
	}
	
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	
	var videos []models.Video
	for _, item := range result.Items {
		publishedAt, _ := time.Parse(time.RFC3339, item.Snippet.PublishedAt)
		videos = append(videos, models.Video{
			Title:        item.Snippet.Title,
			Description:  item.Snippet.Description,
			PublishedAt:  publishedAt,
			ThumbnailURL: item.Snippet.Thumbnails.High.URL,
			Channel:      item.Snippet.ChannelTitle,
		})
	}
	
	return videos, nil
}

func (c *Client) rotateKey() {
	c.currentKeyIndex = (c.currentKeyIndex + 1) % len(c.apiKeys)
} 
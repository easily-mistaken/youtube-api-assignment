package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"gorm.io/gorm"

	"github.com/easily-mistaken/youtube-api-assignment/pkg/models"
)

type ApiResponse struct {
	Items []struct {
		Snippet struct {
			Title        string `json:"title"`
			Description  string `json:"description"`
			PublishedAt  string `json:"publishedAt"`
			ChannelTitle string `json:"channelTitle"`
			Thumbnails   struct {
				High struct {
					URL string `json:"url"`
				} `json:"high"`
			} `json:"thumbnails"`
		} `json:"snippet"`
	} `json:"items"`
}

type YoutubeService struct {
	DB        *gorm.DB
	apiKeys   []string
	keyIndex  int
	client    *http.Client
}

func NewYoutubeService(db *gorm.DB) *YoutubeService {
	apiKeys := []string{
		os.Getenv("API_KEY_1"),
		os.Getenv("API_KEY_2"),
	}

	return &YoutubeService{
		DB:      db,
		apiKeys: apiKeys,
		client:  &http.Client{Timeout: 10 * time.Second},
	}
}

func (s *YoutubeService) FetchVideos(query string) (string, error) {
	url := "https://www.googleapis.com/youtube/v3/search"

	if len(s.apiKeys) == 0 {
		return "", fmt.Errorf("no API keys found in environment variables")
	}

	for i := 0; i < len(s.apiKeys); i++ {
		apiKey := s.apiKeys[s.keyIndex]
		
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return "", fmt.Errorf("failed to create request: %v", err)
		}

		q := req.URL.Query()
		q.Add("part", "snippet")
		q.Add("q", query)
		q.Add("type", "video")
		q.Add("order", "date")
		q.Add("key", apiKey)
		q.Add("maxResults", "10")
		req.URL.RawQuery = q.Encode()

		resp, err := s.client.Do(req)
		if err != nil {
			s.rotateKey()
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode == 403 {
			fmt.Printf("API quota exceeded for key: %s\n", apiKey)
			s.rotateKey()
			continue
		}

		var apiResp ApiResponse
		if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
			return "", fmt.Errorf("error unmarshalling response: %v", err)
		}

		if len(apiResp.Items) == 0 {
			return "No results found for query: " + query, nil
		}

		for _, item := range apiResp.Items {
			publishedAt, err := time.Parse(time.RFC3339, item.Snippet.PublishedAt)
			if err != nil {
				continue
			}

			video := models.Video{
				Title:        item.Snippet.Title,
				Description:  item.Snippet.Description,
				PublishedAt:  publishedAt,
				ThumbnailURL: item.Snippet.Thumbnails.High.URL,
				Channel:      item.Snippet.ChannelTitle,
			}

			if err := s.DB.Create(&video).Error; err != nil {
				fmt.Printf("Error storing video: %v\n", err)
			}
		}

		return fmt.Sprintf("Successfully fetched and inserted %d videos", len(apiResp.Items)), nil
	}

	return "", fmt.Errorf("all API keys have exceeded their quota")
}

func (s *YoutubeService) rotateKey() {
	s.keyIndex = (s.keyIndex + 1) % len(s.apiKeys)
} 
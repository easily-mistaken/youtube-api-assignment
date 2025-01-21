package youtube

import (
    "encoding/json"
    "fmt"
    "net/http"
    "sync"
    "time"
)

type Client struct {
    apiKeys    []string
    currentKey int
    keyMutex   sync.Mutex
    httpClient *http.Client
}

type SearchResponse struct {
    Items []struct {
        ID struct {
            VideoID string `json:"videoId"`
        } `json:"id"`
        Snippet struct {
            PublishedAt time.Time `json:"publishedAt"`
            Title       string    `json:"title"`
            Description string    `json:"description"`
            ChannelTitle string   `json:"channelTitle"`
            Thumbnails  struct {
                High struct {
                    URL string `json:"url"`
                } `json:"high"`
            } `json:"thumbnails"`
        } `json:"snippet"`
    } `json:"items"`
}

func NewClient(apiKeys []string) *Client {
    return &Client{
        apiKeys: apiKeys,
        httpClient: &http.Client{
            Timeout: 10 * time.Second,
        },
    }
}

func (c *Client) getNextKey() string {
    c.keyMutex.Lock()
    defer c.keyMutex.Unlock()
    
    c.currentKey = (c.currentKey + 1) % len(c.apiKeys)
    return c.apiKeys[c.currentKey]
}

func (c *Client) FetchVideos(query string) (*SearchResponse, error) {
    apiKey := c.getNextKey()
    
    url := fmt.Sprintf(
        "https://www.googleapis.com/youtube/v3/search?part=snippet&q=%s&type=video&order=date&maxResults=50&key=%s",
        query, apiKey,
    )

    resp, err := c.httpClient.Get(url)
    if err != nil {
        return nil, fmt.Errorf("failed to fetch videos: %v", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode == http.StatusForbidden {
        return nil, fmt.Errorf("API quota exceeded for key: %s", apiKey)
    }

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("YouTube API error: %d", resp.StatusCode)
    }

    var searchResp SearchResponse
    if err := json.NewDecoder(resp.Body).Decode(&searchResp); err != nil {
        return nil, fmt.Errorf("failed to decode response: %v", err)
    }

    return &searchResp, nil
}
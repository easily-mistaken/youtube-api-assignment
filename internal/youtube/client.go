package youtube

import (
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
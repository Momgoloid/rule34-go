package rule34

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	UserID     string
	APIKey     string
	baseURL    string
	httpClient *http.Client
}

func New(id string, apiKey string) *Client {
	return &Client{
		UserID:  id,
		APIKey:  apiKey,
		baseURL: "https://api.rule34.xxx/index.php?page=dapi&q=index",
		httpClient: &http.Client{
			Timeout: time.Second * 5,
		},
	}
}

func (c *Client) Posts() *PostsRequestBuilder {
	return &PostsRequestBuilder{
		options: PostsOptions{
			Tags:      make([]string, 0),
			BlackList: make([]string, 0),
		},
		client: c,
	}
}

func (c *Client) Comments() {}

func (c *Client) Tags() {}

func (c *Client) doRequest(url string) ([]byte, error) {
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("can't do request: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("can't read body: %v", err)
	}

	return body, nil
}

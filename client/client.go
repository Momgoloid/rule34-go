// Package rule34 provides a client for interacting with the rule34.xxx API.
package rule34

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

// Client represents a client for the rule34.xxx API.
// It holds user credentials and an HTTP client to perform requests.
type Client struct {
	UserID     string
	APIKey     string
	baseURL    string
	httpClient *http.Client
}

// New creates a new instance of the rule34 client.
// It requires a user ID and an API key for authentication.
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

// Posts returns a PostsRequestBuilder for building a request to fetch posts.
// This builder allows for chaining methods to specify query parameters like tags, limits, etc.
func (c *Client) Posts() *PostsRequestBuilder {
	return &PostsRequestBuilder{
		options: PostsOptions{
			Tags:      make([]string, 0),
			BlackList: make([]string, 0),
		},
		client: c,
	}
}

// Comments is intended to return a CommentRequestBuilder for fetching comments.
// TODO: implement
func (c *Client) Comments() {}

// Tags is intended to return a TagRequestBuilder for fetching tags.
// TODO: implement
func (c *Client) Tags() {}

// doRequest performs an HTTP GET request to the specified URL. It returns an
// error for non-200 status codes, network issues, or problems reading the response body.
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

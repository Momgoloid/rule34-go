package rule34

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/Momgoloid/rule34-go/models"
)

const (
	baseURL = "https://api.rule34.xxx/index.php?page=dapi"
)

type Client struct {
	UserID int
	APIKey string
	http.Client
}

func New(id int, apiKey string) *Client {
	return &Client{
		UserID: id,
		APIKey: apiKey,
	}
}

func (c *Client) GetPost(postID int) (*models.PostXML, error) {
	const fn = "client.GetPost"

	post, err := c.getPost(postID)
	if err != nil {
		return nil, fmt.Errorf("%s: failed to get post: %v", fn, err)
	}

	return post, nil
}

func (c *Client) getPost(postID int) (*models.PostXML, error) {
	getPostURL, err := c.buildGetPostURL(postID)
	if err != nil {
		return nil, fmt.Errorf("failed to build get post url: %v", err)
	}

	postBytes, err := c.doRequest(getPostURL)
	if err != nil {
		return nil, fmt.Errorf("failed to do request: %v", err)
	}

	post, err := unmarshalPost(postBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal post: %v", err)
	}

	return post, nil
}

func (c *Client) buildGetPostURL(postID int) (string, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return "", fmt.Errorf("can't parse base URL: %w", err)
	}

	q := u.Query()
	q.Set("s", "post")
	q.Set("q", "index")
	q.Set("id", strconv.Itoa(postID))
	q.Set("user_id", strconv.Itoa(c.UserID))
	q.Set("api_key", c.APIKey)

	u.RawQuery = q.Encode()

	return u.String(), nil
}

func (c *Client) doRequest(url string) ([]byte, error) {
	resp, err := c.Get(url)
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

func unmarshalPost(postBytes []byte) (*models.PostXML, error) {
	var (
		posts models.PostsXML
		post  models.PostXML
	)

	postData := bytes.NewBuffer(postBytes)

	d := xml.NewDecoder(postData)

	for t, _ := d.Token(); t != nil; t, _ = d.Token() {
		switch se := t.(type) {
		case xml.StartElement:
			if se.Name.Local == models.PostElementName {
				err := d.DecodeElement(&post, &se)
				if err != nil {
					return nil, fmt.Errorf("can't decode element: %v", err)
				}
				posts.Post = append(posts.Post, post)
			}
		}
	}

	return &posts.Post[0], nil
}

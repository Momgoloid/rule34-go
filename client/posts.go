package rule34

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/Momgoloid/rule34-go/client/rating"
	"github.com/Momgoloid/rule34-go/models"
)

var (
	ErrNonPositivePostID       = errors.New("post id can't be less than or equal to zero")
	ErrNonPositiveLimit        = errors.New("limit can't be less than or equal to zero")
	ErrNonPositivePageNumber   = errors.New("page number can't be less than or equal to zero")
	ErrUnknownRating           = errors.New("unknown rating was given")
	ErrNonPositiveParentPostID = errors.New("parent post id can't be less than or equal to zero")
	// ErrNegativeScore           = errors.New("score can't be negative")
)

type PostsOptions struct {
	PostID       int
	Limit        int
	PageNumber   int
	Tags         []string
	BlackList    []string
	FilterAI     bool
	Rating       rating.Rating
	ParentPostID int
	// Score        int
}

type PostsRequestBuilder struct {
	options PostsOptions
	client  *Client
	errors  []error
}

func (b *PostsRequestBuilder) PostID(postID int) *PostsRequestBuilder {
	if postID <= 0 {
		b.errors = append(b.errors, ErrNonPositivePostID)
		return b
	}

	b.options.PostID = postID
	return b
}

func (b *PostsRequestBuilder) Limit(limit int) *PostsRequestBuilder {
	if limit <= 0 {
		b.errors = append(b.errors, ErrNonPositiveLimit)
		return b
	}

	b.options.Limit = limit
	return b
}

func (b *PostsRequestBuilder) PageNumber(pageNumber int) *PostsRequestBuilder {
	if pageNumber <= 0 {
		b.errors = append(b.errors, ErrNonPositivePageNumber)
		return b
	}

	b.options.PageNumber = pageNumber
	return b
}

func (b *PostsRequestBuilder) Tags(tags ...string) *PostsRequestBuilder {
	b.options.Tags = append(b.options.Tags, tags...)
	return b
}

func (b *PostsRequestBuilder) BlackList(blackList ...string) *PostsRequestBuilder {
	b.options.BlackList = append(b.options.BlackList, blackList...)
	return b
}

func (b *PostsRequestBuilder) FilterAI() *PostsRequestBuilder {
	b.options.FilterAI = true
	return b
}

func (b *PostsRequestBuilder) Rating(r rating.Rating) *PostsRequestBuilder {
	if r.String() == "" {
		b.errors = append(b.errors, ErrUnknownRating)
		return b
	}

	b.options.Rating = r
	return b
}

func (b *PostsRequestBuilder) ParentPostID(parentPostID int) *PostsRequestBuilder {
	if parentPostID <= 0 {
		b.errors = append(b.errors, ErrNonPositiveParentPostID)
		return b
	}

	b.options.ParentPostID = parentPostID
	return b
}

// func (b *PostsRequestBuilder) Score(score int) *PostsRequestBuilder {
// 	if score < 0 {
// 		b.errors = append(b.errors, ErrNegativeScore)
// 		return b
// 	}

// 	b.options.Score = score
// 	return b
// }

func (b *PostsRequestBuilder) Find() (*models.Posts, error) {
	if len(b.errors) != 0 {
		err := errors.Join(b.errors...)
		return nil, fmt.Errorf("invalid arguments: %v", err)
	}

	url, err := b.buildURL()
	if err != nil {
		return nil, fmt.Errorf("failed to build url: %v", err)
	}

	postsBytes, err := b.client.doRequest(url)
	if err != nil {
		return nil, fmt.Errorf("failed to do get posts request: %v", err)
	}

	posts, err := unmarshalPosts(postsBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal posts: %v", err)
	}

	return posts, nil
}

func (b *PostsRequestBuilder) buildURL() (string, error) {
	u, err := url.Parse(b.client.baseURL)
	if err != nil {
		return "", fmt.Errorf("can't parse base URL: %w", err)
	}

	q := u.Query()
	b.addOptions(&q)
	u.RawQuery = q.Encode()

	return u.String(), nil
}

func (b *PostsRequestBuilder) addOptions(q *url.Values) {
	q.Set("s", "post")

	postID := strconv.Itoa(b.options.PostID)
	if postID != "0" {
		q.Set("id", postID)
	}

	limit := strconv.Itoa(b.options.Limit)
	if limit != "0" {
		q.Set("limit", limit)
	}

	pageNumber := strconv.Itoa(b.options.PageNumber)
	if pageNumber != "0" {
		q.Set("pid", pageNumber)
	}

	tags := b.convertTags()
	q.Set("tags", tags)

	q.Set("user_id", b.client.UserID)
	q.Set("api_key", b.client.APIKey)
}

func (b *PostsRequestBuilder) convertTags() string {
	sb := strings.Builder{}

	for _, tag := range b.options.Tags {
		sb.WriteString(fmt.Sprintf("%s ", tag))
	}

	for _, tag := range b.options.BlackList {
		sb.WriteString(fmt.Sprintf("-%s ", tag))
	}

	if b.options.FilterAI {
		sb.WriteString("-ai_generated ")
	}

	if b.options.Rating != 0 {
		sb.WriteString(fmt.Sprintf("rating:%s ", b.options.Rating.String()))
	}

	parentPostID := strconv.Itoa(b.options.ParentPostID)
	if parentPostID != "0" {
		sb.WriteString(fmt.Sprintf("parent:%s ", parentPostID))
	}

	// if b.options.Score != 0 {
	// 	sb.WriteString(fmt.Sprintf(" score: "))
	// }

	return sb.String()
}

func unmarshalPosts(postsBytes []byte) (*models.Posts, error) {
	var (
		posts models.Posts
		post  models.Post
	)

	postData := bytes.NewBuffer(postsBytes)

	d := xml.NewDecoder(postData)

	for t, _ := d.Token(); t != nil; t, _ = d.Token() {
		switch se := t.(type) {
		case xml.StartElement:
			if se.Name.Local == models.PostElementName {
				err := d.DecodeElement(&post, &se)
				if err != nil {
					return nil, fmt.Errorf("can't decode element: %v", err)
				}
				posts.Posts = append(posts.Posts, post)
			}
		}
	}

	return &posts, nil
}

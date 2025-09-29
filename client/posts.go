// Package rule34 provides a client for interacting with the rule34.xxx API.
package rule34

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/Momgoloid69/rule34-go/customTypes/filtering"
	"github.com/Momgoloid69/rule34-go/customTypes/operators"
	"github.com/Momgoloid69/rule34-go/customTypes/rating"
	"github.com/Momgoloid69/rule34-go/customTypes/sorting"
	"github.com/Momgoloid69/rule34-go/models"
)

// Pre-defined errors for the PostsRequestBuilder.
var (
	// ErrNonPositivePostID is returned when a non-positive post ID is provided.
	ErrNonPositivePostID = errors.New("post id can't be less than or equal to zero")
	// ErrNonPositiveLimit is returned when a non-positive limit is provided.
	ErrNonPositiveLimit = errors.New("limit can't be less than or equal to zero")
	// ErrNonPositivePageNumber is returned when a non-positive page number is provided.
	ErrNonPositivePageNumber = errors.New("page number can't be less than or equal to zero")
	// ErrUnknownRating is returned when an invalid rating is provided.
	ErrUnknownRating = errors.New("unknown rating was given")
	// ErrUnknownSortingType is returned when an invalid sorting type is provided.
	ErrUnknownSortingType = errors.New("unknown sorting type was given")
	// ErrUnknownFilteringType is returned when an invalid filtering type is provided.
	ErrUnknownFilteringType = errors.New("unknown filtering type was given")
	// ErrUnknownOperation is returned when an invalid operator is provided for filtering.
	ErrUnknownOperation = errors.New("unknown operator was given")
	// ErrNonPositiveParentPostID is returned when a non-positive parent post ID is provided.
	ErrNonPositiveParentPostID = errors.New("parent post id can't be less than or equal to zero")
	// ErrSortByWasNotCalled is returned when Asc() or Desc() is called before SortBy().
	ErrSortByWasNotCalled = errors.New("sort by was not called")
	// ErrSortByWasCalledTwiceOrMore is returned when SortBy() is called multiple times on the same builder.
	ErrSortByWasCalledTwiceOrMore = errors.New("sort by was called twice or more")
	// ErrTwoSortingOrders is returned when both Asc() and Desc() are called on the same builder.
	ErrTwoSortingOrders = errors.New("both sorting orders was used")
)

// PostsRequestBuilder is a builder for creating and executing API requests for posts.
type PostsRequestBuilder struct {
	options PostsOptions
	client  *Client
	errors  []error
}

// PostsOptions holds all the configurable parameters for a posts API request.
type PostsOptions struct {
	PostID              int
	Limit               int
	PageNumber          int
	Tags                []string
	BlackList           []string
	FilterAI            bool
	Rating              rating.Rating
	ParentPostID        int
	FilteringConditions []filtering.Condition
	DoSort              bool
	SortableType        sorting.Type
	SortingOrder        string
}

// PostID sets the specific post ID to retrieve.
func (b *PostsRequestBuilder) PostID(postID int) *PostsRequestBuilder {
	if postID <= 0 {
		b.errors = append(b.errors, ErrNonPositivePostID)
		return b
	}

	b.options.PostID = postID
	return b
}

// Limit sets the maximum number of posts to retrieve.
// The API has a hard limit of 1000.
func (b *PostsRequestBuilder) Limit(limit int) *PostsRequestBuilder {
	if limit <= 0 {
		b.errors = append(b.errors, ErrNonPositiveLimit)
		return b
	}

	b.options.Limit = limit
	return b
}

// PageNumber sets the page number for pagination.
func (b *PostsRequestBuilder) PageNumber(pageNumber int) *PostsRequestBuilder {
	if pageNumber <= 0 {
		b.errors = append(b.errors, ErrNonPositivePageNumber)
		return b
	}

	b.options.PageNumber = pageNumber
	return b
}

// Tags adds search tags to the request.
// Multiple calls to this method will append tags.
func (b *PostsRequestBuilder) Tags(tags ...string) *PostsRequestBuilder {
	b.options.Tags = append(b.options.Tags, tags...)
	return b
}

// BlackList adds tags to be excluded from the search results.
// These tags will be prefixed with a '-' in the final query.
func (b *PostsRequestBuilder) BlackList(blackList ...string) *PostsRequestBuilder {
	b.options.BlackList = append(b.options.BlackList, blackList...)
	return b
}

// FilterAI adds a tag to exclude AI-generated content from the search results.
func (b *PostsRequestBuilder) FilterAI() *PostsRequestBuilder {
	b.options.FilterAI = true
	return b
}

// Rating sets the content rating for the search.
func (b *PostsRequestBuilder) Rating(r rating.Rating) *PostsRequestBuilder {
	if !r.IsValid() {
		b.errors = append(b.errors, ErrUnknownRating)
		return b
	}

	b.options.Rating = r
	return b
}

// ParentPostID searches for posts that have the given post ID as a parent.
func (b *PostsRequestBuilder) ParentPostID(parentPostID int) *PostsRequestBuilder {
	if parentPostID <= 0 {
		b.errors = append(b.errors, ErrNonPositiveParentPostID)
		return b
	}

	b.options.ParentPostID = parentPostID
	return b
}

// Where adds a filtering condition to the request.
// For example, Where(filtering.Score, operators.GreaterEqual, 10) finds posts with a score >= 10.
func (b *PostsRequestBuilder) Where(ft filtering.Type, op operators.Operator, arg int) *PostsRequestBuilder {
	if !ft.IsValid() {
		b.errors = append(b.errors, ErrUnknownFilteringType)
		return b
	}

	if !op.IsValid() {
		b.errors = append(b.errors, ErrUnknownOperation)
		return b
	}

	filteringCondition := filtering.Condition{
		FilteringType: ft,
		Operation:     op,
		Argument:      arg,
	}

	b.options.FilteringConditions = append(b.options.FilteringConditions, filteringCondition)

	return b
}

// SortBy specifies the field to sort the results by.
// Must be called before Asc() or Desc(). Defaults to descending order if neither is called.
func (b *PostsRequestBuilder) SortBy(sortableType sorting.Type) *PostsRequestBuilder {
	if b.options.DoSort {
		b.errors = append(b.errors, ErrSortByWasCalledTwiceOrMore)
		return b
	}

	if !sortableType.IsValid() {
		b.errors = append(b.errors, ErrUnknownSortingType)
		return b
	}

	b.options.DoSort = true
	b.options.SortableType = sortableType
	return b
}

// Asc sets the sorting order to ascending.
// Must be called after SortBy().
func (b *PostsRequestBuilder) Asc() *PostsRequestBuilder {
	if !b.checkDoSort() {
		return b
	}

	if b.options.SortingOrder != "" {
		b.errors = append(b.errors, ErrTwoSortingOrders)
		return b
	}

	b.options.SortingOrder = "asc"
	return b
}

// Desc sets the sorting order to descending.
// Must be called after SortBy().
func (b *PostsRequestBuilder) Desc() *PostsRequestBuilder {
	if !b.checkDoSort() {
		return b
	}

	if b.options.SortingOrder != "" {
		b.errors = append(b.errors, ErrTwoSortingOrders)
		return b
	}

	b.options.SortingOrder = "desc"
	return b
}

// Find executes the request to the API and returns the search results.
// It first validates any accumulated errors, then builds the URL, performs the request, and unmarshals the response.
func (b *PostsRequestBuilder) Find() (models.Posts, error) {
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

// buildURL constructs the final request URL from the builder's options.
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

// addOptions adds all configured options as query parameters to the URL.
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

	q.Set("json", "1")
}

// convertTags compiles all tags, blacklisted tags, and meta-tags into a single space-separated string.
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

	if b.options.Rating != "" {
		sb.WriteString(fmt.Sprintf("rating:%s ", b.options.Rating))
	}

	parentPostID := strconv.Itoa(b.options.ParentPostID)
	if parentPostID != "0" {
		sb.WriteString(fmt.Sprintf("parent:%s ", parentPostID))
	}

	if len(b.options.FilteringConditions) != 0 {
		for _, fc := range b.options.FilteringConditions {
			sb.WriteString(fmt.Sprintf("%s:%s%d ", fc.FilteringType, fc.Operation, fc.Argument))
		}
	}

	if b.options.DoSort {
		if b.options.SortingOrder != "" {
			sb.WriteString(fmt.Sprintf("sort:%s:%s ", b.options.SortableType, b.options.SortingOrder))
		} else {
			sb.WriteString(fmt.Sprintf("sort:%s:desc ", b.options.SortableType))
		}
	}

	return sb.String()
}

// unmarshalPosts parses the JSON response body into a slice of Post models.
func unmarshalPosts(postsBytes []byte) (models.Posts, error) {
	if len(postsBytes) == 0 {
		return models.Posts{}, nil
	}

	var posts models.Posts

	err := json.Unmarshal(postsBytes, &posts)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal posts: %v", err)
	}

	return posts, nil
}

// checkDoSort ensures that SortBy has been called before a sorting order method (Asc/Desc) is used.
func (b *PostsRequestBuilder) checkDoSort() bool {
	if !b.options.DoSort {
		b.errors = append(b.errors, ErrSortByWasNotCalled)
		return false
	}

	return true
}

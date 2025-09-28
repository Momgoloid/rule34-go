package rule34

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/Momgoloid/rule34-go/client/filtering"
	"github.com/Momgoloid/rule34-go/client/operators"
	"github.com/Momgoloid/rule34-go/client/rating"
	"github.com/Momgoloid/rule34-go/client/sorting"
	"github.com/Momgoloid/rule34-go/models"
)

var (
	ErrNonPositivePostID          = errors.New("post id can't be less than or equal to zero")
	ErrNonPositiveLimit           = errors.New("limit can't be less than or equal to zero")
	ErrNonPositivePageNumber      = errors.New("page number can't be less than or equal to zero")
	ErrUnknownRating              = errors.New("unknown rating was given")
	ErrUnknownSortingType         = errors.New("unknown sorting type was given")
	ErrUnknownFilteringType       = errors.New("unknown filtering type was given")
	ErrUnknownOperation           = errors.New("unknown operator was given")
	ErrNonPositiveParentPostID    = errors.New("parent post id can't be less than or equal to zero")
	ErrSortByWasNotCalled         = errors.New("sort by was not called")
	ErrSortByWasCalledTwiceOrMore = errors.New("sort by was called twice or more")
	ErrTwoSortingOrders           = errors.New("both sorting orders was used")
	ErrNotFound                   = errors.New("no posts was found")
)

type PostsRequestBuilder struct {
	options PostsOptions
	client  *Client
	errors  []error
}

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
	if !r.IsValid() {
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
		Operation:    op,
		Argument:      arg,
	}

	b.options.FilteringConditions = append(b.options.FilteringConditions, filteringCondition)

	return b
}

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

	q.Set("json", "1")
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

func (b *PostsRequestBuilder) checkDoSort() bool {
	if !b.options.DoSort {
		b.errors = append(b.errors, ErrSortByWasNotCalled)
		return false
	}

	return true
}

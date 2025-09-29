package rule34

import (
	"encoding/xml"
	"fmt"
	"time"

	"encoding/json"
	"strings"
)

// Posts is a slice of Post objects, representing a collection of posts
// returned by the API, typically in JSON format.
type Posts []Post

// Post represents a single post object and its associated metadata.
// The struct tags map the JSON keys from the API response to the fields.
type Post struct {
	PreviewURL   string    `json:"preview_url"`
	SampleURL    string    `json:"sample_url"`
	FileURL      string    `json:"file_url"`
	Directory    int       `json:"directory"`
	Hash         string    `json:"hash"`
	Width        int       `json:"width"`
	Height       int       `json:"height"`
	ID           int       `json:"id"`
	Image        string    `json:"image"`
	Change       int       `json:"change"`
	Owner        string    `json:"owner"`
	ParentID     int       `json:"parent_id"`
	Rating       string    `json:"rating"`
	Sample       bool      `json:"sample"`
	SampleHeight int       `json:"sample_height"`
	SampleWidth  int       `json:"sample_width"`
	Score        int       `json:"score"`
	Tags         TagsSlice `json:"tags"` // Custom type to handle space-separated string
	Source       string    `json:"source"`
	Status       string    `json:"status"`
	HasNotes     bool      `json:"has_notes"`
	CommentCount int       `json:"comment_count"`
}

// Comments represents the top-level structure for a comments API response,
// typically returned in XML format.
type Comments struct {
	XMLName xml.Name  `xml:"comments"`
	Type    string    `xml:"type,attr"`
	Comment []Comment `xml:"comment"`
}

// Comment represents a single comment object with its associated metadata.
// The struct tags map the XML attributes from the API response to the fields.
type Comment struct {
	XMLName   xml.Name  `xml:"comment"`
	CreatedAt CreatedAt `xml:"created_at,attr"`
	PostID    string    `xml:"post_id,attr"`
	Body      string    `xml:"body,attr"`
	Creator   string    `xml:"creator,attr"`
	ID        string    `xml:"id,attr"`
	CreatorID string    `xml:"creator_id,attr"`
}

// Tags represents the top-level structure for a tags API response,
// typically returned in XML format.
type Tags struct {
	XMLName xml.Name `xml:"tags"`
	Type    string   `xml:"type,attr"`
	Tag     []Tag    `xml:"tag"`
}

// Tag represents a single tag object with its associated metadata such as its type,
// count, and name. The struct tags map the XML attributes from the API response to the fields.
type Tag struct {
	XMLName   xml.Name `xml:"tag"`
	Type      int      `xml:"type,attr"`
	Count     int      `xml:"count,attr"`
	Name      string   `xml:"name,attr"`
	Ambiguous bool     `xml:"ambiguous,attr"`
	ID        int      `xml:"id,attr"`
}

// TagsSlice is a custom type representing a slice of strings.
// It is used to unmarshal a single space-separated string of tags
// from API responses (both XML and JSON) into a proper string slice.
type TagsSlice []string

// UnmarshalXMLAttr implements the xml.UnmarshalerAttr interface.
// It takes a space-separated string of tags from an XML attribute,
// trims whitespace, and splits it into a slice of strings.
func (t *TagsSlice) UnmarshalXMLAttr(attr xml.Attr) error {
	tagsStr := strings.Trim(attr.Value, " ")
	tags := strings.Split(tagsStr, " ")

	*t = TagsSlice(tags)
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
// It first unmarshals the raw JSON data into a string, then
// trims whitespace and splits that string by spaces into a slice of strings.
func (t *TagsSlice) UnmarshalJSON(data []byte) error {
	var tagsStr string
	if err := json.Unmarshal(data, &tagsStr); err != nil {
		return err
	}

	tagsStr = strings.TrimSpace(tagsStr)

	tags := strings.Split(tagsStr, " ")

	*t = TagsSlice(tags)
	return nil
}

// CreatedAt is a custom time type that wraps time.Time to handle the specific
// date format ("Mon Jan 2 15:04:05 -0700 2006") found in XML attributes from the API.
type CreatedAt struct {
	time.Time
}

// createdAtFormat defines the specific layout string for parsing the date
// from the API's XML responses.
const createdAtFormat = "Mon Jan 2 15:04:05 -0700 2006"

// UnmarshalXMLAttr implements the xml.UnmarshalerAttr interface.
// It parses the date string from an XML attribute using the custom createdAtFormat
// and assigns the resulting time.Time value to the CreatedAt receiver.
func (cd *CreatedAt) UnmarshalXMLAttr(attr xml.Attr) error {
	dateString := attr.Value

	date, err := time.Parse(createdAtFormat, dateString)
	if err != nil {
		return fmt.Errorf("can't parse date: %v", err)
	}

	*cd = CreatedAt{date}
	return nil
}

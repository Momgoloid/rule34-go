// Package models defines the data structures that map to the API's responses.
package models

import (
	"encoding/xml"

	"github.com/Momgoloid69/rule34-go/customTypes/unmarshaling"
)

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
	XMLName   xml.Name               `xml:"comment"`
	CreatedAt unmarshaling.CreatedAt `xml:"created_at,attr"`
	PostID    string                 `xml:"post_id,attr"`
	Body      string                 `xml:"body,attr"`
	Creator   string                 `xml:"creator,attr"`
	ID        string                 `xml:"id,attr"`
	CreatorID string                 `xml:"creator_id,attr"`
}

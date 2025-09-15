package models

import "encoding/xml"

type Comments struct {
	XMLName xml.Name  `xml:"comments"`
	Type    string    `xml:"type,attr"`
	Comment []Comment `xml:"comment"`
}

type Comment struct {
	XMLName   xml.Name `xml:"comment"`
	CreatedAt string   `xml:"created_at,attr"`
	PostID    string   `xml:"post_id,attr"`
	Body      string   `xml:"body,attr"`
	Creator   string   `xml:"creator,attr"`
	ID        string   `xml:"id,attr"`
	CreatorID string   `xml:"creator_id,attr"`
}

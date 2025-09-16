package models

import (
	"encoding/xml"

	"github.com/Momgoloid/rule34-go/internal/customXML"
)

const PostElementName = "post"

type PostsXML struct {
	XMLName xml.Name  `xml:"posts"`
	Count   string    `xml:"count,attr"`
	Offset  string    `xml:"offset,attr"`
	Post    []PostXML `xml:"post"`
}

type PostXML struct {
	Height        int                 `xml:"height,attr"`
	Score         int                 `xml:"score,attr"`
	FileURL       string              `xml:"file_url,attr"`
	ParentID      string              `xml:"parent_id,attr"`
	SampleURL     string              `xml:"sample_url,attr"`
	SampleWidth   int                 `xml:"sample_width,attr"`
	SampleHeight  int                 `xml:"sample_height,attr"`
	PreviewURL    string              `xml:"preview_url,attr"`
	Rating        string              `xml:"rating,attr"`
	Tags          customXML.Tags      `xml:"tags,attr"`
	ID            int                 `xml:"id,attr"`
	Width         int                 `xml:"width,attr"`
	Change        int                 `xml:"change,attr"`
	Md5           string              `xml:"md5,attr"`
	CreatorID     int                 `xml:"creator_id,attr"`
	HasChildren   bool                `xml:"has_children,attr"`
	CreatedAt     customXML.CreatedAt `xml:"created_at,attr"`
	Status        string              `xml:"status,attr"`
	Source        string              `xml:"source,attr"`
	HasNotes      bool                `xml:"has_notes,attr"`
	HasComments   bool                `xml:"has_comments,attr"`
	PreviewWidth  int                 `xml:"preview_width,attr"`
	PreviewHeight int                 `xml:"preview_height,attr"`
}

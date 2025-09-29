package models

import "github.com/Momgoloid69/rule34-go/customTypes/unmarshaling"

type Posts []Post

type Post struct {
	PreviewURL   string            `json:"preview_url"`
	SampleURL    string            `json:"sample_url"`
	FileURL      string            `json:"file_url"`
	Directory    int               `json:"directory"`
	Hash         string            `json:"hash"`
	Width        int               `json:"width"`
	Height       int               `json:"height"`
	ID           int               `json:"id"`
	Image        string            `json:"image"`
	Change       int               `json:"change"`
	Owner        string            `json:"owner"`
	ParentID     int               `json:"parent_id"`
	Rating       string            `json:"rating"`
	Sample       bool              `json:"sample"`
	SampleHeight int               `json:"sample_height"`
	SampleWidth  int               `json:"sample_width"`
	Score        int               `json:"score"`
	Tags         unmarshaling.Tags `json:"tags"`
	Source       string            `json:"source"`
	Status       string            `json:"status"`
	HasNotes     bool              `json:"has_notes"`
	CommentCount int               `json:"comment_count"`
}

// Package unmarshaling provides custom types with special unmarshaling logic
// for handling non-standard data formats from the API.
package unmarshaling

import (
	"encoding/json"
	"encoding/xml"
	"strings"
)

// Tags is a custom type representing a slice of strings.
// It is used to unmarshal a single space-separated string of tags
// from API responses (both XML and JSON) into a proper string slice.
type Tags []string

// UnmarshalXMLAttr implements the xml.UnmarshalerAttr interface.
// It takes a space-separated string of tags from an XML attribute,
// trims whitespace, and splits it into a slice of strings.
func (t *Tags) UnmarshalXMLAttr(attr xml.Attr) error {
	tagsStr := strings.Trim(attr.Value, " ")
	tags := strings.Split(tagsStr, " ")

	*t = Tags(tags)
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
// It first unmarshals the raw JSON data into a string, then
// trims whitespace and splits that string by spaces into a slice of strings.
func (t *Tags) UnmarshalJSON(data []byte) error {
	var tagsStr string
	if err := json.Unmarshal(data, &tagsStr); err != nil {
		return err
	}

	tagsStr = strings.TrimSpace(tagsStr)

	tags := strings.Split(tagsStr, " ")

	*t = Tags(tags)
	return nil
}

// Package models defines the data structures that map to the API's responses.
package models

import "encoding/xml"

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

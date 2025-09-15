package models

import "encoding/xml"

type Tags struct {
	XMLName xml.Name `xml:"tags"`
	Type    string   `xml:"type,attr"`
	Tag     []Tag    `xml:"tag"`
}

type Tag struct {
	XMLName   xml.Name `xml:"tag"`
	Type      int      `xml:"type,attr"`
	Count     int      `xml:"count,attr"`
	Name      string   `xml:"name,attr"`
	Ambiguous bool     `xml:"ambiguous,attr"`
	ID        int      `xml:"id,attr"`
}

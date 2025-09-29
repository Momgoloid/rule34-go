// Package unmarshaling provides custom types with special unmarshaling logic
// for handling non-standard data formats from the API.
package unmarshaling

import (
	"encoding/xml"
	"fmt"
	"time"
)

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

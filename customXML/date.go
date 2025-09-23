package customXML

import (
	"encoding/xml"
	"fmt"
	"time"
)

type CreatedAt struct {
	time.Time
}

const createdAtFormat = "Mon Jan 2 15:04:05 -0700 2006"

func (cd *CreatedAt) UnmarshalXMLAttr(attr xml.Attr) error {
	dateString := attr.Value

	date, err := time.Parse(createdAtFormat, dateString)
	if err != nil {
		return fmt.Errorf("can't parse date: %v", err)
	}

	*cd = CreatedAt{date}
	return nil
}

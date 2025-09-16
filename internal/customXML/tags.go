package customXML

import (
	"encoding/xml"
	"strings"
)

type Tags []string

func (t *Tags) UnmarshalXMLAttr(attr xml.Attr) error {
	tagsStr := strings.Trim(attr.Value, " ")
	tags := strings.Split(tagsStr, " ")

	*t = Tags(tags)
	return nil
}

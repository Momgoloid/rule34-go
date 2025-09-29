package unmarshaling

import (
	"encoding/json"
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

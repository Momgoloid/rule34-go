// Package sorting defines the types used for specifying the sorting order in API requests.
package sorting

// Type represents a field that can be used for sorting search results.
type Type string

// Defines the valid sortable types for API requests.
const (
	ID      Type = "id"
	Score   Type = "score"
	Rating  Type = "rating"
	User    Type = "user"
	Height  Type = "height"
	Width   Type = "width"
	Parent  Type = "parent"
	Source  Type = "source"
	Updated Type = "updated"
)

// ValidTypes is a set of all valid sortable types for quick validation.
var ValidTypes = map[Type]struct{}{
	ID:      {},
	Score:   {},
	Rating:  {},
	User:    {},
	Height:  {},
	Width:   {},
	Parent:  {},
	Source:  {},
	Updated: {},
}

// IsValid checks if the sorting type is a valid, known type.
func (t Type) IsValid() bool {
	_, ok := ValidTypes[t]
	return ok
}

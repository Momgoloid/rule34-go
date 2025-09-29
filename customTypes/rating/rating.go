// Package rating defines the types for content ratings used in API requests.
package rating

// Rating represents the content rating of a post.
type Rating string

// Defines the valid content ratings.
const (
	Safe         Rating = "safe"
	Questionable Rating = "questionable"
	Explicit     Rating = "explicit"
)

// ValidRatings is a set of all valid rating types for quick validation.
var ValidRatings = map[Rating]struct{}{
	Safe:         {},
	Questionable: {},
	Explicit:     {},
}

// IsValid checks if the rating is a valid, known rating.
func (r Rating) IsValid() bool {
	_, ok := ValidRatings[r]
	return ok
}

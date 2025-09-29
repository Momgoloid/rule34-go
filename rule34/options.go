package rule34

// Condition represents a single filtering rule, combining a type, an operator, and an argument.
// For example, it can represent "score >= 10".
type Condition struct {
	FilteringType FilterType
	Operation     Operator
	Argument      int
}

// FilterType represents a field that can be used for filtering, such as score, width, or height.
type FilterType string

// Defines the valid filtering FilterTypes that can be used in API requests.
const (
	FilterByID      FilterType = "id"
	FilterByScore   FilterType = "score"
	FilterByHeight  FilterType = "height"
	FilterByWidth   FilterType = "width"
	FilterByParent  FilterType = "parent"
	FilterByUpdated FilterType = "updated"
)

// ValidFilterTypes is a set of all valid filtering FilterTypes for quick validation.
var ValidFilterTypes = map[FilterType]struct{}{
	FilterByID:      {},
	FilterByScore:   {},
	FilterByHeight:  {},
	FilterByWidth:   {},
	FilterByParent:  {},
	FilterByUpdated: {},
}

// IsValid checks if the filtering type is a valid, known type.
func (t FilterType) IsValid() bool {
	_, ok := ValidFilterTypes[t]
	return ok
}

// Operator represents a comparison operator, such as "=", ">=", or "<".
type Operator string

// Defines the valid comparison operators for filtering.
const (
	Equal        Operator = "="
	NotEqual     Operator = "!="
	Greater      Operator = ">"
	GreaterEqual Operator = ">="
	Less         Operator = "<"
	LessEqual    Operator = "<="
)

// ValidOperators is a set of all valid comparison operators for quick validation.
var ValidOperators = map[Operator]struct{}{
	Equal:        {},
	NotEqual:     {},
	Greater:      {},
	GreaterEqual: {},
	Less:         {},
	LessEqual:    {},
}

// IsValid checks if the operator is a valid, known operator.
func (o Operator) IsValid() bool {
	_, ok := ValidOperators[o]
	return ok
}

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

// SortableType represents a field that can be used for sorting search results.
type SortableType string

// Defines the valid sortable types for API requests.
const (
	SortByID      SortableType = "id"
	SortByScore   SortableType = "score"
	SortByRating  SortableType = "rating"
	SortByUser    SortableType = "user"
	SortByHeight  SortableType = "height"
	SortByWidth   SortableType = "width"
	SortByParent  SortableType = "parent"
	SortBySource  SortableType = "source"
	SortByUpdated SortableType = "updated"
)

// ValidSortableTypes is a set of all valid sortable types for quick validation.
var ValidSortableTypes = map[SortableType]struct{}{
	SortByID:      {},
	SortByScore:   {},
	SortByRating:  {},
	SortByUser:    {},
	SortByHeight:  {},
	SortByWidth:   {},
	SortByParent:  {},
	SortBySource:  {},
	SortByUpdated: {},
}

// IsValid checks if the sorting type is a valid, known type.
func (t SortableType) IsValid() bool {
	_, ok := ValidSortableTypes[t]
	return ok
}

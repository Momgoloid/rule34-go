// Package filtering defines the types used for creating filtering conditions in API requests.
package filtering

import "github.com/Momgoloid69/rule34-go/customTypes/operators"

// Condition represents a single filtering rule, combining a type, an operator, and an argument.
// For example, it can represent "score >= 10".
type Condition struct {
	FilteringType Type
	Operation     operators.Operator
	Argument      int
}

// Type represents a field that can be used for filtering, such as score, width, or height.
type Type string

// Defines the valid filtering types that can be used in API requests.
const (
	ID      Type = "id"
	Score   Type = "score"
	Height  Type = "height"
	Width   Type = "width"
	Parent  Type = "parent"
	Updated Type = "updated"
)

// ValidTypes is a set of all valid filtering types for quick validation.
var ValidTypes = map[Type]struct{}{
	ID:      {},
	Score:   {},
	Height:  {},
	Width:   {},
	Parent:  {},
	Updated: {},
}

// IsValid checks if the filtering type is a valid, known type.
func (t Type) IsValid() bool {
	_, ok := ValidTypes[t]
	return ok
}

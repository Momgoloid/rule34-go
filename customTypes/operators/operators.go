// Package operators defines the comparison operators used in filtering conditions.
package operators

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

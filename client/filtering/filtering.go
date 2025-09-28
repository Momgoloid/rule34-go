package filtering

import "github.com/Momgoloid/rule34-go/client/operators"

type Condition struct {
	FilteringType Type
	Operation     operators.Operator
	Argument      int
}

type Type string

const (
	ID      Type = "id"
	Score   Type = "score"
	Rating  Type = "rating"
	Height  Type = "height"
	Width   Type = "width"
	Parent  Type = "parent"
	Updated Type = "updated"
)

var ValidTypes = map[Type]struct{}{
	ID:      {},
	Score:   {},
	Rating:  {},
	Height:  {},
	Width:   {},
	Parent:  {},
	Updated: {},
}

func (t Type) IsValid() bool {
	_, ok := ValidTypes[t]
	return ok
}

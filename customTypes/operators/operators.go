package operators

type Operator string

const (
	Equal        Operator = "="
	NotEqual     Operator = "!="
	Greater      Operator = ">"
	GreaterEqual Operator = ">="
	Less         Operator = "<"
	LessEqual    Operator = "<="
)

var ValidOperators = map[Operator]struct{}{
	Equal:        {},
	NotEqual:     {},
	Greater:      {},
	GreaterEqual: {},
	Less:         {},
	LessEqual:    {},
}

func (o Operator) IsValid() bool {
	_, ok := ValidOperators[o]
	return ok
}

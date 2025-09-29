package sorting

type Type string

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

func (t Type) IsValid() bool {
	_, ok := ValidTypes[t]
	return ok
}

package rating

type Rating int

const (
	None Rating = iota
	Safe
	Questionable
	Explicit
)

func (r Rating) String() string {
	return []string{"", "safe", "questionable", "explicit"}[r]
}

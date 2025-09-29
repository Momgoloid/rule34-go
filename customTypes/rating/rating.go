package rating

type Rating string

const (
	Safe         Rating = "safe"
	Questionable Rating = "questionable"
	Explicit     Rating = "explicit"
)

var ValidRatings = map[Rating]struct{}{
	Safe:         {},
	Questionable: {},
	Explicit:     {},
}

func (r Rating) IsValid() bool {
	_, ok := ValidRatings[r]
	return ok
}

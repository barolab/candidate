package reddit

import (
	"unicode/utf8"

	"github.com/barolab/candidate"
)

const (
	minLength = 1
	maxLength = 15
)

// func init() {
// 	candidate.Register(New())
// }

// Redit is the social network provider that checks for a username validity / availability
type Redit struct{}

// New creates a new instance of a Redit NameProvider
func New() *Redit {
	return &Redit{}
}

func (r *Redit) String() string {
	return "Redit"
}

// Validate the username using Tiwtter rules
func (r *Redit) Validate(username string) (violations candidate.Violations) {
	length := utf8.RuneCountInString(username)

	if length > maxLength {
		violations = append(violations, candidate.NameTooLong)
	}

	if length < minLength {
		violations = append(violations, candidate.NameTooShort)
	}

	return violations
}

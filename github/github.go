package github

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

// Github is the social network provider that checks for a username validity / availability
type Github struct{}

// New creates a new instance of a Github NameProvider
func New() *Github {
	return &Github{}
}

func (g *Github) String() string {
	return "Github"
}

// Validate the username using Github rules
func (g *Github) Validate(username string) (violations candidate.Violations) {
	length := utf8.RuneCountInString(username)

	if length > maxLength {
		violations = append(violations, candidate.NameTooLong)
	}

	if length < minLength {
		violations = append(violations, candidate.NameTooShort)
	}

	return violations
}

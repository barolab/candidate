package twitter

import (
	"regexp"
	"unicode/utf8"

	"github.com/barolab/candidate"
)

const (
	minLength = 1
	maxLength = 15
)

var (
	illegalPatternRegexp = regexp.MustCompile("(?i)twitter")
	legalRunesRegexp     = regexp.MustCompile("^[0-9A-Za-z_]*$")
)

func init() {
	candidate.Register(New())
}

// Twitter is the social network provider that checks for a username validity / availability
type Twitter struct{}

// New creates a new instance of a Twitter NameProvider
func New() *Twitter {
	return &Twitter{}
}

func (t *Twitter) String() string {
	return "Twitter"
}

// Validate the username using Twitter rules
func (t *Twitter) Validate(username string) (violations candidate.Violations) {
	length := utf8.RuneCountInString(username)

	if length > maxLength {
		violations = append(violations, candidate.NameTooLong)
	}

	if length < minLength {
		violations = append(violations, candidate.NameTooShort)
	}

	if illegalPatternRegexp.MatchString(username) {
		violations = append(violations, candidate.NameContainsIllegalPattern)
	}

	if !legalRunesRegexp.MatchString(username) {
		violations = append(violations, candidate.NameContainsIllegalCharacters)
	}

	return violations
}

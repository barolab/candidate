package twitter

import (
	"regexp"
	"unicode/utf8"

	"github.com/barolab/candidate/lib"
)

const (
	minLength = 1
	maxLength = 15
)

var (
	illegalPatternRegexp = regexp.MustCompile("(?i)twitter")
	legalRunesRegexp     = regexp.MustCompile("^[0-9A-Za-z_]*$")
)

// Validate the username using Tiwtter rules
func (t *Twitter) Validate(username string) (violations lib.Violations) {
	length := utf8.RuneCountInString(username)

	if length > maxLength {
		violations = append(violations, lib.NameTooLong)
	}

	if length < minLength {
		violations = append(violations, lib.NameTooShort)
	}

	if illegalPatternRegexp.MatchString(username) {
		violations = append(violations, lib.NameContainsIllegalPattern)
	}

	if !legalRunesRegexp.MatchString(username) {
		violations = append(violations, lib.NameContainsIllegalCharacters)
	}

	return violations
}

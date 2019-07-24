package twitter

import (
	"fmt"
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

// Twitter is the social network provider that checks for a username validity / availability
type Twitter struct {
	url       string
	apiKey    string
	apiSecret string
}

// NewTwitter creates a new instance of a Twitter Social Network Provider
func NewTwitter(url, key, secret string) *Twitter {
	return &Twitter{
		url:       url,
		apiKey:    key,
		apiSecret: secret,
	}
}

// Name of the twitter social network provider
func (t *Twitter) Name() string {
	return "Twitter"
}

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

// DoesAlreadyExists check whenever the given username exist on twitter
func (t *Twitter) DoesAlreadyExists(username string) (bool, error) {
	return false, fmt.Errorf("Not implemented yet")
}

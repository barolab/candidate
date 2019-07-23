package social

import (
	"fmt"
	"regexp"
	"unicode/utf8"
)

const (
	minLength      = 1
	maxLength      = 15
	illegalPattern = "(?i)twitter"
	legalRunes     = "^[0-9A-Za-z_]+$"
)

var (
	illegalPatternRegexp = regexp.MustCompile(illegalPattern)
	legalRunesRegexp     = regexp.MustCompile(legalRunes)
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
func (t *Twitter) Validate(username string) error {
	if !isShortEnough(username) {
		return fmt.Errorf("%s is too long", username)
	}

	if !isLongEnough(username) {
		return fmt.Errorf("%s is too short", username)
	}

	if containsNoIllegalPattern(username) {
		return fmt.Errorf("%s contains illegal pattern", username)
	}

	if !onlyContainsLegalRunes(username) {
		return fmt.Errorf("%s contains illegal characters", username)
	}

	return nil
}

// DoesAlreadyExists check whenever the given username exist on twitter
func (t *Twitter) DoesAlreadyExists(username string) (bool, error) {
	return false, fmt.Errorf("Not implemented yet")
}

func isShortEnough(username string) bool {
	return utf8.RuneCountInString(username) < 15
}

func isLongEnough(username string) bool {
	return utf8.RuneCountInString(username) > 1
}

func containsNoIllegalPattern(username string) bool {
	return illegalPatternRegexp.MatchString(username)
}

func onlyContainsLegalRunes(username string) bool {
	return legalRunesRegexp.MatchString(username)
}

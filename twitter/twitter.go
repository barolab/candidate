package twitter

import (
	"fmt"
	"net/url"
	"regexp"
	"unicode/utf8"

	"github.com/barolab/candidate"
	"github.com/barolab/candidate/fetch"
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
type Twitter struct {
	client fetch.Fetcher
}

// New creates a new instance of a Twitter NameProvider
func New() *Twitter {
	return &Twitter{
		client: fetch.DefaultFetcher,
	}
}

func (t *Twitter) String() string {
	return "Twitter"
}

// IsAvailable check if the given name is available in Twitter
func (t *Twitter) IsAvailable(name string) (bool, error) {
	return fetch.IsNotFound(t.client, fmt.Sprintf("https://twitter.com/%s", url.QueryEscape(name)))
}

// WithFetcher can be set behavior of the Twitter HTTP request
func (t *Twitter) WithFetcher(f fetch.Fetcher) {
	t.client = f
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

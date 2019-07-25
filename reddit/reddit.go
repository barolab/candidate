package reddit

import (
	"fmt"
	"net/url"
	"regexp"
	"unicode/utf8"

	"github.com/barolab/candidate"
	"github.com/barolab/candidate/fetch"
)

const (
	minLength = 3
	maxLength = 20
)

var (
	legalRunesRegexp = regexp.MustCompile("^[0-9A-Za-z_]*$")
)

func init() {
	candidate.Register(New())
}

// Reddit is the social network provider that checks for a username validity / availability
type Reddit struct {
	client fetch.Fetcher
}

// New creates a new instance of a Reddit NameProvider
func New() *Reddit {
	return &Reddit{
		client: fetch.DefaultFetcher,
	}
}

func (r *Reddit) String() string {
	return "Reddit"
}

// WithFetcher can be set behavior of the Twitter HTTP request
func (r *Reddit) WithFetcher(f fetch.Fetcher) {
	r.client = f
}

// IsAvailable check if the given name is available in Reddit
func (r *Reddit) IsAvailable(name string) (bool, error) {
	return fetch.IsNotFound(r.client, fmt.Sprintf("https://www.reddit.com/user/%s", url.QueryEscape(name)))
}

// Validate the username using Tiwtter rules
func (r *Reddit) Validate(username string) (violations candidate.Violations) {
	length := utf8.RuneCountInString(username)

	if length > maxLength {
		violations = append(violations, candidate.NameTooLong)
	}

	if length < minLength {
		violations = append(violations, candidate.NameTooShort)
	}

	if !legalRunesRegexp.MatchString(username) {
		violations = append(violations, candidate.NameContainsIllegalCharacters)
	}

	return violations
}

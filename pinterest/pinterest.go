package pinterest

import (
	"fmt"
	"net/url"
	"unicode/utf8"

	"github.com/barolab/candidate"
	"github.com/barolab/candidate/fetch"
)

const (
	minLength = 1
	maxLength = 15
)

func init() {
	candidate.Register(New())
}

// Pinterest is the social network provider that checks for a username validity / availability
type Pinterest struct {
	client fetch.Fetcher
}

// New creates a new instance of a Pinterest NameProvider
func New() *Pinterest {
	return &Pinterest{
		client: fetch.DefaultFetcher,
	}
}

func (p *Pinterest) String() string {
	return "Pinterest"
}

// IsAvailable check if the given name is available in Pinterest
func (p *Pinterest) IsAvailable(name string) (bool, error) {
	return fetch.IsNotFound(p.client, fmt.Sprintf("https://www.pinterest.com/%s/", url.QueryEscape(name)))
}

// WithFetcher can be set behavior of the Pinterest HTTP request
func (p *Pinterest) WithFetcher(f fetch.Fetcher) {
	p.client = f
}

// Validate the username using Pinterest rules
func (p *Pinterest) Validate(username string) (violations candidate.Violations) {
	length := utf8.RuneCountInString(username)

	if length > maxLength {
		violations = append(violations, candidate.NameTooLong)
	}

	if length < minLength {
		violations = append(violations, candidate.NameTooShort)
	}

	return violations
}

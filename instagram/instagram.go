package instagram

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

// Instagram is the social network provider that checks for a username validity / availability
type Instagram struct {
	client fetch.Fetcher
}

// New creates a new instance of a Instagram NameProvider
func New() *Instagram {
	return &Instagram{
		client: fetch.DefaultFetcher,
	}
}

func (i *Instagram) String() string {
	return "Instagram"
}

// IsAvailable check if the given name is available in Instagram
func (i *Instagram) IsAvailable(name string) (bool, error) {
	return fetch.IsNotFound(i.client, fmt.Sprintf("https://www.instagram.com/%s/", url.QueryEscape(name)))
}

// WithFetcher can be set behavior of the Instagram HTTP request
func (i *Instagram) WithFetcher(f fetch.Fetcher) {
	i.client = f
}

// Validate the username using Instagram rules
func (i *Instagram) Validate(username string) (violations candidate.Violations) {
	length := utf8.RuneCountInString(username)

	if length > maxLength {
		violations = append(violations, candidate.NameTooLong)
	}

	if length < minLength {
		violations = append(violations, candidate.NameTooShort)
	}

	return violations
}

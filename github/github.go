package github

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

// Github is the social network provider that checks for a username validity / availability
type Github struct {
	client fetch.Fetcher
}

// New creates a new instance of a Github NameProvider
func New() *Github {
	return &Github{
		client: fetch.DefaultFetcher,
	}
}

func (g *Github) String() string {
	return "Github"
}

// WithFetcher can be set behavior of the Github HTTP request
func (g *Github) WithFetcher(f fetch.Fetcher) {
	g.client = f
}

// IsAvailable check if the given name is available in Github
func (g *Github) IsAvailable(name string) (bool, error) {
	return fetch.IsNotFound(g.client, fmt.Sprintf("https://github.com/%s", url.QueryEscape(name)))
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

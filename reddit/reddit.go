package reddit

import (
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

// Redit is the social network provider that checks for a username validity / availability
type Redit struct {
	client fetch.Fetcher
}

// New creates a new instance of a Redit NameProvider
func New() *Redit {
	return &Redit{
		client: fetch.DefaultFetcher,
	}
}

func (r *Redit) String() string {
	return "Redit"
}

// WithFetcher can be set behavior of the Twitter HTTP request
func (r *Redit) WithFetcher(f fetch.Fetcher) {
	r.client = f
}

// Validate the username using Tiwtter rules
func (r *Redit) Validate(username string) (violations candidate.Violations) {
	length := utf8.RuneCountInString(username)

	if length > maxLength {
		violations = append(violations, candidate.NameTooLong)
	}

	if length < minLength {
		violations = append(violations, candidate.NameTooShort)
	}

	return violations
}

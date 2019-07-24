package candidate

import (
	"github.com/barolab/candidate/fetch"
)

var providers []NameProvider

// NameProvider represents a service providing names, it then should be able to checks for existence of a given name
type NameProvider interface {
	String() string
	WithFetcher
	IsAvailabler
	Validator
}

// Validator interface for remote system that have special rules about user names
type Validator interface {
	// Validate if the given string respect the SocialNetwork restrictions
	Validate(name string) Violations
}

// IsAvailabler interface for remote system that can check for a username availability
type IsAvailabler interface {
	// IsAvailable check if the given name exists in the provider
	IsAvailable(name string) (bool, error)
}

// WithFetcher is an interface to set the behavior of all the name provider
type WithFetcher interface {
	// WithFetcher can bind a new HTTP Fetcher to a name provider to use different mechanism
	WithFetcher(f fetch.Fetcher)
}

// SocialNetworks returns a full list of NameProvider supported by this package
func SocialNetworks() []NameProvider {
	return providers
}

// UseFetcher will return all the supported providers but initialized with a custom http client
func UseFetcher(f fetch.Fetcher) {
	for _, p := range providers {
		p.WithFetcher(f)
	}
}

// Register a NameProvider in this package
func Register(np NameProvider) {
	providers = append(providers, np)
}

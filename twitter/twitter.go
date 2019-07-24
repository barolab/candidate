package twitter

import (
	"fmt"
)

const (
	name = "Twitter"
)

// Twitter is the social network provider that checks for a username validity / availability
type Twitter struct {
	url       string
	apiKey    string
	apiSecret string
}

// NewTwitter creates a new instance of a Twitter Social Network Provider
func NewTwitter(configurators ...configurator) *Twitter {
	t := &Twitter{}

	for _, c := range configurators {
		c(t)
	}

	return t
}

// Name of the twitter social network provider
func (t *Twitter) Name() string {
	return name
}

// URL to the twitter API
func (t *Twitter) URL() string {
	return t.url
}

// APIKey to contact twitter
func (t *Twitter) APIKey() string {
	return t.apiKey
}

// APISecret to contact twitter
func (t *Twitter) APISecret() string {
	return t.apiSecret
}

// DoesAlreadyExists check whenever the given username exist on twitter
func (t *Twitter) DoesAlreadyExists(username string) (bool, error) {
	return false, fmt.Errorf("Not implemented yet")
}

package twitter

import (
	"fmt"
	"net/http"
	"net/url"
)

// ParseError when making sure the url is secure
type ParseError struct {
	url    *url.URL
	err    string
	parent error
}

// NetworkError when contacting twitter
type NetworkError struct {
	url    *url.URL
	err    string
	parent error
}

func newParseError(u *url.URL, parent error) ParseError {
	return ParseError{
		url:    u,
		parent: parent,
	}
}

func newNetworkError(u *url.URL, parent error) NetworkError {
	return NetworkError{
		url:    u,
		parent: parent,
	}
}

func (e ParseError) Error() string {
	return fmt.Sprintf("Failed to parse url %s, cannot check for username availability", e.url)
}

// Unwrap the parent error
func (e *ParseError) Unwrap() error {
	return e.parent
}

func (e NetworkError) Error() string {
	return fmt.Sprintf("Failed to contact url %s, cannot check for username availability", e.url)
}

// Unwrap the parent error
func (e *NetworkError) Unwrap() error {
	return e.parent
}

// IsAvailable check if the given name is available in Twitter
func (t *Twitter) IsAvailable(name string) (bool, error) {
	url, err := url.Parse(fmt.Sprintf("https://twitter.com/%s", url.QueryEscape(name)))
	if err != nil {
		return false, newParseError(url, err)
	}

	res, err := http.Get(url.Host + url.Path)
	if err != nil {
		return false, newNetworkError(url, err)
	}

	defer res.Body.Close()
	return res.StatusCode == http.StatusNotFound, nil
}

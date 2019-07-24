package fetch

import (
	"net/http"
	"time"
)

// Fetcher is an interface helper that is used to execute an http request
type Fetcher interface {
	// Fetch perform an HTTP request and return the response, with error if any
	Fetch(req *http.Request) (*http.Response, error)
}

type nativeFetcher struct {
	client *http.Client
}

func (f *nativeFetcher) Fetch(req *http.Request) (*http.Response, error) {
	return f.client.Do(req)
}

// DefaultFetcher is the default mechanism for performing HTTP request in this module
var DefaultFetcher = &nativeFetcher{
	client: &http.Client{
		Timeout: time.Second * 5,
	},
}

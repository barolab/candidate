package twitter_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/barolab/candidate/twitter"
)

type IsAvailableTestCase struct {
	argument string
	expected bool
	err      error
	fetcher  *FakeFetcher
}

type FakeFetcher struct {
	res *http.Response
	err error
}

func (f *FakeFetcher) Fetch(req *http.Request) (*http.Response, error) {
	if f.err != nil {
		return nil, f.err
	}

	return f.res, nil
}

var (
	OkFetcher = &FakeFetcher{
		err: nil,
		res: &http.Response{
			Status:     "200 Ok",
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBufferString("Testing")),
		},
	}

	NotFoundFetcher = &FakeFetcher{
		err: nil,
		res: &http.Response{
			Status:     "404 Not Found",
			StatusCode: http.StatusNotFound,
			Body:       ioutil.NopCloser(bytes.NewBufferString("Testing")),
		},
	}

	ErrorFetcher = &FakeFetcher{
		err: fmt.Errorf("Failed to contact remote server"),
		res: nil,
	}
)

func TestIsAvailable(T *testing.T) {
	client := twitter.New()
	cases := []IsAvailableTestCase{
		{argument: "candidate", expected: false, err: nil, fetcher: OkFetcher},
		{argument: "candidate", expected: true, err: nil, fetcher: NotFoundFetcher},
		{argument: "candidate", expected: false, err: fmt.Errorf("Failed to contact %s at https://twitter.com/candidate with error %s", client, ErrorFetcher.err), fetcher: ErrorFetcher},
	}

	for _, c := range cases {
		client.WithFetcher(c.fetcher)
		ok, err := client.IsAvailable(c.argument)

		if ok != c.expected {
			T.Errorf("IsAvailable should have return %v, got %v for username %s", c.expected, ok, c.argument)
		}

		if err != nil && c.err == nil {
			T.Errorf("IsAvailable returned unexpected error %s for username %s", err, c.argument)
		}

		if err == nil && c.err != nil {
			T.Errorf("IsAvailable should have return error %s, got no error instead for username %s", c.err, c.argument)
		}

		if err != nil && c.err != nil && err.Error() != c.err.Error() {
			T.Errorf("IsAvailable should have return error %s, got %s for username %s", c.err, err, c.argument)
		}
	}
}

package pinterest_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/barolab/candidate"
	"github.com/barolab/candidate/pinterest"
)

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

type ValidationTestCase struct {
	username   string
	violations candidate.Violations
}

type IsAvailableTestCase struct {
	argument string
	expected bool
	err      error
	fetcher  *FakeFetcher
}

func TestValidate(T *testing.T) {
	client := pinterest.New()
	cases := []ValidationTestCase{
		{username: "", violations: candidate.Violations{candidate.NameTooShort}},
		{username: "this-string-has-too-much-runes", violations: candidate.Violations{candidate.NameTooLong}},
		{username: "This_is_Valid", violations: candidate.Violations{}},
		{username: "-illegal-rune", violations: candidate.Violations{}},
		{username: "000__000", violations: candidate.Violations{}},
	}

	for _, c := range cases {
		if violations := client.Validate(c.username); !violations.IsEqual(c.violations) {
			T.Errorf("%s failed to validate username %s, expected violations \n%s but got \n%s", client, c.username, c.violations, violations)
		}
	}
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
	client := pinterest.New()
	cases := []IsAvailableTestCase{
		{argument: "candidate", expected: false, err: nil, fetcher: OkFetcher},
		{argument: "candidate", expected: true, err: nil, fetcher: NotFoundFetcher},
		{argument: "candidate", expected: false, err: fmt.Errorf("Request to https://www.pinterest.com/candidate/ failed: %s", ErrorFetcher.err), fetcher: ErrorFetcher},
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

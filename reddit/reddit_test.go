package reddit_test

import (
	"testing"

	"github.com/barolab/candidate"
	"github.com/barolab/candidate/reddit"
)

type ValidationTestCase struct {
	username   string
	violations candidate.Violations
}

func TestValidate(T *testing.T) {
	client := reddit.New()
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

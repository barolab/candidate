package candidate_test

import (
	"testing"

	"github.com/barolab/candidate"
	"github.com/barolab/candidate/twitter"
)

type ValidationTestCase struct {
	username   string
	violations candidate.Violations
}

func TestTwitter(T *testing.T) {
	exec(twitter.NewTwitter(twitter.WithURL("FAKE")), []ValidationTestCase{
		{username: "", violations: candidate.Violations{candidate.NameTooShort}},
		{username: "This_is_Valid", violations: candidate.Violations{}},
		{username: "TWITTER", violations: candidate.Violations{candidate.NameContainsIllegalPattern}},
		{username: "this-string-has-too-much-runes", violations: candidate.Violations{candidate.NameTooLong, candidate.NameContainsIllegalCharacters}},
	}, T)
}

func exec(np candidate.NameProvider, cases []ValidationTestCase, T *testing.T) {
	for _, c := range cases {
		if violations := np.Validate(c.username); !violations.IsEqual(c.violations) {
			T.Errorf("%s failed to validate username %s, expected violations \n%s but got \n%s", np, c.username, c.violations, violations)
		}
	}
}

package twitter_test

import (
	"testing"

	"github.com/barolab/candidate"
	"github.com/barolab/candidate/twitter"
)

type ValidationTestCase struct {
	username   string
	violations candidate.Violations
}

func TestValidate(T *testing.T) {
	client := twitter.NewTwitter(
		twitter.WithURL("test"),
		twitter.WithAPIKey("test"),
		twitter.WithSecretKey("test"),
	)

	cases := []ValidationTestCase{
		{username: "", violations: candidate.Violations{candidate.NameTooShort}},
		{username: "this-string-has-too-much-runes", violations: candidate.Violations{candidate.NameTooLong, candidate.NameContainsIllegalCharacters}},
		{username: "TWITTER", violations: candidate.Violations{candidate.NameContainsIllegalPattern}},
		{username: "Twitter", violations: candidate.Violations{candidate.NameContainsIllegalPattern}},
		{username: "twitter", violations: candidate.Violations{candidate.NameContainsIllegalPattern}},
		{username: "twittEr", violations: candidate.Violations{candidate.NameContainsIllegalPattern}},
		{username: "TWITTTT", violations: candidate.Violations{}},
		{username: "This_is_Valid", violations: candidate.Violations{}},
		{username: "-illegal-rune", violations: candidate.Violations{candidate.NameContainsIllegalCharacters}},
		{username: "000__000", violations: candidate.Violations{}},
	}

	for _, c := range cases {
		if violations := client.Validate(c.username); !violations.IsEqual(c.violations) {
			T.Errorf("%s failed to validate username %s, expected violations \n%s but got \n%s", client, c.username, c.violations, violations)
		}
	}
}

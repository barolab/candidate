package twitter_test

import (
	"testing"

	"github.com/barolab/candidate/lib"
	"github.com/barolab/candidate/twitter"
)

var client = twitter.NewTwitter("test", "test", "test")

type ValidateTestCase struct {
	name       string
	violations lib.Violations
}

func TestValidate(T *testing.T) {
	cases := []ValidateTestCase{
		{name: "", violations: lib.Violations{lib.NameTooShort}},
		{name: "this-string-has-too-much-runes", violations: lib.Violations{lib.NameTooLong, lib.NameContainsIllegalCharacters}},
		{name: "TWITTER", violations: lib.Violations{lib.NameContainsIllegalPattern}},
		{name: "Twitter", violations: lib.Violations{lib.NameContainsIllegalPattern}},
		{name: "twitter", violations: lib.Violations{lib.NameContainsIllegalPattern}},
		{name: "twittEr", violations: lib.Violations{lib.NameContainsIllegalPattern}},
		{name: "TWITTTT", violations: lib.Violations{}},
		{name: "This_is_Valid", violations: lib.Violations{}},
		{name: "-illegal-rune", violations: lib.Violations{lib.NameContainsIllegalCharacters}},
		{name: "000__000", violations: lib.Violations{}},
	}

	for _, c := range cases {
		if violations := client.Validate(c.name); !violations.IsEqual(c.violations) {
			T.Errorf("Failed to validate username %s, expected violations \n%s but got \n%s", c.name, c.violations, violations)
		}
	}
}

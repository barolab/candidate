package twitter_test

import (
	"testing"

	"github.com/barolab/candidate/social"
	"github.com/barolab/candidate/twitter"
)

var client = twitter.NewTwitter("test", "test", "test")

type ValidateTestCase struct {
	name       string
	violations social.Violations
}

func TestValidate(T *testing.T) {
	cases := []ValidateTestCase{
		{name: "", violations: social.Violations{social.NameTooShort}},
		{name: "this-string-has-too-much-runes", violations: social.Violations{social.NameTooLong, social.NameContainsIllegalCharacters}},
		{name: "TWITTER", violations: social.Violations{social.NameContainsIllegalPattern}},
		{name: "Twitter", violations: social.Violations{social.NameContainsIllegalPattern}},
		{name: "twitter", violations: social.Violations{social.NameContainsIllegalPattern}},
		{name: "twittEr", violations: social.Violations{social.NameContainsIllegalPattern}},
		{name: "TWITTTT", violations: social.Violations{}},
		{name: "This_is_Valid", violations: social.Violations{}},
		{name: "-illegal-rune", violations: social.Violations{social.NameContainsIllegalCharacters}},
		{name: "000__000", violations: social.Violations{}},
	}

	for _, c := range cases {
		if violations := client.Validate(c.name); !violations.IsEqual(c.violations) {
			T.Errorf("Failed to validate username %s, expected violations \n%s but got \n%s", c.name, c.violations, violations)
		}
	}
}

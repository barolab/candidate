package candidate_test

import (
	"testing"

	"github.com/barolab/candidate"
	"github.com/barolab/candidate/twitter"
)

type Stringer interface {
	String() string
}

type supported []string
type ValidationTestCase struct {
	username   string
	violations candidate.Violations
}

func (s supported) contains(needle Stringer) bool {
	for _, sp := range s {
		if sp == needle.String() {
			return true
		}
	}

	return false
}

func TestAllProviders(T *testing.T) {
	providers := candidate.SocialNetworks()
	sp := supported{
		"Twitter",
		"Github",
		"Reddit",
	}

	for _, p := range providers {
		if !sp.contains(p) {
			T.Errorf("Provider %s is not in the SocialNetwork Result", p)
		}
	}
}

func TestTwitterProviders(T *testing.T) {
	exec(twitter.New(), []ValidationTestCase{
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

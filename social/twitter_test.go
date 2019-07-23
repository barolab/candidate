package social_test

import (
	"fmt"
	"testing"

	"github.com/barolab/candidate/social"
)

var twitter = social.NewTwitter("test", "test", "test")

type assert struct {
	arg      string
	expected error
}

func TestValidate(T *testing.T) {
	assertions := []assert{
		assert{arg: "", expected: fmt.Errorf(" is too short")},
		assert{arg: "this-string-has-too-much-runes", expected: fmt.Errorf("this-string-has-too-much-runes is too long")},
		assert{arg: "TWITTER", expected: fmt.Errorf("TWITTER contains illegal pattern")},
		assert{arg: "Twitter", expected: fmt.Errorf("Twitter contains illegal pattern")},
		assert{arg: "twitter", expected: fmt.Errorf("twitter contains illegal pattern")},
		assert{arg: "twittEr", expected: fmt.Errorf("twittEr contains illegal pattern")},
		assert{arg: "TWITTTT", expected: nil},
		assert{arg: "This_is_Valid", expected: nil},
		assert{arg: "-illegal-rune", expected: fmt.Errorf("-illegal-rune contains illegal characters")},
		assert{arg: "000__000", expected: nil},
	}

	for _, assertion := range assertions {
		err := twitter.Validate(assertion.arg)
		if assertion.expected == nil && err != nil {
			T.Errorf("Failed to validate %s, expected no errors, got %s", assertion.arg, err)
		} else if assertion.expected != nil && err == nil {
			T.Errorf("Failed to validate %s, expected %s, got no error", assertion.arg, assertion.expected)
		} else if err != nil && assertion.expected != nil && err.Error() != assertion.expected.Error() {
			T.Errorf("Failed to validate %s, expected %s, got %s", assertion.arg, assertion.expected, err)
		}
	}
}
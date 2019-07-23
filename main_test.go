package main

import "testing"

type assert struct {
	arg      string
	expected bool
}

func TestIsShortEnough(T *testing.T) {
	assertions := []assert{
		assert{arg: "", expected: true},
		assert{arg: "this-string-has-too-much-runes", expected: false},
	}

	for _, assertion := range assertions {
		if isShortEnough(assertion.arg) != assertion.expected {
			T.Errorf("Failed to validate that %s is short enough", assertion.arg)
		}
	}
}

func TestIsLongEnough(T *testing.T) {
	assertions := []assert{
		assert{arg: "", expected: false},
		assert{arg: "this-string-has-too-much-runes", expected: true},
	}

	for _, assertion := range assertions {
		if isLongEnough(assertion.arg) != assertion.expected {
			T.Errorf("Failed to validate that %s is long enough", assertion.arg)
		}
	}
}

func TestContainssIllegalPattern(T *testing.T) {
	assertions := []assert{
		assert{arg: "TWITTER", expected: true},
		assert{arg: "Twitter", expected: true},
		assert{arg: "twitter", expected: true},
		assert{arg: "twittEr", expected: true},
		assert{arg: "TWITTTT", expected: false},
	}

	for _, assertion := range assertions {
		if containsNoIllegalPattern(assertion.arg) != assertion.expected {
			T.Errorf("Failed to validate that %s does not contains illegal pattern", assertion.arg)
		}
	}
}

func TestContainsOnlyLegalRunes(t *testing.T) {
	assertions := []assert{
		assert{arg: "", expected: false},
		assert{arg: "This_is_Valid_Runes", expected: true},
		assert{arg: "ThisContains-illegal-rune", expected: false},
		assert{arg: "000__000", expected: true},
	}

	for _, assertion := range assertions {
		if onlyContainsLegalRunes(assertion.arg) != assertion.expected {
			t.Errorf("Failed to validate that %s does not contains illegal runes", assertion.arg)
		}
	}
}

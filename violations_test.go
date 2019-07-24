package candidate_test

import (
	"fmt"
	"testing"

	"github.com/barolab/candidate"
)

var expectedAllViolationsString = fmt.Sprintf(
	"- %s\n- %s\n- %s\n- %s\n- %s\n- \n",
	"NAME_TOO_LONG",
	"NAME_TOO_SHORT",
	"NAME_ALREADY_EXIST",
	"NAME_CONTAINS_ILLEGAL_PATTERNS",
	"NAME_CONTAINS_ILLEGAL_CHARACTERS",
)

func TestViolations(T *testing.T) {
	violations := candidate.Violations{
		candidate.NameTooLong,
		candidate.NameTooShort,
		candidate.NameAlreadyExist,
		candidate.NameContainsIllegalPattern,
		candidate.NameContainsIllegalCharacters,
		11, // unkown violation
	}

	if violations.String() != expectedAllViolationsString {
		T.Errorf("Failed to validate that all violations string match expected %s, got %s", expectedAllViolationsString, violations.String())
	}

	if violations.IsNil() {
		T.Errorf("Failed to validate that a full violations is not nil...")
	}

	if violations.IsEqual(candidate.Violations{}) {
		T.Errorf("Fail to validate that isEqual returns false for an empty violations and a full one")
	}

	if !violations.IsEqual(violations) {
		T.Errorf("Failed to validate that a full violations is equal with itsel")
	}
}

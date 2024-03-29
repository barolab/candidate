package candidate

import "fmt"

// Violation represent a username rule break
type Violation int

// Violations is a slice of violation
type Violations []Violation

const (
	// NameTooLong is the violation message for a name which does respect the max length requirement
	NameTooLong Violation = iota + 1

	// NameTooShort is the violation message for a name which does respect the min length requirement
	NameTooShort

	// NameContainsIllegalPattern is the violation message for a name that contains an illegal patter
	NameContainsIllegalPattern

	// NameContainsIllegalCharacters is the violation message for a name containing bad characters
	NameContainsIllegalCharacters
)

func (v Violation) String() string {
	switch v {
	case NameTooLong:
		return "NAME_TOO_LONG"
	case NameTooShort:
		return "NAME_TOO_SHORT"
	case NameContainsIllegalPattern:
		return "NAME_CONTAINS_ILLEGAL_PATTERNS"
	case NameContainsIllegalCharacters:
		return "NAME_CONTAINS_ILLEGAL_CHARACTERS"
	default:
		return ""
	}
}

func (v Violation) Error() string {
	return v.String()
}

func (violations Violations) String() string {
	var result = ""
	for _, v := range violations {
		result += fmt.Sprintf("- %s\n", v)
	}

	return result
}

// IsNil checks if the slice does not contains any violations
func (violations Violations) IsNil() bool {
	return len(violations) == 0
}

// IsEqual check if two violations slice are equals
func (violations Violations) IsEqual(vs Violations) bool {
	if len(violations) != len(vs) {
		return false
	}

	for i, v := range violations {
		if vs[i] != v {
			return false
		}
	}

	return true
}

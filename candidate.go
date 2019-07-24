package candidate

var providers = make([]NameProvider, 1)

// NameProvider represents a service providing names, it then should be able to checks for existence of a given name
type NameProvider interface {
	String() string
	IsAvailabler
	Validator
}

// Validator interface for remote system that have special rules about user names
type Validator interface {
	// Validate if the given string respect the SocialNetwork restrictions
	Validate(name string) Violations
}

// IsAvailabler interface for remote system that can check for a username availability
type IsAvailabler interface {
	// IsAvailable check if the given name exists in the provider
	IsAvailable(name string) (bool, error)
}

// SocialNetworks returns a full list of NameProvider supported by this package
func SocialNetworks() []NameProvider {
	return providers
}

// Register a NameProvider in this package
func Register(np NameProvider) {
	providers[0] = np
}

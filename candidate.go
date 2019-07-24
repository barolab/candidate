package candidate

var providers []NameProvider

// NameProvider represents a service providing names, it then should be able to checks for existence of a given name
type NameProvider interface {

	// Valide if the given string respect the SocialNetwork restrictions
	Validate(name string) Violations
}

// SocialNetworks returns a full list of NameProvider supported by this package
func SocialNetworks() []NameProvider {
	return providers
}

// Register a NameProvider in this package
func Register(np NameProvider) {
	providers = append(providers, np)
}

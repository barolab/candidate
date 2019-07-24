package lib

// SocialNetwork is an interface for any given social network that should check whenever a username
// match the requirements & does not already exists
type SocialNetwork interface {

	// Name of the social network
	Name() string

	// Valide if the given string respect the SocialNetwork restrictions
	Validate(username string) Violations

	// DoesAlreadyExists fetch the given username from the social network, return false if it exist, true otherwise
	DoesAlreadyExists(username string) (bool, error)
}

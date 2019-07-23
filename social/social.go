package social

// Network is an interface for any given social network that should check whenever a username
// match the requirements & does not already exists
type Network interface {

	// Name of the social network
	Name() string

	// Valide if the given string respect the SocialNetwork restrictions
	Validate(username string) error

	// DoesAlreadyExists fetch the given username from the social network, return false if it exist, true otherwise
	DoesAlreadyExists(username string) (bool, error)
}

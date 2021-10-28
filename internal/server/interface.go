package server

// ClientRegisterer is the interface implemented by types that register clients.
type ClientRegisterer interface {
	// Register registers a new client and returns the ID.
	Register() string
}

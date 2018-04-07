package users

// User represents a single user of a Resell project.
type User struct {
	// ID is a unique id of a user.
	ID string `json:"id"`

	// Name represents the name of a user.
	Name string `json:"name"`

	// Enabled shows if user is active or it was disabled.
	Enabled bool `json:"enabled"`
}

package roles

// Role represents a single Resell subnet.
type Role struct {
	// ProjectID represents an associated Resell project.
	ProjectID string `json:"project_id"`

	// UserID represents an associated Resell user.
	UserID string `json:"user_id"`
}

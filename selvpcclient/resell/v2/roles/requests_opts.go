package roles

// RoleOpt represents options for a single Resell role.
type RoleOpt struct {
	// ProjectID represents needed Resell project.
	ProjectID string `json:"project_id"`

	// UserID represents needed Resell user.
	UserID string `json:"user_id"`
}
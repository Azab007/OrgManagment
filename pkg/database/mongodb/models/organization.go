// models/organization.go

package models

// Member represents a member of the organization
type Member struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	AccessLevel string `json:"access_level"`
}

// Organization represents an organization entity
type Organization struct {
	ID          string   `json:"organization_id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Members     []Member `json:"organization_members"`
}

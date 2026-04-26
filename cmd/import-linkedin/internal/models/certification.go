// Package models defines the data structures for LinkedIn import entities.
package models

// Certification represents a professional certification imported from LinkedIn.
type Certification struct {
	Name           string `yaml:"name" json:"name"`
	Organization   string `yaml:"organization" json:"organization"`
	IssueDate      string `yaml:"issue_date" json:"issue_date"`
	ExpirationDate string `yaml:"expiration_date,omitempty" json:"expiration_date,omitempty"`
	CredentialID   string `yaml:"credential_id,omitempty" json:"credential_id,omitempty"`
	CredentialURL  string `yaml:"credential_url,omitempty" json:"credential_url,omitempty"`
}

// ID returns a unique identifier for this certification.
// Format: name#organization
func (c Certification) ID() string {
	return c.Name + "#" + c.Organization
}

// Validate checks if the certification has all required fields.
func (c Certification) Validate() error {
	if c.Name == "" {
		return NewValidationError("name", "cannot be empty")
	}
	if c.Organization == "" {
		return NewValidationError("organization", "cannot be empty")
	}
	if c.IssueDate == "" {
		return NewValidationError("issue_date", "cannot be empty")
	}
	return nil
}

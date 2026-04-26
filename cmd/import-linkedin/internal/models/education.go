// Package models defines the data structures for LinkedIn import entities.
package models

// Education represents an academic qualification imported from LinkedIn.
type Education struct {
	Institution string   `yaml:"institution" json:"institution"`
	Degree      string   `yaml:"degree" json:"degree"`
	Field       string   `yaml:"field,omitempty" json:"field,omitempty"`
	StartDate   string   `yaml:"start_date,omitempty" json:"start_date,omitempty"`
	EndDate     string   `yaml:"end_date,omitempty" json:"end_date,omitempty"`
	Description []string `yaml:"description,omitempty" json:"description,omitempty"`
}

// ID returns a unique identifier for this education entry.
// Format: institution#degree#field
func (e Education) ID() string {
	return e.Institution + "#" + e.Degree + "#" + e.Field
}

// Validate checks if the education has all required fields.
func (e Education) Validate() error {
	if e.Institution == "" {
		return NewValidationError("institution", "cannot be empty")
	}
	if e.Degree == "" {
		return NewValidationError("degree", "cannot be empty")
	}
	if e.StartDate == "" && e.EndDate == "" {
		return NewValidationError("dates", "at least one date (start or end) must be present")
	}
	return nil
}

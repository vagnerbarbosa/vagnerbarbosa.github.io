// Package models defines the data structures for LinkedIn import entities.
package models

// Experience represents a professional work experience imported from LinkedIn.
type Experience struct {
	Company     string   `yaml:"company" json:"company"`
	Title       string   `yaml:"title" json:"title"`
	StartDate   string   `yaml:"start_date" json:"start_date"`
	EndDate     string   `yaml:"end_date" json:"end_date"`
	Description []string `yaml:"description" json:"description"`
	TechStack   string   `yaml:"tech_stack,omitempty" json:"tech_stack,omitempty"`
	Location    string   `yaml:"location,omitempty" json:"location,omitempty"`
}

// ID returns a unique identifier for this experience.
// Format: company#title
func (e Experience) ID() string {
	return e.Company + "#" + e.Title
}

// Validate checks if the experience has all required fields.
func (e Experience) Validate() error {
	if e.Company == "" {
		return NewValidationError("company", "cannot be empty")
	}
	if e.Title == "" {
		return NewValidationError("title", "cannot be empty")
	}
	if e.StartDate == "" {
		return NewValidationError("start_date", "cannot be empty")
	}
	return nil
}

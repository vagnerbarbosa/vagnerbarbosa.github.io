// Package models defines the data structures for LinkedIn import entities.
package models

// Experience represents a professional work experience imported from LinkedIn.
type Experience struct {
	Company       string   `yaml:"company" json:"company"`
	Title         string   `yaml:"title" json:"title"`
	TitleEN       string   `yaml:"title_en,omitempty" json:"title_en,omitempty"`
	StartDate     string   `yaml:"start_date" json:"start_date"`
	StartDateEN   string   `yaml:"start_date_en,omitempty" json:"start_date_en,omitempty"`
	EndDate       string   `yaml:"end_date" json:"end_date"`
	EndDateEN     string   `yaml:"end_date_en,omitempty" json:"end_date_en,omitempty"`
	Description   []string `yaml:"description" json:"description"`
	DescriptionEN []string `yaml:"description_en,omitempty" json:"description_en,omitempty"`
	TechStack     string   `yaml:"tech_stack,omitempty" json:"tech_stack,omitempty"`
	Location      string   `yaml:"location,omitempty" json:"location,omitempty"`
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

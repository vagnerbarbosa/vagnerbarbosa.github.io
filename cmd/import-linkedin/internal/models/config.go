// Package models defines the data structures for LinkedIn import entities.
package models

// Social represents social media links.
type Social struct {
	LinkedIn string `yaml:"linkedin,omitempty" json:"linkedin,omitempty"`
	GitHub   string `yaml:"github,omitempty" json:"github,omitempty"`
	YouTube  string `yaml:"youtube,omitempty" json:"youtube,omitempty"`
}

// Title represents bilingual titles.
type Title struct {
	EN string `yaml:"en" json:"en"`
	PT string `yaml:"pt" json:"pt"`
}

// About represents bilingual about text.
type About struct {
	EN string `yaml:"en" json:"en"`
	PT string `yaml:"pt" json:"pt"`
}

// ConfigPortfolio represents the complete portfolio configuration structure.
// This maps to the config.yaml file structure.
type ConfigPortfolio struct {
	Title          Title           `yaml:"title" json:"title"`
	Name           string          `yaml:"name" json:"name"`
	About          About           `yaml:"about" json:"about"`
	Social         Social          `yaml:"social" json:"social"`
	Experiences    []Experience    `yaml:"experiences" json:"experiences"`
	Education      []Education     `yaml:"education" json:"education"`
	Certifications []Certification `yaml:"certifications" json:"certifications"`
}

// NewConfigPortfolio creates a new empty config.
func NewConfigPortfolio() *ConfigPortfolio {
	return &ConfigPortfolio{
		Experiences:    make([]Experience, 0),
		Education:      make([]Education, 0),
		Certifications: make([]Certification, 0),
	}
}

// GetExperienceMap returns experiences as a map keyed by ID for easy lookup.
func (c *ConfigPortfolio) GetExperienceMap() map[string]Experience {
	result := make(map[string]Experience, len(c.Experiences))
	for _, exp := range c.Experiences {
		result[exp.ID()] = exp
	}
	return result
}

// GetEducationMap returns education as a map keyed by ID for easy lookup.
func (c *ConfigPortfolio) GetEducationMap() map[string]Education {
	result := make(map[string]Education, len(c.Education))
	for _, edu := range c.Education {
		result[edu.ID()] = edu
	}
	return result
}

// GetCertificationMap returns certifications as a map keyed by ID for easy lookup.
func (c *ConfigPortfolio) GetCertificationMap() map[string]Certification {
	result := make(map[string]Certification, len(c.Certifications))
	for _, cert := range c.Certifications {
		result[cert.ID()] = cert
	}
	return result
}

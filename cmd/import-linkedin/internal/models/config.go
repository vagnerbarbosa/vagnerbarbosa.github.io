// Package models defines the data structures for LinkedIn import entities.
package models

// Social represents social media links.
type Social struct {
	LinkedIn string `yaml:"linkedin,omitempty" json:"linkedin,omitempty"`
	GitHub   string `yaml:"github,omitempty" json:"github,omitempty"`
	YouTube  string `yaml:"youtube,omitempty" json:"youtube,omitempty"`
}

// About represents bilingual about text.
type About struct {
	EN string `yaml:"en" json:"en"`
	PT string `yaml:"pt" json:"pt"`
}

// SiteConfig represents the site configuration part of the YAML.
type SiteConfig struct {
	Title           string  `yaml:"title"`
	Description     string  `yaml:"description"`
	BaseURL         string  `yaml:"baseurl"`
	Username        string  `yaml:"username"`
	UserDescription string  `yaml:"user_description"`
	UserTitle       string  `yaml:"user_title"`
	Email           string  `yaml:"email"`
	Social          Social  `yaml:"social"`
}

// Content represents the content part of the YAML.
type Content struct {
	About        About          `yaml:"about"`
	Experiences  []Experience   `yaml:"experiences"`
	Education    []Education    `yaml:"education"`
	Certifications []Certification `yaml:"certifications"`
	Technologies []string      `yaml:"technologies"`
}

// ConfigPortfolio represents the complete portfolio configuration structure.
// This matches the structure expected by the site generator.
type ConfigPortfolio struct {
	Site    SiteConfig `yaml:"site"`
	Content Content    `yaml:"content"`
}

// NewConfigPortfolio creates a new empty config.
func NewConfigPortfolio() *ConfigPortfolio {
	return &ConfigPortfolio{
		Site: SiteConfig{},
		Content: Content{
			Experiences:    make([]Experience, 0),
			Education:      make([]Education, 0),
			Certifications: make([]Certification, 0),
		},
	}
}

// GetExperienceMap returns experiences as a map keyed by ID for easy lookup.
func (c *ConfigPortfolio) GetExperienceMap() map[string]Experience {
	result := make(map[string]Experience, len(c.Content.Experiences))
	for _, exp := range c.Content.Experiences {
		result[exp.ID()] = exp
	}
	return result
}

// GetEducationMap returns education as a map keyed by ID for easy lookup.
func (c *ConfigPortfolio) GetEducationMap() map[string]Education {
	result := make(map[string]Education, len(c.Content.Education))
	for _, edu := range c.Content.Education {
		result[edu.ID()] = edu
	}
	return result
}

// GetCertificationMap returns certifications as a map keyed by ID for easy lookup.
func (c *ConfigPortfolio) GetCertificationMap() map[string]Certification {
	result := make(map[string]Certification, len(c.Content.Certifications))
	for _, cert := range c.Content.Certifications {
		result[cert.ID()] = cert
	}
	return result
}

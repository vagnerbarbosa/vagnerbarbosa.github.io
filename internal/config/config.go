// Package config handles site configuration loading
package config

import (
	"fmt"
	"os"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

// SiteConfig represents the site configuration
type SiteConfig struct {
	Title           string `yaml:"title"`
	Description     string `yaml:"description"`
	BaseURL         string `yaml:"baseurl"`
	Username        string `yaml:"username"`
	UserDescription string `yaml:"user_description"`
	UserTitle       string `yaml:"user_title"`
	Email           string `yaml:"email"`
}

type Experience struct {
	Title       string   `yaml:"title"`
	TitleEN     string   `yaml:"title_en,omitempty"`
	Company     string   `yaml:"company"`
	StartDate   string   `yaml:"start_date"`
	StartDateEN string   `yaml:"start_date_en,omitempty"`
	EndDate     string   `yaml:"end_date"`
	EndDateEN   string   `yaml:"end_date_en,omitempty"`
	Description []string `yaml:"description"`
	DescriptionEN []string `yaml:"description_en,omitempty"`
	TechStack   string   `yaml:"tech_stack"`
	Location    string   `yaml:"location,omitempty"`
}

type Education struct {
	Institution string   `yaml:"institution"`
	Degree      string   `yaml:"degree"`
	DegreeEN    string   `yaml:"degree_en,omitempty"`
	Field       string   `yaml:"field,omitempty"`
	StartDate   string   `yaml:"start_date,omitempty"`
	StartDateEN string   `yaml:"start_date_en,omitempty"`
	EndDate     string   `yaml:"end_date,omitempty"`
	EndDateEN   string   `yaml:"end_date_en,omitempty"`
	Description []string `yaml:"description,omitempty"`
}

type Certification struct {
	Name            string `yaml:"name"`
	Organization    string `yaml:"organization"`
	IssueDate       string `yaml:"issue_date"`
	ExpirationDate  string `yaml:"expiration_date,omitempty"`
	CredentialID    string `yaml:"credential_id,omitempty"`
	CredentialURL   string `yaml:"credential_url,omitempty"`
}

// AboutContent holds the about section content
type AboutContent struct {
	PT string `yaml:"pt"`
	EN string `yaml:"en"`
}

// ParsedAbout holds parsed paragraphs
type ParsedAbout struct {
	ParagraphsPT []string
	ParagraphsEN []string
}

// Content holds all dynamic content
type Content struct {
	About        AboutContent `yaml:"about"`
	ParsedAbout  ParsedAbout
	Experiences  []Experience `yaml:"experiences"`
	Education    []Education  `yaml:"education"`
	Certifications []Certification `yaml:"certifications"`
	Technologies []string     `yaml:"technologies"`
	Privacy      AboutContent `yaml:"privacy"`
}

// Config is the main configuration structure
type Config struct {
	Site    SiteConfig `yaml:"site"`
	Content Content    `yaml:"content"`
}

// TemplateData holds data for template rendering
type TemplateData struct {
	Site    SiteConfig
	Content Content
	Year    int
}

// Load reads configuration from a YAML file
func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	// Parse about paragraphs
	cfg.Content.ParsedAbout.ParagraphsPT = splitParagraphs(cfg.Content.About.PT)
	cfg.Content.ParsedAbout.ParagraphsEN = splitParagraphs(cfg.Content.About.EN)

	return &cfg, nil
}

// splitParagraphs splits text into paragraphs
func splitParagraphs(text string) []string {
	lines := strings.Split(strings.TrimSpace(text), "\n")
	var paragraphs []string
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed != "" {
			paragraphs = append(paragraphs, trimmed)
		}
	}
	return paragraphs
}

// CurrentYear returns the current year for copyright
func CurrentYear() int {
	return time.Now().Year()
}

// ToTemplateData converts Config to TemplateData
func (c *Config) ToTemplateData() TemplateData {
	// Copy parsed about to content for template access
	c.Content.ParsedAbout = ParsedAbout{
		ParagraphsPT: splitParagraphs(c.Content.About.PT),
		ParagraphsEN: splitParagraphs(c.Content.About.EN),
	}

	return TemplateData{
		Site:    c.Site,
		Content: c.Content,
		Year:    CurrentYear(),
	}
}

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

// Experience represents a work experience entry
type Experience struct {
	Title      string   `yaml:"title"`
	TitleEN    string   `yaml:"title_en"`
	Company    string   `yaml:"company"`
	Period     string   `yaml:"period"`
	PeriodEN   string   `yaml:"period_en"`
	Details    []string `yaml:"details"`
	DetailsEN  []string `yaml:"details_en"`
	TechStack  string   `yaml:"tech_stack"`
}

// Education represents an education entry
type Education struct {
	Title    string `yaml:"title"`
	TitleEN  string `yaml:"title_en"`
	School   string `yaml:"school"`
	Period   string `yaml:"period"`
	PeriodEN string `yaml:"period_en"`
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
	Technologies []string     `yaml:"technologies"`
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

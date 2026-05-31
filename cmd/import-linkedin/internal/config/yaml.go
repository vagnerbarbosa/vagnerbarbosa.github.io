// Package config handles YAML configuration file operations.
package config

import (
	"fmt"
	"os"

	"github.com/vagnerbarbosa/vagnerbarbosa.github.io/cmd/import-linkedin/internal/models"
	"gopkg.in/yaml.v3"
)

// ReadYAML reads the config.yaml file and returns the parsed configuration.
func ReadYAML(filepath string) (*models.ConfigPortfolio, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config models.ConfigPortfolio
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	return &config, nil
}

// WriteYAML writes the configuration to a YAML file.
func WriteYAML(filepath string, config *models.ConfigPortfolio) error {
	data, err := yaml.Marshal(config)
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	if err := os.WriteFile(filepath, data, 0644); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	return nil
}

// YAMLManager handles YAML file operations with node preservation.
type YAMLManager struct {
	filepath string
	root     *yaml.Node
	config   *models.ConfigPortfolio
}

// NewYAMLManager creates a new YAML manager for the given file.
func NewYAMLManager(filepath string) (*YAMLManager, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var root yaml.Node
	if err := yaml.Unmarshal(data, &root); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	var config models.ConfigPortfolio
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse config into struct: %w", err)
	}

	return &YAMLManager{
		filepath: filepath,
		root:     &root,
		config:   &config,
	}, nil
}

// GetConfig returns the parsed configuration.
func (m *YAMLManager) GetConfig() *models.ConfigPortfolio {
	return m.config
}

// UpdateExperiences updates the experiences section.
func (m *YAMLManager) UpdateExperiences(experiences []models.Experience) error {
	m.config.Content.Experiences = experiences
	return nil
}

// UpdateEducation updates the education section.
func (m *YAMLManager) UpdateEducation(education []models.Education) error {
	m.config.Content.Education = education
	return nil
}

// UpdateCertifications updates the certifications section.
func (m *YAMLManager) UpdateCertifications(certifications []models.Certification) error {
	m.config.Content.Certifications = certifications
	return nil
}

// Save writes the updated configuration back to the file.
func (m *YAMLManager) Save() error {
	return WriteYAML(m.filepath, m.config)
}

// FileExists checks if a file exists.
func FileExists(filepath string) bool {
	_, err := os.Stat(filepath)
	return !os.IsNotExist(err)
}

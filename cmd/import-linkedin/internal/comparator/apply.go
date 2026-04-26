// Package comparator handles comparison between LinkedIn data and current config.
package comparator

import (
	"fmt"

	"github.com/vagnerbarbosa/vagnerbarbosa.github.io/cmd/import-linkedin/internal/models"
)

// ApplyChanges applies all detected changes to the config.
func ApplyChanges(diff *Diff, config *models.ConfigPortfolio) error {
	if diff == nil || config == nil {
		return fmt.Errorf("diff and config cannot be nil")
	}

	// Apply experience changes
	if err := applyExperienceChanges(diff.Experiences, config); err != nil {
		return fmt.Errorf("failed to apply experience changes: %w", err)
	}

	// Apply education changes
	if err := applyEducationChanges(diff.Education, config); err != nil {
		return fmt.Errorf("failed to apply education changes: %w", err)
	}

	// Apply certification changes
	if err := applyCertificationChanges(diff.Certifications, config); err != nil {
		return fmt.Errorf("failed to apply certification changes: %w", err)
	}

	return nil
}

// applyExperienceChanges applies experience changes to the config.
func applyExperienceChanges(diff EntityDiff[models.Experience], config *models.ConfigPortfolio) error {
	// Create a map of current experiences
	currentMap := config.GetExperienceMap()

	// Add new experiences
	for _, exp := range diff.Added {
		currentMap[exp.ID()] = exp
	}

	// Update modified experiences
	for _, pair := range diff.Modified {
		currentMap[pair.New.ID()] = pair.New
	}

	// Remove deleted experiences
	for _, exp := range diff.Removed {
		delete(currentMap, exp.ID())
	}

	// Convert map back to slice
	config.Experiences = make([]models.Experience, 0, len(currentMap))
	for _, exp := range currentMap {
		config.Experiences = append(config.Experiences, exp)
	}

	return nil
}

// applyEducationChanges applies education changes to the config.
func applyEducationChanges(diff EntityDiff[models.Education], config *models.ConfigPortfolio) error {
	// Create a map of current education
	currentMap := config.GetEducationMap()

	// Add new education
	for _, edu := range diff.Added {
		currentMap[edu.ID()] = edu
	}

	// Update modified education
	for _, pair := range diff.Modified {
		currentMap[pair.New.ID()] = pair.New
	}

	// Remove deleted education
	for _, edu := range diff.Removed {
		delete(currentMap, edu.ID())
	}

	// Convert map back to slice
	config.Education = make([]models.Education, 0, len(currentMap))
	for _, edu := range currentMap {
		config.Education = append(config.Education, edu)
	}

	return nil
}

// applyCertificationChanges applies certification changes to the config.
func applyCertificationChanges(diff EntityDiff[models.Certification], config *models.ConfigPortfolio) error {
	// Create a map of current certifications
	currentMap := config.GetCertificationMap()

	// Add new certifications
	for _, cert := range diff.Added {
		currentMap[cert.ID()] = cert
	}

	// Update modified certifications
	for _, pair := range diff.Modified {
		currentMap[pair.New.ID()] = pair.New
	}

	// Remove deleted certifications
	for _, cert := range diff.Removed {
		delete(currentMap, cert.ID())
	}

	// Convert map back to slice
	config.Certifications = make([]models.Certification, 0, len(currentMap))
	for _, cert := range currentMap {
		config.Certifications = append(config.Certifications, cert)
	}

	return nil
}

// MergeConfigs merges LinkedIn data with existing config.
// Returns a new config with all LinkedIn data, preserving other sections.
func MergeConfigs(linkedinData *models.ConfigPortfolio, existingConfig *models.ConfigPortfolio) *models.ConfigPortfolio {
	result := *existingConfig

	// Override with LinkedIn data
	if len(linkedinData.Experiences) > 0 {
		result.Experiences = linkedinData.Experiences
	}
	if len(linkedinData.Education) > 0 {
		result.Education = linkedinData.Education
	}
	if len(linkedinData.Certifications) > 0 {
		result.Certifications = linkedinData.Certifications
	}

	return &result
}

// ImportResult represents the result of an import operation.
type ImportResult struct {
	ExperiencesAdded    int
	ExperiencesModified int
	ExperiencesRemoved  int
	EducationAdded      int
	EducationModified   int
	EducationRemoved    int
	CertificationsAdded int
	CertificationsModified int
	CertificationsRemoved  int
}

// GetImportResult returns a summary of the import operation.
func GetImportResult(diff *Diff) ImportResult {
	return ImportResult{
		ExperiencesAdded:       len(diff.Experiences.Added),
		ExperiencesModified:    len(diff.Experiences.Modified),
		ExperiencesRemoved:     len(diff.Experiences.Removed),
		EducationAdded:         len(diff.Education.Added),
		EducationModified:      len(diff.Education.Modified),
		EducationRemoved:       len(diff.Education.Removed),
		CertificationsAdded:    len(diff.Certifications.Added),
		CertificationsModified: len(diff.Certifications.Modified),
		CertificationsRemoved:  len(diff.Certifications.Removed),
	}
}

// PrintImportResult prints a summary of the import operation.
func PrintImportResult(result ImportResult) {
	totalAdded := result.ExperiencesAdded + result.EducationAdded + result.CertificationsAdded
	totalModified := result.ExperiencesModified + result.EducationModified + result.CertificationsModified
	totalRemoved := result.ExperiencesRemoved + result.EducationRemoved + result.CertificationsRemoved

	fmt.Printf("\n✓ Importação concluída:\n")
	fmt.Printf("  %d adições, %d modificações, %d remoções\n",
		totalAdded, totalModified, totalRemoved)

	if result.ExperiencesAdded > 0 || result.ExperiencesModified > 0 || result.ExperiencesRemoved > 0 {
		fmt.Printf("  Experiências: %d adições, %d modificações, %d remoções\n",
			result.ExperiencesAdded, result.ExperiencesModified, result.ExperiencesRemoved)
	}

	if result.EducationAdded > 0 || result.EducationModified > 0 || result.EducationRemoved > 0 {
		fmt.Printf("  Educação: %d adições, %d modificações, %d remoções\n",
			result.EducationAdded, result.EducationModified, result.EducationRemoved)
	}

	if result.CertificationsAdded > 0 || result.CertificationsModified > 0 || result.CertificationsRemoved > 0 {
		fmt.Printf("  Certificações: %d adições, %d modificações, %d remoções\n",
			result.CertificationsAdded, result.CertificationsModified, result.CertificationsRemoved)
	}
}

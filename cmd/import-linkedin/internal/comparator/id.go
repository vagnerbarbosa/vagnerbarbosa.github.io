// Package comparator handles comparison between LinkedIn data and current config.
package comparator

import "github.com/vagnerbarbosa/vagnerbarbosa.github.io/cmd/import-linkedin/internal/models"

// GenerateExperienceID creates a unique identifier for an experience.
// Format: company#role
func GenerateExperienceID(exp models.Experience) string {
	return exp.ID()
}

// GenerateEducationID creates a unique identifier for an education entry.
// Format: institution#degree#field
func GenerateEducationID(edu models.Education) string {
	return edu.ID()
}

// GenerateCertificationID creates a unique identifier for a certification.
// Format: name#organization
func GenerateCertificationID(cert models.Certification) string {
	return cert.ID()
}

// IDGenerator provides a unified interface for generating entity IDs.
type IDGenerator struct{}

// NewIDGenerator creates a new ID generator.
func NewIDGenerator() *IDGenerator {
	return &IDGenerator{}
}

// Generate creates an ID based on entity type.
func (g *IDGenerator) Generate(entity any) string {
	switch e := entity.(type) {
	case models.Experience:
		return GenerateExperienceID(e)
	case *models.Experience:
		if e == nil {
			return ""
		}
		return GenerateExperienceID(*e)
	case models.Education:
		return GenerateEducationID(e)
	case *models.Education:
		if e == nil {
			return ""
		}
		return GenerateEducationID(*e)
	case models.Certification:
		return GenerateCertificationID(e)
	case *models.Certification:
		if e == nil {
			return ""
		}
		return GenerateCertificationID(*e)
	default:
		return ""
	}
}

// GetEntityType returns the entity type name.
func GetEntityType(entity any) string {
	switch entity.(type) {
	case models.Experience, *models.Experience:
		return "experience"
	case models.Education, *models.Education:
		return "education"
	case models.Certification, *models.Certification:
		return "certification"
	default:
		return "unknown"
	}
}

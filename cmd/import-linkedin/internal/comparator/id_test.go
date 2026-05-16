package comparator

import (
	"testing"

	"github.com/vagnerbarbosa/vagnerbarbosa.github.io/cmd/import-linkedin/internal/models"
)

func TestIDGenerator(t *testing.T) {
	gen := NewIDGenerator()

	tests := []struct {
		name     string
		entity   any
		expected string
	}{
		{
			name:     "experience",
			entity:   models.Experience{Company: "Google", Role: "Engineer"},
			expected: "Google#Engineer",
		},
		{
			name:     "experience pointer",
			entity:   &models.Experience{Company: "Google", Role: "Engineer"},
			expected: "Google#Engineer",
		},
		{
			name:     "experience pointer nil",
			entity:   (*models.Experience)(nil),
			expected: "",
		},
		{
			name:     "education",
			entity:   models.Education{Institution: "MIT", Degree: "BS", Field: "CS"},
			expected: "MIT#BS#CS",
		},
		{
			name:     "education pointer",
			entity:   &models.Education{Institution: "MIT", Degree: "BS", Field: "CS"},
			expected: "MIT#BS#CS",
		},
		{
			name:     "education pointer nil",
			entity:   (*models.Education)(nil),
			expected: "",
		},
		{
			name:     "certification",
			entity:   models.Certification{Name: "AWS", Organization: "Amazon"},
			expected: "AWS#Amazon",
		},
		{
			name:     "certification pointer",
			entity:   &models.Certification{Name: "AWS", Organization: "Amazon"},
			expected: "AWS#Amazon",
		},
		{
			name:     "certification pointer nil",
			entity:   (*models.Certification)(nil),
			expected: "",
		},
		{
			name:     "unknown entity",
			entity:   "some random string",
			expected: "",
		},
		{
			name:     "nil entity",
			entity:   nil,
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gen.Generate(tt.entity); got != tt.expected {
				t.Errorf("Generate() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestGetEntityType(t *testing.T) {
	tests := []struct {
		name     string
		entity   any
		expected string
	}{
		{
			name:     "experience",
			entity:   models.Experience{},
			expected: "experience",
		},
		{
			name:     "experience pointer",
			entity:   &models.Experience{},
			expected: "experience",
		},
		{
			name:     "education",
			entity:   models.Education{},
			expected: "education",
		},
		{
			name:     "education pointer",
			entity:   &models.Education{},
			expected: "education",
		},
		{
			name:     "certification",
			entity:   models.Certification{},
			expected: "certification",
		},
		{
			name:     "certification pointer",
			entity:   &models.Certification{},
			expected: "certification",
		},
		{
			name:     "unknown",
			entity:   "random",
			expected: "unknown",
		},
		{
			name:     "nil",
			entity:   nil,
			expected: "unknown",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetEntityType(tt.entity); got != tt.expected {
				t.Errorf("GetEntityType() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestGenerateIDs(t *testing.T) {
	exp := models.Experience{Company: "Google", Role: "Engineer"}
	if GenerateExperienceID(exp) != "Google#Engineer" {
		t.Error("GenerateExperienceID failed")
	}

	edu := models.Education{Institution: "MIT", Degree: "BS", Field: "CS"}
	if GenerateEducationID(edu) != "MIT#BS#CS" {
		t.Error("GenerateEducationID failed")
	}

	cert := models.Certification{Name: "AWS", Organization: "Amazon"}
	if GenerateCertificationID(cert) != "AWS#Amazon" {
		t.Error("GenerateCertificationID failed")
	}
}

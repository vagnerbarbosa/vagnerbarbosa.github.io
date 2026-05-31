package comparator

import (
	"testing"

	"github.com/vagnerbarbosa/vagnerbarbosa.github.io/cmd/import-linkedin/internal/models"
)

func TestApplyChanges(t *testing.T) {
	config := &models.ConfigPortfolio{
		Content: models.Content{
			Experiences: []models.Experience{
				{Company: "OldCo", Title: "Dev", StartDate: "Jan 2020"},
			},
			Education: []models.Education{
				{Institution: "OldUni", Degree: "BS", StartDate: "Jan 2010"},
			},
			Certifications: []models.Certification{
				{Name: "OldCert", Organization: "OldOrg", IssueDate: "Jan 2015"},
			},
		},
	}

	diff := NewDiff()
	diff.Experiences.Added = []models.Experience{
		{Company: "NewCo", Title: "Dev", StartDate: "Jan 2021"},
	}
	diff.Education.Modified = []ChangePair[models.Education]{
		{
			Old: models.Education{Institution: "OldUni", Degree: "BS", StartDate: "Jan 2010"},
			New: models.Education{Institution: "OldUni", Degree: "BS", StartDate: "Jan 2011"},
		},
	}
	diff.Certifications.Removed = []models.Certification{
		{Name: "OldCert", Organization: "OldOrg", IssueDate: "Jan 2015"},
	}

	err := ApplyChanges(diff, config)
	if err != nil {
		t.Fatalf("ApplyChanges failed: %v", err)
	}

	// Verify Experiences
	if len(config.Content.Experiences) != 2 {
		t.Errorf("Expected 2 experiences, got %d", len(config.Content.Experiences))
	}

	// Verify Education
	if len(config.Content.Education) != 1 || config.Content.Education[0].StartDate != "Jan 2011" {
		t.Errorf("Expected education to be updated to Jan 2011, got %v", config.Content.Education)
	}

	// Verify Certifications
	if len(config.Content.Certifications) != 0 {
		t.Errorf("Expected 0 certifications, got %d", len(config.Content.Certifications))
	}
}

func TestApplyChanges_NilPointers(t *testing.T) {
	config := &models.ConfigPortfolio{}
	diff := NewDiff()

	if err := ApplyChanges(nil, config); err == nil {
		t.Error("Expected error when diff is nil")
	}
	if err := ApplyChanges(diff, nil); err == nil {
		t.Error("Expected error when config is nil")
	}
}

func TestApplyChanges_Comprehensive(t *testing.T) {
	config := &models.ConfigPortfolio{
		Content: models.Content{
			Experiences: []models.Experience{
				{Company: "Exp1", Title: "R1", StartDate: "S1"},
				{Company: "Exp2", Title: "R2", StartDate: "S2"},
				{Company: "Exp3", Title: "R3", StartDate: "S3"},
			},
			Education: []models.Education{
				{Institution: "Edu1", Degree: "D1", StartDate: "S1"},
				{Institution: "Edu2", Degree: "D2", StartDate: "S2"},
				{Institution: "Edu3", Degree: "D3", StartDate: "S3"},
			},
			Certifications: []models.Certification{
				{Name: "Cert1", Organization: "O1", IssueDate: "I1"},
				{Name: "Cert2", Organization: "O2", IssueDate: "I2"},
				{Name: "Cert3", Organization: "O3", IssueDate: "I3"},
			},
		},
	}

	diff := NewDiff()

	// Experiences
	diff.Experiences.Added = []models.Experience{{Company: "ExpAdd", Title: "RAdd", StartDate: "SAdd"}}
	diff.Experiences.Modified = []ChangePair[models.Experience]{
		{
			Old: models.Experience{Company: "Exp1", Title: "R1", StartDate: "S1"},
			New: models.Experience{Company: "Exp1", Title: "R1", StartDate: "S1-mod"},
		},
	}
	diff.Experiences.Removed = []models.Experience{{Company: "Exp2", Title: "R2", StartDate: "S2"}}

	// Education
	diff.Education.Added = []models.Education{{Institution: "EduAdd", Degree: "DAdd", StartDate: "SAdd"}}
	diff.Education.Modified = []ChangePair[models.Education]{
		{
			Old: models.Education{Institution: "Edu1", Degree: "D1", StartDate: "S1"},
			New: models.Education{Institution: "Edu1", Degree: "D1", StartDate: "S1-mod"},
		},
	}
	diff.Education.Removed = []models.Education{{Institution: "Edu2", Degree: "D2", StartDate: "S2"}}

	// Certifications
	diff.Certifications.Added = []models.Certification{{Name: "CertAdd", Organization: "OAdd", IssueDate: "IAdd"}}
	diff.Certifications.Modified = []ChangePair[models.Certification]{
		{
			Old: models.Certification{Name: "Cert1", Organization: "O1", IssueDate: "I1"},
			New: models.Certification{Name: "Cert1", Organization: "O1", IssueDate: "I1-mod"},
		},
	}
	diff.Certifications.Removed = []models.Certification{{Name: "Cert2", Organization: "O2", IssueDate: "I2"}}

	err := ApplyChanges(diff, config)
	if err != nil {
		t.Fatalf("ApplyChanges failed: %v", err)
	}

	if len(config.Content.Experiences) != 3 { // 3 - 1 (removed) + 1 (added) = 3
		t.Errorf("Expected 3 experiences, got %d", len(config.Content.Experiences))
	}
	if len(config.Content.Education) != 3 {
		t.Errorf("Expected 3 educations, got %d", len(config.Content.Education))
	}
	if len(config.Content.Certifications) != 3 {
		t.Errorf("Expected 3 certifications, got %d", len(config.Content.Certifications))
	}
}

func TestMergeConfigs(t *testing.T) {
	tests := []struct {
		name     string
		existing *models.ConfigPortfolio
		linkedin *models.ConfigPortfolio
		verify   func(*models.ConfigPortfolio) bool
	}{
		{
			name: "all overridden",
			existing: &models.ConfigPortfolio{
				Content: models.Content{
					Experiences:    []models.Experience{{Company: "OldCo"}},
					Education:      []models.Education{{Institution: "OldUni"}},
					Certifications: []models.Certification{{Name: "OldCert"}},
				},
			},
			linkedin: &models.ConfigPortfolio{
				Content: models.Content{
					Experiences:    []models.Experience{{Company: "NewCo"}},
					Education:      []models.Education{{Institution: "NewUni"}},
					Certifications: []models.Certification{{Name: "NewCert"}},
				},
			},
			verify: func(m *models.ConfigPortfolio) bool {
				return len(m.Content.Experiences) == 1 && m.Content.Experiences[0].Company == "NewCo" &&
					len(m.Content.Education) == 1 && m.Content.Education[0].Institution == "NewUni" &&
					len(m.Content.Certifications) == 1 && m.Content.Certifications[0].Name == "NewCert"
			},
		},
		{
			name: "some overridden",
			existing: &models.ConfigPortfolio{
				Content: models.Content{
					Experiences:    []models.Experience{{Company: "OldCo"}},
					Education:      []models.Education{{Institution: "OldUni"}},
					Certifications: []models.Certification{{Name: "OldCert"}},
				},
			},
			linkedin: &models.ConfigPortfolio{
				Content: models.Content{
					Experiences: []models.Experience{{Company: "NewCo"}},
					// Education and Certifications are empty
				},
			},
			verify: func(m *models.ConfigPortfolio) bool {
				return len(m.Content.Experiences) == 1 && m.Content.Experiences[0].Company == "NewCo" &&
					len(m.Content.Education) == 1 && m.Content.Education[0].Institution == "OldUni" &&
					len(m.Content.Certifications) == 1 && m.Content.Certifications[0].Name == "OldCert"
			},
		},
		{
			name: "none overridden",
			existing: &models.ConfigPortfolio{
				Content: models.Content{
					Experiences:    []models.Experience{{Company: "OldCo"}},
					Education:      []models.Education{{Institution: "OldUni"}},
					Certifications: []models.Certification{{Name: "OldCert"}},
				},
			},
			linkedin: &models.ConfigPortfolio{},
			verify: func(m *models.ConfigPortfolio) bool {
				return len(m.Content.Experiences) == 1 && m.Content.Experiences[0].Company == "OldCo" &&
					len(m.Content.Education) == 1 && m.Content.Education[0].Institution == "OldUni" &&
					len(m.Content.Certifications) == 1 && m.Content.Certifications[0].Name == "OldCert"
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merged := MergeConfigs(tt.linkedin, tt.existing)
			if !tt.verify(merged) {
				t.Errorf("MergeConfigs failed for %s", tt.name)
			}
		})
	}
}

func TestGetImportResult(t *testing.T) {
	diff := NewDiff()
	diff.Experiences.Added = make([]models.Experience, 2)
	diff.Education.Modified = make([]ChangePair[models.Education], 3)
	diff.Certifications.Removed = make([]models.Certification, 1)

	result := GetImportResult(diff)

	if result.ExperiencesAdded != 2 || result.EducationModified != 3 || result.CertificationsRemoved != 1 {
		t.Errorf("ImportResult unexpected: %+v", result)
	}
}

func TestPrintImportResult(t *testing.T) {
	result := ImportResult{
		ExperiencesAdded: 1,
		EducationAdded:    1,
		CertificationsAdded: 1,
	}
	// This is mainly to cover the code
	PrintImportResult(result)
}

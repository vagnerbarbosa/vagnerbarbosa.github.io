// Package comparator handles comparison between LinkedIn data and current config.
package comparator

import (
	"testing"

	"github.com/vagnerbarbosa/vagnerbarbosa.github.io/cmd/import-linkedin/internal/models"
)

func TestCompareExperiences(t *testing.T) {
	linkedin := []models.Experience{
		{Company: "Google", Role: "Engineer", StartDate: "Jan 2020"},
		{Company: "Microsoft", Role: "Dev", StartDate: "Jan 2021"},
	}

	current := []models.Experience{
		{Company: "Google", Role: "Engineer", StartDate: "Jan 2020"},
		{Company: "Apple", Role: "Designer", StartDate: "Jan 2019"},
	}

	diff := CompareExperiences(linkedin, current)

	if len(diff.Added) != 1 || diff.Added[0].Company != "Microsoft" {
		t.Errorf("Expected 1 added experience (Microsoft), got %d", len(diff.Added))
	}

	if len(diff.Removed) != 1 || diff.Removed[0].Company != "Apple" {
		t.Errorf("Expected 1 removed experience (Apple), got %d", len(diff.Removed))
	}

	if len(diff.Modified) != 0 {
		t.Errorf("Expected 0 modified experiences, got %d", len(diff.Modified))
	}
}

func TestCompareExperiences_WithModifications(t *testing.T) {
	// Note: When role changes, the ID (company#role) changes, so it's treated
	// as remove + add, not modify. This is expected behavior.
	linkedin := []models.Experience{
		{Company: "Google", Role: "Senior Engineer", StartDate: "Jan 2020", Location: "SF"},
	}

	current := []models.Experience{
		{Company: "Google", Role: "Engineer", StartDate: "Jan 2020", Location: "NYC"},
	}

	diff := CompareExperiences(linkedin, current)

	// Since company#role forms the ID, changing role creates new/remove
	// not modify. This is the expected behavior.
	if len(diff.Modified) != 0 {
		t.Errorf("Expected 0 modified experiences (role change = new ID), got %d", len(diff.Modified))
	}

	if len(diff.Added) != 1 {
		t.Errorf("Expected 1 added experience (new role = new ID), got %d", len(diff.Added))
	}

	if len(diff.Removed) != 1 {
		t.Errorf("Expected 1 removed experience (old role), got %d", len(diff.Removed))
	}
}

func TestCompareEducation(t *testing.T) {
	tests := []struct {
		name     string
		linkedin []models.Education
		current  []models.Education
		added    int
		modified int
		removed  int
	}{
		{
			name: "added and removed",
			linkedin: []models.Education{
				{Institution: "MIT", Degree: "Bachelor", StartDate: "Jan 2015"},
				{Institution: "Stanford", Degree: "Master", StartDate: "Jan 2020"},
			},
			current: []models.Education{
				{Institution: "MIT", Degree: "Bachelor", StartDate: "Jan 2015"},
			},
			added:    1,
			modified: 0,
			removed:  0,
		},
		{
			name: "modified",
			linkedin: []models.Education{
				{Institution: "MIT", Degree: "Bachelor", StartDate: "Jan 2016"},
			},
			current: []models.Education{
				{Institution: "MIT", Degree: "Bachelor", StartDate: "Jan 2015"},
			},
			added:    0,
			modified: 1,
			removed:  0,
		},
		{
			name: "identical",
			linkedin: []models.Education{
				{Institution: "MIT", Degree: "Bachelor", StartDate: "Jan 2015"},
			},
			current: []models.Education{
				{Institution: "MIT", Degree: "Bachelor", StartDate: "Jan 2015"},
			},
			added:    0,
			modified: 0,
			removed:  0,
		},
		{
			name: "all removed",
			linkedin: []models.Education{},
			current: []models.Education{
				{Institution: "MIT", Degree: "Bachelor", StartDate: "Jan 2015"},
			},
			added:    0,
			modified: 0,
			removed:  1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			diff := CompareEducation(tt.linkedin, tt.current)
			if len(diff.Added) != tt.added || len(diff.Modified) != tt.modified || len(diff.Removed) != tt.removed {
				t.Errorf("Got added=%d, mod=%d, rem=%d; want added=%d, mod=%d, rem=%d",
					len(diff.Added), len(diff.Modified), len(diff.Removed), tt.added, tt.modified, tt.removed)
			}
		})
	}
}

func TestCompareCertifications(t *testing.T) {
	tests := []struct {
		name     string
		linkedin []models.Certification
		current  []models.Certification
		added    int
		modified int
		removed  int
	}{
		{
			name: "added and removed",
			linkedin: []models.Certification{
				{Name: "AWS", Organization: "Amazon", IssueDate: "Jan 2020"},
			},
			current: []models.Certification{
				{Name: "Azure", Organization: "Microsoft", IssueDate: "Jan 2019"},
			},
			added:    1,
			modified: 0,
			removed:  1,
		},
		{
			name: "modified",
			linkedin: []models.Certification{
				{Name: "AWS", Organization: "Amazon", IssueDate: "Jan 2021"},
			},
			current: []models.Certification{
				{Name: "AWS", Organization: "Amazon", IssueDate: "Jan 2020"},
			},
			added:    0,
			modified: 1,
			removed:  0,
		},
		{
			name: "identical",
			linkedin: []models.Certification{
				{Name: "AWS", Organization: "Amazon", IssueDate: "Jan 2020"},
			},
			current: []models.Certification{
				{Name: "AWS", Organization: "Amazon", IssueDate: "Jan 2020"},
			},
			added:    0,
			modified: 0,
			removed:  0,
		},
		{
			name: "all removed",
			linkedin: []models.Certification{},
			current: []models.Certification{
				{Name: "AWS", Organization: "Amazon", IssueDate: "Jan 2020"},
			},
			added:    0,
			modified: 0,
			removed:  1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			diff := CompareCertifications(tt.linkedin, tt.current)
			if len(diff.Added) != tt.added || len(diff.Modified) != tt.modified || len(diff.Removed) != tt.removed {
				t.Errorf("Got added=%d, mod=%d, rem=%d; want added=%d, mod=%d, rem=%d",
					len(diff.Added), len(diff.Modified), len(diff.Removed), tt.added, tt.modified, tt.removed)
			}
		})
	}
}

func TestNewDiff(t *testing.T) {
	diff := NewDiff()

	if diff == nil {
		t.Fatal("NewDiff() returned nil")
	}

	if diff.HasChanges() {
		t.Error("Empty diff should not have changes")
	}

	if diff.CountAdded() != 0 || diff.CountModified() != 0 || diff.CountRemoved() != 0 {
		t.Error("Empty diff should have zero counts")
	}
}

func TestDiff_HasChanges(t *testing.T) {
	tests := []struct {
		name     string
		diff     *Diff
		expected bool
	}{
		{
			name:     "empty diff",
			diff:     NewDiff(),
			expected: false,
		},
		{
			name: "with added",
			diff: &Diff{
				Experiences: EntityDiff[models.Experience]{
					Added: []models.Experience{{Company: "Test", Role: "Dev"}},
				},
			},
			expected: true,
		},
		{
			name: "with modified",
			diff: &Diff{
				Education: EntityDiff[models.Education]{
					Modified: []ChangePair[models.Education]{{}},
				},
			},
			expected: true,
		},
		{
			name: "with removed",
			diff: &Diff{
				Certifications: EntityDiff[models.Certification]{
					Removed: []models.Certification{{Name: "Test"}},
				},
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.diff.HasChanges(); got != tt.expected {
				t.Errorf("HasChanges() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestGetChangedFields(t *testing.T) {
	tests := []struct {
		name     string
		old      models.Experience
		new      models.Experience
		expected []string
	}{
		{
			name: "no changes",
			old: models.Experience{
				Company: "Google", Role: "Engineer", StartDate: "Jan 2020",
				EndDate: "Dec 2022", Location: "NYC", Description: []string{"work"},
			},
			new: models.Experience{
				Company: "Google", Role: "Engineer", StartDate: "Jan 2020",
				EndDate: "Dec 2022", Location: "NYC", Description: []string{"work"},
			},
			expected: []string{},
		},
		{
			name: "all changes",
			old: models.Experience{
				Company: "Google", Role: "Engineer", StartDate: "Jan 2020",
				EndDate: "Dec 2022", Location: "NYC", Description: []string{"work"},
			},
			new: models.Experience{
				Company: "Microsoft", Role: "Senior Engineer", StartDate: "Feb 2020",
				EndDate: "Jan 2023", Location: "Seattle", Description: []string{"work", "lead"},
			},
			expected: []string{"company", "role", "start_date", "end_date", "location", "description"},
		},
		{
			name: "some changes",
			old: models.Experience{
				Company: "Google", Role: "Engineer", StartDate: "Jan 2020",
				EndDate: "", Location: "NYC", Description: []string{"work"},
			},
			new: models.Experience{
				Company: "Google", Role: "Senior Engineer", StartDate: "Jan 2020",
				EndDate: "Dec 2022", Location: "NYC", Description: []string{"work", "lead"},
			},
			expected: []string{"role", "end_date", "description"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fields := GetChangedFields(tt.old, tt.new)
			if len(fields) != len(tt.expected) {
				t.Errorf("Expected %d changed fields, got %d: %v", len(tt.expected), len(fields), fields)
			}

			fieldMap := make(map[string]bool)
			for _, f := range fields {
				fieldMap[f] = true
			}

			for _, exp := range tt.expected {
				if !fieldMap[exp] {
					t.Errorf("Expected field %s to be present in changed fields", exp)
				}
			}
		})
	}
}

func TestFormatID(t *testing.T) {
	tests := []struct {
		name     string
		entity   any
		expected string
	}{
		{
			name:     "experience",
			entity:   models.Experience{Company: "Google", Role: "Engineer"},
			expected: "Engineer @ Google",
		},
		{
			name:     "experience pointer",
			entity:   &models.Experience{Company: "Google", Role: "Engineer"},
			expected: "Engineer @ Google",
		},
		{
			name:     "experience pointer nil",
			entity:   (*models.Experience)(nil),
			expected: "",
		},
		{
			name:     "education",
			entity:   models.Education{Institution: "MIT", Degree: "Bachelor"},
			expected: "Bachelor at MIT",
		},
		{
			name:     "education pointer",
			entity:   &models.Education{Institution: "MIT", Degree: "Bachelor"},
			expected: "Bachelor at MIT",
		},
		{
			name:     "education pointer nil",
			entity:   (*models.Education)(nil),
			expected: "",
		},
		{
			name:     "certification",
			entity:   models.Certification{Name: "AWS", Organization: "Amazon"},
			expected: "AWS (Amazon)",
		},
		{
			name:     "certification pointer",
			entity:   &models.Certification{Name: "AWS", Organization: "Amazon"},
			expected: "AWS (Amazon)",
		},
		{
			name:     "certification pointer nil",
			entity:   (*models.Certification)(nil),
			expected: "",
		},
		{
			name:     "unknown",
			entity:   "string",
			expected: "Unknown",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatID(tt.entity); got != tt.expected {
				t.Errorf("FormatID() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestEqualityHelpers(t *testing.T) {
	t.Run("experiencesEqual", func(t *testing.T) {
		exp1 := models.Experience{
			Company: "Google", Role: "Engineer", StartDate: "Jan 2020",
			EndDate: "Dec 2022", Location: "NYC", Description: []string{"desc1"},
		}
		exp2 := models.Experience{
			Company: "Google", Role: "Engineer", StartDate: "Jan 2020",
			EndDate: "Dec 2022", Location: "NYC", Description: []string{"desc1"},
		}
		exp3 := models.Experience{
			Company: "Google", Role: "Lead", StartDate: "Jan 2020",
			EndDate: "Dec 2022", Location: "NYC", Description: []string{"desc1"},
		}

		if !experiencesEqual(exp1, exp2) {
			t.Error("Expected exp1 and exp2 to be equal")
		}
		if experiencesEqual(exp1, exp3) {
			t.Error("Expected exp1 and exp3 to be different")
		}
	})

	t.Run("educationEqual", func(t *testing.T) {
		edu1 := models.Education{
			Institution: "MIT", Degree: "Bachelor", Field: "CS",
			StartDate: "Jan 2015", EndDate: "Dec 2019", Description: []string{"desc1"},
		}
		edu2 := models.Education{
			Institution: "MIT", Degree: "Bachelor", Field: "CS",
			StartDate: "Jan 2015", EndDate: "Dec 2019", Description: []string{"desc1"},
		}
		edu3 := models.Education{
			Institution: "MIT", Degree: "Master", Field: "CS",
			StartDate: "Jan 2015", EndDate: "Dec 2019", Description: []string{"desc1"},
		}

		if !educationEqual(edu1, edu2) {
			t.Error("Expected edu1 and edu2 to be equal")
		}
		if educationEqual(edu1, edu3) {
			t.Error("Expected edu1 and edu3 to be different")
		}
	})

	t.Run("certificationsEqual", func(t *testing.T) {
		cert1 := models.Certification{
			Name: "AWS", Organization: "Amazon", IssueDate: "Jan 2020",
			ExpirationDate: "Jan 2023", CredentialID: "123", CredentialURL: "http://aws.com",
		}
		cert2 := models.Certification{
			Name: "AWS", Organization: "Amazon", IssueDate: "Jan 2020",
			ExpirationDate: "Jan 2023", CredentialID: "123", CredentialURL: "http://aws.com",
		}
		cert3 := models.Certification{
			Name: "AWS", Organization: "Amazon", IssueDate: "Jan 2020",
			ExpirationDate: "Jan 2023", CredentialID: "456", CredentialURL: "http://aws.com",
		}

		if !certificationsEqual(cert1, cert2) {
			t.Error("Expected cert1 and cert2 to be equal")
		}
		if certificationsEqual(cert1, cert3) {
			t.Error("Expected cert1 and cert3 to be different")
		}
	})
}

func TestCompareAll(t *testing.T) {
	linkedinExp := []models.Experience{
		{Company: "Google", Role: "Engineer", StartDate: "Jan 2020"}, // Modified
		{Company: "Microsoft", Role: "Dev", StartDate: "Jan 2021"},   // Added
	}
	currentExp := []models.Experience{
		{Company: "Google", Role: "Engineer", StartDate: "Jan 2019"}, // Modified (StartDate differs)
		{Company: "Apple", Role: "Designer", StartDate: "Jan 2018"},   // Removed
	}

	linkedinEdu := []models.Education{
		{Institution: "MIT", Degree: "Bachelor", StartDate: "Jan 2010"}, // Same
		{Institution: "Stanford", Degree: "Master", StartDate: "Jan 2015"}, // Added
	}
	currentEdu := []models.Education{
		{Institution: "MIT", Degree: "Bachelor", StartDate: "Jan 2010"}, // Same
	}

	linkedinCert := []models.Certification{}
	currentCert := []models.Certification{
		{Name: "AWS", Organization: "Amazon", IssueDate: "Jan 2020"}, // Removed
	}

	diff := CompareAll(linkedinExp, currentExp, linkedinEdu, currentEdu, linkedinCert, currentCert)

	if diff.CountAdded() != 2 { // Microsoft Exp + Stanford Edu
		t.Errorf("Expected 2 added items, got %d", diff.CountAdded())
	}
	if diff.CountModified() != 1 { // Google Exp
		t.Errorf("Expected 1 modified item, got %d", diff.CountModified())
	}
	if diff.CountRemoved() != 2 { // Apple Exp + AWS Cert
		t.Errorf("Expected 2 removed items, got %d", diff.CountRemoved())
	}
}

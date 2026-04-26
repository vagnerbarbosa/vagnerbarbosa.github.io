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
	linkedin := []models.Education{
		{Institution: "MIT", Degree: "Bachelor", StartDate: "Jan 2015"},
		{Institution: "Stanford", Degree: "Master", StartDate: "Jan 2020"},
	}

	current := []models.Education{
		{Institution: "MIT", Degree: "Bachelor", StartDate: "Jan 2015"},
	}

	diff := CompareEducation(linkedin, current)

	if len(diff.Added) != 1 || diff.Added[0].Institution != "Stanford" {
		t.Errorf("Expected 1 added education (Stanford), got %d", len(diff.Added))
	}

	if len(diff.Removed) != 0 {
		t.Errorf("Expected 0 removed education, got %d", len(diff.Removed))
	}
}

func TestCompareCertifications(t *testing.T) {
	linkedin := []models.Certification{
		{Name: "AWS", Organization: "Amazon", IssueDate: "Jan 2020"},
	}

	current := []models.Certification{
		{Name: "Azure", Organization: "Microsoft", IssueDate: "Jan 2019"},
	}

	diff := CompareCertifications(linkedin, current)

	if len(diff.Added) != 1 || diff.Added[0].Name != "AWS" {
		t.Errorf("Expected 1 added certification (AWS), got %d", len(diff.Added))
	}

	if len(diff.Removed) != 1 || diff.Removed[0].Name != "Azure" {
		t.Errorf("Expected 1 removed certification (Azure), got %d", len(diff.Removed))
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
	old := models.Experience{
		Company:     "Google",
		Role:        "Engineer",
		StartDate:   "Jan 2020",
		EndDate:     "",
		Location:    "NYC",
		Description: []string{"work"},
	}

	new := models.Experience{
		Company:     "Google",
		Role:        "Senior Engineer",
		StartDate:   "Jan 2020",
		EndDate:     "Dec 2022",
		Location:    "NYC",
		Description: []string{"work", "lead"},
	}

	fields := GetChangedFields(old, new)

	if len(fields) != 3 {
		t.Errorf("Expected 3 changed fields, got %d: %v", len(fields), fields)
	}

	// Check that expected fields are present
	fieldMap := make(map[string]bool)
	for _, f := range fields {
		fieldMap[f] = true
	}

	if !fieldMap["role"] {
		t.Error("Expected 'role' in changed fields")
	}
	if !fieldMap["end_date"] {
		t.Error("Expected 'end_date' in changed fields")
	}
	if !fieldMap["description"] {
		t.Error("Expected 'description' in changed fields")
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
			name:     "education",
			entity:   models.Education{Institution: "MIT", Degree: "Bachelor"},
			expected: "Bachelor at MIT",
		},
		{
			name:     "certification",
			entity:   models.Certification{Name: "AWS", Organization: "Amazon"},
			expected: "AWS (Amazon)",
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

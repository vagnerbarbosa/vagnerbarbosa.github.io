// Package parser handles CSV parsing for LinkedIn export files.
package parser

import (
	"strings"
	"testing"

	"github.com/vagnerbarbosa/vagnerbarbosa.github.io/cmd/import-linkedin/internal/models"
)

func TestEducationParser_ParseAll(t *testing.T) {
	csvData := `School Name,Degree Name,Field Of Study,Started On,Finished On,Description
Universidade Federal,Bachelor's Degree,Computer Science,Jan 2015,Dec 2019,Studied algorithms and data structures
Tech Institute,Master's Degree,AI,Jan 2020,Present,Research in machine learning`

	parser, err := NewEducationParserFromReader(strings.NewReader(csvData))
	if err != nil {
		t.Fatalf("Failed to create parser: %v", err)
	}

	education, err := parser.ParseAll()
	if err != nil {
		t.Fatalf("Failed to parse education: %v", err)
	}

	if len(education) != 2 {
		t.Errorf("Expected 2 education entries, got %d", len(education))
	}

	// Check first entry
	edu1 := education[0]
	if edu1.Institution != "Universidade Federal" {
		t.Errorf("Expected institution 'Universidade Federal', got '%s'", edu1.Institution)
	}
	if edu1.Degree != "Bachelor's Degree" {
		t.Errorf("Expected degree 'Bachelor's Degree', got '%s'", edu1.Degree)
	}
	if edu1.Field != "Computer Science" {
		t.Errorf("Expected field 'Computer Science', got '%s'", edu1.Field)
	}
	if edu1.StartDate != "Jan 2015" {
		t.Errorf("Expected start date 'Jan 2015', got '%s'", edu1.StartDate)
	}
	if edu1.EndDate != "Dez 2019" {
		t.Errorf("Expected end date 'Dez 2019', got '%s'", edu1.EndDate)
	}

	// Check second entry (with Present date)
	edu2 := education[1]
	if edu2.EndDate != "Presente" {
		t.Errorf("Expected end date 'Presente', got '%s'", edu2.EndDate)
	}
}

func TestEducationParser_Validate(t *testing.T) {
	tests := []struct {
		name        string
		csvData     string
		expectError bool
		errorCount  int
	}{
		{
			name: "valid data",
			csvData: `School Name,Degree Name,Field Of Study,Started On,Finished On,Description
University,Bachelor,CS,Jan 2015,Dec 2019,Studied`,
			expectError: false,
			errorCount:  0,
		},
		{
			name: "missing institution",
			csvData: `School Name,Degree Name,Field Of Study,Started On,Finished On,Description
,Bachelor,CS,Jan 2015,Dec 2019,Studied`,
			expectError: true,
			errorCount:  1,
		},
		{
			name: "missing degree",
			csvData: `School Name,Degree Name,Field Of Study,Started On,Finished On,Description
University,,CS,Jan 2015,Dec 2019,Studied`,
			expectError: true,
			errorCount:  1,
		},
		{
			name: "missing both dates",
			csvData: `School Name,Degree Name,Field Of Study,Started On,Finished On,Description
University,Bachelor,CS,,,Studied`,
			expectError: true,
			errorCount:  1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser, err := NewEducationParserFromReader(strings.NewReader(tt.csvData))
			if err != nil {
				t.Fatalf("Failed to create parser: %v", err)
			}

			errors := parser.Validate()
			hasErrors := len(errors) > 0

			if hasErrors != tt.expectError {
				t.Errorf("Expected error=%v, got errors=%v", tt.expectError, errors)
			}

			if tt.expectError && len(errors) != tt.errorCount {
				t.Errorf("Expected %d errors, got %d", tt.errorCount, len(errors))
			}
		})
	}
}

func TestEducation_Validate(t *testing.T) {
	tests := []struct {
		name    string
		edu     models.Education
		wantErr bool
	}{
		{
			name: "valid education",
			edu: models.Education{
				Institution: "University",
				Degree:      "Bachelor",
				StartDate:   "Jan 2015",
			},
			wantErr: false,
		},
		{
			name: "missing institution",
			edu: models.Education{
				Institution: "",
				Degree:      "Bachelor",
				StartDate:   "Jan 2015",
			},
			wantErr: true,
		},
		{
			name: "missing degree",
			edu: models.Education{
				Institution: "University",
				Degree:      "",
				StartDate:   "Jan 2015",
			},
			wantErr: true,
		},
		{
			name: "missing dates",
			edu: models.Education{
				Institution: "University",
				Degree:      "Bachelor",
				StartDate:   "",
				EndDate:     "",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.edu.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// Package parser handles CSV parsing for LinkedIn export files.
package parser

import (
	"os"
	"strings"
	"testing"

	"github.com/vagnerbarbosa/vagnerbarbosa.github.io/cmd/import-linkedin/internal/models"
)

func TestNewEducationParser(t *testing.T) {
	content := "School Name,Degree Name,Field Of Study,Start Date,End Date,Description\nMIT,BSc,CS,Jan 2020,Jan 2024,Studied"
	tmpFile, err := os.CreateTemp("", "edu*.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.Write([]byte(content)); err != nil {
		t.Fatal(err)
	}
	tmpFile.Close()

	t.Run("valid file", func(t *testing.T) {
		parser, err := NewEducationParser(tmpFile.Name())
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if parser == nil {
			t.Error("Expected parser to be non-nil")
		}
	})

	t.Run("missing required column", func(t *testing.T) {
		contentBad := "Start Date\nJan 2020"
		tmpFileBad, _ := os.CreateTemp("", "edubad*.csv")
		defer os.Remove(tmpFileBad.Name())
		tmpFileBad.Write([]byte(contentBad))
		tmpFileBad.Close()

		_, err := NewEducationParser(tmpFileBad.Name())
		if err == nil {
			t.Error("Expected error for missing required columns")
		}
	})

	t.Run("non-existent file", func(t *testing.T) {
		_, err := NewEducationParser("non_existent.csv")
		if err == nil {
			t.Error("Expected error for non-existent file")
		}
	})
}

func TestNewEducationParserFromReader(t *testing.T) {
	t.Run("invalid csv header", func(t *testing.T) {
		parser, err := NewEducationParserFromReader(strings.NewReader(""))
		if err == nil {
			t.Error("Expected error for empty input")
		}
		if parser != nil {
			t.Error("Expected parser to be nil on error")
		}
	})

	t.Run("simulated read error", func(t *testing.T) {
		parser, err := NewEducationParserFromReader(&failingReader{failAt: 0})
		if err == nil {
			t.Error("Expected error for simulated read failure")
		}
		if parser != nil {
			t.Error("Expected parser to be nil")
		}
	})
}

func TestEducationParser_Close(t *testing.T) {
	parser, _ := NewEducationParserFromReader(strings.NewReader("School Name,Degree Name\nMIT,BSc"))
	if err := parser.Close(); err != nil {
		t.Errorf("Close() error = %v", err)
	}
}

func TestEducationParser_ParseAll(t *testing.T) {
	csvData := `School Name,Degree Name,Field Of Study,Start Date,End Date,Description
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
			csvData: `School Name,Degree Name,Field Of Study,Start Date,End Date,Description
University,Bachelor,CS,Jan 2015,Dec 2019,Studied`,
			expectError: false,
			errorCount:  0,
		},
		{
			name: "missing institution",
			csvData: `School Name,Degree Name,Field Of Study,Start Date,End Date,Description
,Bachelor,CS,Jan 2015,Dec 2019,Studied`,
			expectError: true,
			errorCount:  1,
		},
		{
			name: "missing degree",
			csvData: `School Name,Degree Name,Field Of Study,Start Date,End Date,Description
University,,CS,Jan 2015,Dec 2019,Studied`,
			expectError: true,
			errorCount:  1,
		},
		{
			name: "missing both dates",
			csvData: `School Name,Degree Name,Field Of Study,Start Date,End Date,Description
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

func TestEducationParser_ParseAll_Errors(t *testing.T) {
	tests := []struct {
		name        string
		csvData     string
		expectError bool
		expectedLen int
	}{
		{
			name:        "invalid row - missing school",
			csvData:     "School Name,Degree Name,Start Date,End Date\n,BSc,Jan 2020,Jan 2024",
			expectError: false,
			expectedLen: 0,
		},
		{
			name:        "invalid row - missing degree",
			csvData:     "School Name,Degree Name,Start Date,End Date\nMIT,,Jan 2020,Jan 2024",
			expectError: false,
			expectedLen: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser, err := NewEducationParserFromReader(strings.NewReader(tt.csvData))
			if err != nil {
				t.Fatalf("Failed to create parser: %v", err)
			}

			education, err := parser.ParseAll()
			if (err != nil) != tt.expectError {
				t.Errorf("ParseAll() error = %v, expectError %v", err, tt.expectError)
			}

			if len(education) != tt.expectedLen {
				t.Errorf("Expected %d education entries, got %d", tt.expectedLen, len(education))
			}
		})
	}
}

func TestEducationParser_Validate_Errors(t *testing.T) {
	csvData := "School Name,Degree Name\nMIT,BSc\nStanford,PhD"

	t.Run("ParseAll failure", func(t *testing.T) {
		reader := &failingReader{
			data:   []byte(csvData),
			failAt: 20,
		}
		parser, _ := NewEducationParserFromReader(reader)
		_, err := parser.ParseAll()
		if err == nil {
			t.Error("Expected error during ParseAll")
		}
	})
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

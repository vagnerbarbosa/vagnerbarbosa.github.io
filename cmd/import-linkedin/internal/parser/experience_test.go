// Package parser handles CSV parsing for LinkedIn export files.
package parser

import (
	"os"
	"strings"
	"testing"

	"github.com/vagnerbarbosa/vagnerbarbosa.github.io/cmd/import-linkedin/internal/models"
)

func TestNewExperienceParser(t *testing.T) {
	content := "Company Name,Title,Started On,Finished On,Description,Location\nAcme Corp,Engineer,Jan 2020,Mar 2022,Work,NY"
	tmpFile, err := os.CreateTemp("", "exp*.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.Write([]byte(content)); err != nil {
		t.Fatal(err)
	}
	tmpFile.Close()

	t.Run("valid file", func(t *testing.T) {
		parser, err := NewExperienceParser(tmpFile.Name())
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if parser == nil {
			t.Error("Expected parser to be non-nil")
		}
	})

	t.Run("missing required column", func(t *testing.T) {
		contentBad := "Started On\nJan 2020"
		tmpFileBad, _ := os.CreateTemp("", "expbad*.csv")
		defer os.Remove(tmpFileBad.Name())
		tmpFileBad.Write([]byte(contentBad))
		tmpFileBad.Close()

		_, err := NewExperienceParser(tmpFileBad.Name())
		if err == nil {
			t.Error("Expected error for missing required columns")
		}
	})

	t.Run("non-existent file", func(t *testing.T) {
		_, err := NewExperienceParser("non_existent.csv")
		if err == nil {
			t.Error("Expected error for non-existent file")
		}
	})
}

func TestNewExperienceParserFromReader(t *testing.T) {
	t.Run("invalid csv header", func(t *testing.T) {
		parser, err := NewExperienceParserFromReader(strings.NewReader(""))
		if err == nil {
			t.Error("Expected error for empty input")
		}
		if parser != nil {
			t.Error("Expected parser to be nil on error")
		}
	})

	t.Run("simulated read error", func(t *testing.T) {
		parser, err := NewExperienceParserFromReader(&failingReader{failAt: 0})
		if err == nil {
			t.Error("Expected error for simulated read failure")
		}
		if parser != nil {
			t.Error("Expected parser to be nil")
		}
	})
}

func TestExperienceParser_Close(t *testing.T) {
	parser, _ := NewExperienceParserFromReader(strings.NewReader("Company Name,Title,Started On\nAcme Corp,Engineer,Jan 2020"))
	if err := parser.Close(); err != nil {
		t.Errorf("Close() error = %v", err)
	}
}

func TestExperienceParser_ParseAll(t *testing.T) {
	// Test CSV data
	csvData := `Company Name,Title,Started On,Finished On,Description,Location
Acme Corp,Software Engineer,Jan 2020,Mar 2022,Developed key features. Led team of 5.,New York
Tech Inc,Senior Developer,Apr 2022,Present,Managing projects and mentoring junior devs,San Francisco
StartupXYZ,Intern,Jun 2019,Dec 2019,Learned the ropes,Remote`

	parser, err := NewExperienceParserFromReader(strings.NewReader(csvData))
	if err != nil {
		t.Fatalf("Failed to create parser: %v", err)
	}

	experiences, err := parser.ParseAll()
	if err != nil {
		t.Fatalf("Failed to parse experiences: %v", err)
	}

	if len(experiences) != 3 {
		t.Errorf("Expected 3 experiences, got %d", len(experiences))
	}

	// Check first experience
	exp1 := experiences[0]
	if exp1.Company != "Acme Corp" {
		t.Errorf("Expected company 'Acme Corp', got '%s'", exp1.Company)
	}
	if exp1.Role != "Software Engineer" {
		t.Errorf("Expected role 'Software Engineer', got '%s'", exp1.Role)
	}
	if exp1.StartDate != "Jan 2020" {
		t.Errorf("Expected start date 'Jan 2020', got '%s'", exp1.StartDate)
	}
	if exp1.EndDate != "Mar 2022" {
		t.Errorf("Expected end date 'Mar 2022', got '%s'", exp1.EndDate)
	}
	if exp1.Location != "New York" {
		t.Errorf("Expected location 'New York', got '%s'", exp1.Location)
	}
	if len(exp1.Description) != 2 {
		t.Errorf("Expected 2 description bullets, got %d", len(exp1.Description))
	}

	// Check second experience (with Present date)
	exp2 := experiences[1]
	if exp2.EndDate != "Presente" {
		t.Errorf("Expected end date 'Presente', got '%s'", exp2.EndDate)
	}

	// Check ID generation
	expectedID := "Acme Corp#Software Engineer"
	if exp1.ID() != expectedID {
		t.Errorf("Expected ID '%s', got '%s'", expectedID, exp1.ID())
	}
}

func TestExperienceParser_ParseAll_EmptyCSV(t *testing.T) {
	csvData := `Company Name,Title,Started On,Finished On,Description,Location
`

	parser, err := NewExperienceParserFromReader(strings.NewReader(csvData))
	if err != nil {
		t.Fatalf("Failed to create parser: %v", err)
	}

	experiences, err := parser.ParseAll()
	if err != nil {
		t.Fatalf("Failed to parse experiences: %v", err)
	}

	if len(experiences) != 0 {
		t.Errorf("Expected 0 experiences for empty CSV, got %d", len(experiences))
	}
}

func TestExperienceParser_Validate(t *testing.T) {
	tests := []struct {
		name        string
		csvData     string
		expectError bool
		errorCount  int
	}{
		{
			name: "valid data",
			csvData: `Company Name,Title,Started On,Finished On,Description,Location
Acme Corp,Engineer,Jan 2020,Mar 2022,Work,NY`,
			expectError: false,
			errorCount:  0,
		},
		{
			name: "missing company",
			csvData: `Company Name,Title,Started On,Finished On,Description,Location
,Engineer,Jan 2020,,,NY`,
			expectError: true,
			errorCount:  1,
		},
		{
			name: "missing title",
			csvData: `Company Name,Title,Started On,Finished On,Description,Location
Acme Corp,,Jan 2020,,,NY`,
			expectError: true,
			errorCount:  1,
		},
		{
			name: "invalid date",
			csvData: `Company Name,Title,Started On,Finished On,Description,Location
Acme Corp,Engineer,Invalid Date,,,NY`,
			expectError: true,
			errorCount:  1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser, err := NewExperienceParserFromReader(strings.NewReader(tt.csvData))
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

func TestExperienceParser_ParseAll_Errors(t *testing.T) {
	tests := []struct {
		name        string
		csvData     string
		expectError bool
	}{
		{
			name:        "invalid row - missing company",
			csvData:     "Company Name,Title,Started On\n,Engineer,Jan 2020",
			expectError: true,
		},
		{
			name:        "invalid row - missing title",
			csvData:     "Company Name,Title,Started On\nAcme Corp,,Jan 2020",
			expectError: true,
		},
		{
			name:        "invalid row - missing start date",
			csvData:     "Company Name,Title,Started On\nAcme Corp,Engineer,",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser, err := NewExperienceParserFromReader(strings.NewReader(tt.csvData))
			if err != nil {
				t.Fatalf("Failed to create parser: %v", err)
			}

			_, err = parser.ParseAll()
			if (err != nil) != tt.expectError {
				t.Errorf("ParseAll() error = %v, expectError %v", err, tt.expectError)
			}
		})
	}
}

func TestExperienceParser_Validate_Errors(t *testing.T) {
	csvData := "Company Name,Title,Started On\nAcme Corp,Engineer,Jan 2020\nTech Inc,Dev,Mar 2021"

	t.Run("ParseAll failure with read error", func(t *testing.T) {
		reader := &failingReader{
			data:   []byte(csvData),
			failAt: 20,
		}
		parser, _ := NewExperienceParserFromReader(reader)
		_, err := parser.ParseAll()
		if err == nil {
			t.Error("Expected error during ParseAll")
		}
	})
}

func TestExperience_Validate(t *testing.T) {
	tests := []struct {
		name    string
		exp     models.Experience
		wantErr bool
	}{
		{
			name: "valid experience",
			exp: models.Experience{
				Company:   "Acme Corp",
				Role:      "Engineer",
				StartDate: "Jan 2020",
			},
			wantErr: false,
		},
		{
			name: "missing company",
			exp: models.Experience{
				Company:   "",
				Role:      "Engineer",
				StartDate: "Jan 2020",
			},
			wantErr: true,
		},
		{
			name: "missing role",
			exp: models.Experience{
				Company:   "Acme Corp",
				Role:      "",
				StartDate: "Jan 2020",
			},
			wantErr: true,
		},
		{
			name: "missing start date",
			exp: models.Experience{
				Company:   "Acme Corp",
				Role:      "Engineer",
				StartDate: "",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.exp.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestConvertDate_Experiences(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Jan 2020", "Jan 2020"},
		{"Feb 2020", "Fev 2020"},
		{"Mar 2020", "Mar 2020"},
		{"Apr 2020", "Abr 2020"},
		{"May 2020", "Mai 2020"},
		{"Jun 2020", "Jun 2020"},
		{"Jul 2020", "Jul 2020"},
		{"Aug 2020", "Ago 2020"},
		{"Sep 2020", "Set 2020"},
		{"Oct 2020", "Out 2020"},
		{"Nov 2020", "Nov 2020"},
		{"Dec 2020", "Dez 2020"},
		{"Present", "Presente"},
		{"Jan 2020 - Mar 2022", "Jan 2020 - Mar 2022"},
		{"Feb 2020 - Present", "Fev 2020 - Presente"},
	}

	for _, tt := range tests {
		result := ConvertDate(tt.input)
		if result != tt.expected {
			t.Errorf("ConvertDate(%q) = %q, want %q", tt.input, result, tt.expected)
		}
	}
}

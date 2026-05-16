package parser

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/vagnerbarbosa/vagnerbarbosa.github.io/cmd/import-linkedin/internal/models"
)

type failingReader struct {
	data   []byte
	off    int
	failAt int
}

func (r *failingReader) Read(p []byte) (n int, err error) {
	if r.off >= r.failAt {
		return 0, fmt.Errorf("simulated read failure")
	}
	n = copy(p, r.data[r.off:])
	r.off += n
	return n, nil
}

func TestNewCertificationParser(t *testing.T) {
	content := "Name,Url,Authority,Started On,Finished On,License Number\nAWS,url,Amazon,Jan 2020,Jan 2023,ABC"
	tmpFile, err := os.CreateTemp("", "cert*.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.Write([]byte(content)); err != nil {
		t.Fatal(err)
	}
	tmpFile.Close()

	t.Run("valid file", func(t *testing.T) {
		parser, err := NewCertificationParser(tmpFile.Name())
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if parser == nil {
			t.Error("Expected parser to be non-nil")
		}
	})

	t.Run("missing required column", func(t *testing.T) {
		contentBad := "Url,Started On\nurl,Jan 2020"
		tmpFileBad, _ := os.CreateTemp("", "certbad*.csv")
		defer os.Remove(tmpFileBad.Name())
		tmpFileBad.Write([]byte(contentBad))
		tmpFileBad.Close()

		_, err := NewCertificationParser(tmpFileBad.Name())
		if err == nil {
			t.Error("Expected error for missing required columns")
		}
	})

	t.Run("non-existent file", func(t *testing.T) {
		_, err := NewCertificationParser("non_existent.csv")
		if err == nil {
			t.Error("Expected error for non-existent file")
		}
	})
}

func TestNewCertificationParserFromReader(t *testing.T) {
	t.Run("invalid csv header", func(t *testing.T) {
		parser, err := NewCertificationParserFromReader(strings.NewReader(""))
		if err == nil {
			t.Error("Expected error for empty input")
		}
		if parser != nil {
			t.Error("Expected parser to be nil on error")
		}
	})

	t.Run("simulated read error", func(t *testing.T) {
		// Use a reader that fails immediately
		parser, err := NewCertificationParserFromReader(&failingReader{failAt: 0})
		if err == nil {
			t.Error("Expected error for simulated read failure")
		}
		if parser != nil {
			t.Error("Expected parser to be nil")
		}
	})
}

func TestCertificationParser_ParseAll(t *testing.T) {
	tests := []struct {
		name        string
		csvData     string
		expectError bool
		expectedLen int
	}{
		{
			name: "valid data",
			csvData: `Name,Url,Authority,Started On,Finished On,License Number
AWS Solutions Architect,https://aws.amazon.com/certification,Amazon Web Services,Jan 2020,Jan 2023,ABC123
Kubernetes Admin,https://cncf.io/certification,CNCF,Mar 2021,Mar 2024,XYZ789`,
			expectError: false,
			expectedLen: 2,
		},
		{
			name: "invalid row - missing name",
			csvData: `Name,Url,Authority,Started On,Finished On,License Number
,https://test.com,Amazon,Jan 2020,Jan 2023,ABC`,
			expectError: true,
		},
		{
			name: "invalid row - missing organization",
			csvData: `Name,Url,Authority,Started On,Finished On,License Number
AWS Cert,https://test.com,,Jan 2020,Jan 2023,ABC`,
			expectError: true,
		},
		{
			name: "invalid row - missing issue date",
			csvData: `Name,Url,Authority,Started On,Finished On,License Number
AWS Cert,https://test.com,Amazon,,Jan 2023,ABC`,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser, err := NewCertificationParserFromReader(strings.NewReader(tt.csvData))
			if err != nil {
				t.Fatalf("Failed to create parser: %v", err)
			}

			certs, err := parser.ParseAll()
			if (err != nil) != tt.expectError {
				t.Errorf("ParseAll() error = %v, expectError %v", err, tt.expectError)
			}

			if !tt.expectError && len(certs) != tt.expectedLen {
				t.Errorf("Expected %d certifications, got %d", tt.expectedLen, len(certs))
			}
		})
	}
}

func TestCertificationParser_Close(t *testing.T) {
	parser, _ := NewCertificationParserFromReader(strings.NewReader("Name,Authority\nCert,Org"))
	if err := parser.Close(); err != nil {
		t.Errorf("Close() error = %v", err)
	}
}

func TestCertificationParser_Validate(t *testing.T) {
	tests := []struct {
		name        string
		csvData     string
		expectError bool
		errorCount  int
	}{
		{
			name: "valid data",
			csvData: `Name,Url,Authority,Started On,Finished On,License Number
AWS Cert,https://test.com,Amazon,Jan 2020,Jan 2023,ABC`,
			expectError: false,
			errorCount:  0,
		},
		{
			name: "missing name",
			csvData: `Name,Url,Authority,Started On,Finished On,License Number
,https://test.com,Amazon,Jan 2020,,ABC`,
			expectError: true,
			errorCount:  1,
		},
		{
			name: "missing organization",
			csvData: `Name,Url,Authority,Started On,Finished On,License Number
AWS Cert,https://test.com,,Jan 2020,,ABC`,
			expectError: true,
			errorCount:  1,
		},
		{
			name: "invalid date format",
			csvData: `Name,Url,Authority,Started On,Finished On,License Number
AWS Cert,https://test.com,Amazon,Invalid Date,Jan 2023,ABC`,
			expectError: true,
			errorCount:  1,
		},
		{
			name: "multiple errors",
			csvData: `Name,Url,Authority,Started On,Finished On,License Number
,https://test.com,,Invalid Date,Invalid Date,ABC`,
			expectError: true,
			errorCount:  4, // Name, Authority, Started On, Finished On
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser, err := NewCertificationParserFromReader(strings.NewReader(tt.csvData))
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

func TestCertificationParser_ReadFailure(t *testing.T) {
	csvData := "Name,Authority\nCert,Org\nAnother,AnotherOrg"

	t.Run("ParseAll failure", func(t *testing.T) {
		reader := &failingReader{
			data:   []byte(csvData),
			failAt: 20, // Approximately after first record
		}
		parser, _ := NewCertificationParserFromReader(reader)
		_, err := parser.ParseAll()
		if err == nil {
			t.Error("Expected error during ParseAll")
		}
	})

	t.Run("Validate failure", func(t *testing.T) {
		reader := &failingReader{
			data:   []byte(csvData),
			failAt: 20,
		}
		parser, _ := NewCertificationParserFromReader(reader)
		errors := parser.Validate()
		if len(errors) == 0 {
			t.Error("Expected error during Validate")
		}
	})
}

func TestCertification_Validate(t *testing.T) {
	tests := []struct {
		name    string
		cert    models.Certification
		wantErr bool
	}{
		{
			name: "valid certification",
			cert: models.Certification{
				Name:         "AWS Cert",
				Organization: "Amazon",
				IssueDate:    "Jan 2020",
			},
			wantErr: false,
		},
		{
			name: "missing name",
			cert: models.Certification{
				Name:         "",
				Organization: "Amazon",
				IssueDate:    "Jan 2020",
			},
			wantErr: true,
		},
		{
			name: "missing organization",
			cert: models.Certification{
				Name:         "AWS Cert",
				Organization: "",
				IssueDate:    "Jan 2020",
			},
			wantErr: true,
		},
		{
			name: "missing issue date",
			cert: models.Certification{
				Name:         "AWS Cert",
				Organization: "Amazon",
				IssueDate:    "",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.cert.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

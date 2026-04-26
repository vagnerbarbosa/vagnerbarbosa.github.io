// Package parser handles CSV parsing for LinkedIn export files.
package parser

import (
	"strings"
	"testing"

	"github.com/vagnerbarbosa/vagnerbarbosa.github.io/cmd/import-linkedin/internal/models"
)

func TestCertificationParser_ParseAll(t *testing.T) {
	csvData := `Certification Name,Certification Authority,Started On,Finished On,License Number,Certification URL
AWS Solutions Architect,Amazon Web Services,Jan 2020,Jan 2023,ABC123,https://aws.amazon.com/certification
Kubernetes Admin,CNCF,Mar 2021,Mar 2024,XYZ789,https://cncf.io/certification`

	parser, err := NewCertificationParserFromReader(strings.NewReader(csvData))
	if err != nil {
		t.Fatalf("Failed to create parser: %v", err)
	}

	certs, err := parser.ParseAll()
	if err != nil {
		t.Fatalf("Failed to parse certifications: %v", err)
	}

	if len(certs) != 2 {
		t.Errorf("Expected 2 certifications, got %d", len(certs))
	}

	// Check first certification
	cert1 := certs[0]
	if cert1.Name != "AWS Solutions Architect" {
		t.Errorf("Expected name 'AWS Solutions Architect', got '%s'", cert1.Name)
	}
	if cert1.Organization != "Amazon Web Services" {
		t.Errorf("Expected organization 'Amazon Web Services', got '%s'", cert1.Organization)
	}
	if cert1.IssueDate != "Jan 2020" {
		t.Errorf("Expected issue date 'Jan 2020', got '%s'", cert1.IssueDate)
	}
	if cert1.ExpirationDate != "Jan 2023" {
		t.Errorf("Expected expiration date 'Jan 2023', got '%s'", cert1.ExpirationDate)
	}
	if cert1.CredentialID != "ABC123" {
		t.Errorf("Expected credential ID 'ABC123', got '%s'", cert1.CredentialID)
	}
	if cert1.CredentialURL != "https://aws.amazon.com/certification" {
		t.Errorf("Expected credential URL 'https://aws.amazon.com/certification', got '%s'", cert1.CredentialURL)
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
			csvData: `Certification Name,Certification Authority,Started On,Finished On,License Number,Certification URL
AWS Cert,Amazon,Jan 2020,Jan 2023,ABC,https://test.com`,
			expectError: false,
			errorCount:  0,
		},
		{
			name: "missing name",
			csvData: `Certification Name,Certification Authority,Started On,Finished On,License Number,Certification URL
,Amazon,Jan 2020,,ABC,`,
			expectError: true,
			errorCount:  1,
		},
		{
			name: "missing organization",
			csvData: `Certification Name,Certification Authority,Started On,Finished On,License Number,Certification URL
AWS Cert,,Jan 2020,,ABC,`,
			expectError: true,
			errorCount:  1,
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

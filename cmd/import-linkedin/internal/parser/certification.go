// Package parser handles CSV parsing for LinkedIn export files.
package parser

import (
	"fmt"
	"io"

	"github.com/vagnerbarbosa/vagnerbarbosa.github.io/cmd/import-linkedin/internal/models"
)

// CertificationParser handles parsing of Certifications.csv from LinkedIn.
type CertificationParser struct {
	csvPath string
	reader  *CSVReader
}

// Required columns for certification parsing (LinkedIn export format)
var certificationRequiredColumns = []string{
	"Name",
	"Authority",
}

// Optional columns for certification parsing (LinkedIn export format)
var certificationOptionalColumns = []string{
	"Started On",
	"Finished On",
	"License Number",
	"Url",
}

// NewCertificationParser creates a new certification parser for the given CSV file.
func NewCertificationParser(csvPath string) (*CertificationParser, error) {
	reader, err := NewCSVReader(csvPath)
	if err != nil {
		return nil, fmt.Errorf("failed to create CSV reader: %w", err)
	}

	// Validate required columns
	missing := reader.ValidateColumns(certificationRequiredColumns)
	if len(missing) > 0 {
		return nil, fmt.Errorf("missing required columns: %v", missing)
	}

	return &CertificationParser{
		csvPath: csvPath,
		reader:  reader,
	}, nil
}

// NewCertificationParserFromReader creates a parser from an io.Reader (for testing).
func NewCertificationParserFromReader(r io.Reader) (*CertificationParser, error) {
	reader, err := NewCSVReaderFromIO(r)
	if err != nil {
		return nil, fmt.Errorf("failed to create CSV reader: %w", err)
	}

	return &CertificationParser{
		reader: reader,
	}, nil
}

// ParseAll parses all certifications from the CSV file.
func (p *CertificationParser) ParseAll() ([]models.Certification, error) {
	var certifications []models.Certification
	lineNum := 1 // Start at 1 to account for header

	for {
		row, err := p.reader.Next()
		lineNum++

		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("error reading row at line %d: %w", lineNum, err)
		}

		cert, err := p.parseRow(row, lineNum)
		if err != nil {
			return nil, err
		}

		certifications = append(certifications, *cert)
	}

	return certifications, nil
}

// parseRow converts a CSV row to a Certification struct.
func (p *CertificationParser) parseRow(row map[string]string, lineNum int) (*models.Certification, error) {
	cert := models.Certification{
		Name:           row["Name"],
		Organization:   row["Authority"],
		IssueDate:      ConvertDate(row["Started On"]),
		ExpirationDate: ConvertDate(row["Finished On"]),
		CredentialID:   row["License Number"],
		CredentialURL:  row["Url"],
	}

	// Validate the certification
	if err := cert.Validate(); err != nil {
		return nil, fmt.Errorf("validation error at line %d: %w", lineNum, err)
	}

	return &cert, nil
}

// Close closes the parser and its underlying resources.
func (p *CertificationParser) Close() error {
	return p.reader.Close()
}

// Validate validates the certifications CSV file without parsing all entries.
// Returns a slice of validation errors found.
func (p *CertificationParser) Validate() []error {
	var errors []error
	lineNum := 1

	for {
		row, err := p.reader.Next()
		lineNum++

		if err == io.EOF {
			break
		}
		if err != nil {
			errors = append(errors, fmt.Errorf("parse error at line %d: %w", lineNum, err))
			continue
		}

		// Check required fields
		if row["Name"] == "" {
			errors = append(errors, models.NewParseError(lineNum, "Name", "cannot be empty"))
		}
		if row["Authority"] == "" {
			errors = append(errors, models.NewParseError(lineNum, "Authority", "cannot be empty"))
		}

		// Validate date format
		if row["Started On"] != "" && !ValidateDate(row["Started On"]) {
			errors = append(errors, models.NewParseError(lineNum, "Started On", "invalid date format"))
		}
		if row["Finished On"] != "" && !ValidateDate(row["Finished On"]) {
			errors = append(errors, models.NewParseError(lineNum, "Finished On", "invalid date format"))
		}
	}

	return errors
}

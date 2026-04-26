// Package parser handles CSV parsing for LinkedIn export files.
package parser

import (
	"fmt"
	"io"

	"github.com/vagnerbarbosa/vagnerbarbosa.github.io/cmd/import-linkedin/internal/models"
	"github.com/vagnerbarbosa/vagnerbarbosa.github.io/cmd/import-linkedin/internal/transformer"
)

// EducationParser handles parsing of Education.csv from LinkedIn.
type EducationParser struct {
	csvPath string
	reader  *CSVReader
}

// Required columns for education parsing
var educationRequiredColumns = []string{
	"School Name",
	"Degree Name",
}

// Optional columns for education parsing
var educationOptionalColumns = []string{
	"Field Of Study",
	"Start Date",
	"End Date",
	"Description",
}

// NewEducationParser creates a new education parser for the given CSV file.
func NewEducationParser(csvPath string) (*EducationParser, error) {
	reader, err := NewCSVReader(csvPath)
	if err != nil {
		return nil, fmt.Errorf("failed to create CSV reader: %w", err)
	}

	// Validate required columns
	missing := reader.ValidateColumns(educationRequiredColumns)
	if len(missing) > 0 {
		return nil, fmt.Errorf("missing required columns: %v", missing)
	}

	return &EducationParser{
		csvPath: csvPath,
		reader:  reader,
	}, nil
}

// NewEducationParserFromReader creates a parser from an io.Reader (for testing).
func NewEducationParserFromReader(r io.Reader) (*EducationParser, error) {
	reader, err := NewCSVReaderFromIO(r)
	if err != nil {
		return nil, fmt.Errorf("failed to create CSV reader: %w", err)
	}

	return &EducationParser{
		reader: reader,
	}, nil
}

// ParseAll parses all education entries from the CSV file.
func (p *EducationParser) ParseAll() ([]models.Education, error) {
	var education []models.Education
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

		edu, err := p.parseRow(row, lineNum)
		if err != nil {
			return nil, err
		}

		education = append(education, *edu)
	}

	return education, nil
}

// parseRow converts a CSV row to an Education struct.
func (p *EducationParser) parseRow(row map[string]string, lineNum int) (*models.Education, error) {
	edu := models.Education{
		Institution: row["School Name"],
		Degree:      row["Degree Name"],
		Field:       row["Field Of Study"],
		StartDate:   ConvertDate(row["Start Date"]),
		EndDate:     ConvertDate(row["End Date"]),
	}

	// Split description into bullets
	description := row["Description"]
	if description != "" {
		edu.Description = transformer.SplitDescription(description)
	}

	// Validate the education entry
	if err := edu.Validate(); err != nil {
		return nil, fmt.Errorf("validation error at line %d: %w", lineNum, err)
	}

	return &edu, nil
}

// Close closes the parser and its underlying resources.
func (p *EducationParser) Close() error {
	return p.reader.Close()
}

// Validate validates the education CSV file without parsing all entries.
// Returns a slice of validation errors found.
func (p *EducationParser) Validate() []error {
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
		if row["School Name"] == "" {
			errors = append(errors, models.NewParseError(lineNum, "School Name", "cannot be empty"))
		}
		if row["Degree Name"] == "" {
			errors = append(errors, models.NewParseError(lineNum, "Degree Name", "cannot be empty"))
		}

		// At least one date should be present
		if row["Start Date"] == "" && row["End Date"] == "" {
			errors = append(errors, models.NewParseError(lineNum, "dates", "at least one date must be present"))
		}

		// Validate date format
		if row["Start Date"] != "" && !ValidateDate(row["Start Date"]) {
			errors = append(errors, models.NewParseError(lineNum, "Start Date", "invalid date format"))
		}
		if row["End Date"] != "" && !ValidateDate(row["End Date"]) {
			errors = append(errors, models.NewParseError(lineNum, "End Date", "invalid date format"))
		}
	}

	return errors
}

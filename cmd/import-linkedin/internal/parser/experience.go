// Package parser handles CSV parsing for LinkedIn export files.
package parser

import (
	"fmt"
	"io"

	"github.com/vagnerbarbosa/vagnerbarbosa.github.io/cmd/import-linkedin/internal/models"
	"github.com/vagnerbarbosa/vagnerbarbosa.github.io/cmd/import-linkedin/internal/transformer"
)

// ExperienceParser handles parsing of Experiences.csv from LinkedIn.
type ExperienceParser struct {
	csvPath string
	reader  *CSVReader
}

// RequiredColumns lists the CSV columns required for experience parsing.
var experienceRequiredColumns = []string{
	"Company Name",
	"Title",
	"Started On",
}

// OptionalColumns lists optional columns that enhance the experience data.
var experienceOptionalColumns = []string{
	"Finished On",
	"Description",
	"Location",
}

// NewExperienceParser creates a new experience parser for the given CSV file.
func NewExperienceParser(csvPath string) (*ExperienceParser, error) {
	reader, err := NewCSVReader(csvPath)
	if err != nil {
		return nil, fmt.Errorf("failed to create CSV reader: %w", err)
	}

	// Validate required columns
	missing := reader.ValidateColumns(experienceRequiredColumns)
	if len(missing) > 0 {
		return nil, fmt.Errorf("missing required columns: %v", missing)
	}

	return &ExperienceParser{
		csvPath: csvPath,
		reader:  reader,
	}, nil
}

// NewExperienceParserFromReader creates a parser from an io.Reader (for testing).
func NewExperienceParserFromReader(r io.Reader) (*ExperienceParser, error) {
	reader, err := NewCSVReaderFromIO(r)
	if err != nil {
		return nil, fmt.Errorf("failed to create CSV reader: %w", err)
	}

	return &ExperienceParser{
		reader: reader,
	}, nil
}

// ParseAll parses all experiences from the CSV file.
func (p *ExperienceParser) ParseAll() ([]models.Experience, error) {
	var experiences []models.Experience
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

		experience, err := p.parseRow(row, lineNum)
		if err != nil {
			return nil, err
		}

		experiences = append(experiences, *experience)
	}

	return experiences, nil
}

// parseRow converts a CSV row to an Experience struct.
func (p *ExperienceParser) parseRow(row map[string]string, lineNum int) (*models.Experience, error) {
	experience := models.Experience{
		Company:   row["Company Name"],
		Role:      row["Title"],
		StartDate: ConvertDate(row["Started On"]),
		EndDate:   ConvertDate(row["Finished On"]),
		Location:  row["Location"],
	}

	// Split description into bullets
	description := row["Description"]
	if description != "" {
		experience.Description = transformer.SplitDescription(description)
	}

	// Validate the experience
	if err := experience.Validate(); err != nil {
		return nil, fmt.Errorf("validation error at line %d: %w", lineNum, err)
	}

	return &experience, nil
}

// Close closes the parser and its underlying resources.
func (p *ExperienceParser) Close() error {
	return p.reader.Close()
}

// Validate validates the experiences CSV file without parsing all entries.
// Returns a slice of validation errors found.
func (p *ExperienceParser) Validate() []error {
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
		if row["Company Name"] == "" {
			errors = append(errors, models.NewParseError(lineNum, "Company Name", "cannot be empty"))
		}
		if row["Title"] == "" {
			errors = append(errors, models.NewParseError(lineNum, "Title", "cannot be empty"))
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

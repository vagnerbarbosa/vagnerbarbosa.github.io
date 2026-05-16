// Package parser handles CSV parsing for LinkedIn export files.
package parser

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
)

// CSVReader wraps a CSV reader with additional functionality.
type CSVReader struct {
	reader *csv.Reader
	header map[string]int
}

// NewCSVReader creates a new CSV reader from a file path.
func NewCSVReader(filepath string) (*CSVReader, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to open CSV file: %w", err)
	}
	return NewCSVReaderFromIO(file)
}

// NewCSVReaderFromIO creates a new CSV reader from an io.Reader.
func NewCSVReaderFromIO(r io.Reader) (*CSVReader, error) {
	reader := csv.NewReader(r)
	reader.FieldsPerRecord = 0 // Auto-detect number of fields
	reader.LazyQuotes = true   // Handle inconsistent quoting

	// Read header
	headerRow, err := reader.Read()
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV header: %w", err)
	}

	// Create header map
	header := make(map[string]int, len(headerRow))
	for i, col := range headerRow {
		header[col] = i
	}

	return &CSVReader{
		reader: reader,
		header: header,
	}, nil
}

// Next reads the next row from the CSV.
// Returns io.EOF when done.
func (r *CSVReader) Next() (map[string]string, error) {
	record, err := r.reader.Read()
	if err != nil {
		if !errors.Is(err, csv.ErrFieldCount) {
			return nil, err
		}
	}

	// Convert to map
	row := make(map[string]string, len(r.header))
	for col, idx := range r.header {
		if idx < len(record) {
			row[col] = record[idx]
		} else {
			row[col] = ""
		}
	}

	return row, nil
}

// GetHeader returns the column header map.
func (r *CSVReader) GetHeader() map[string]int {
	return r.header
}

// HasColumn checks if a column exists in the CSV.
func (r *CSVReader) HasColumn(name string) bool {
	_, exists := r.header[name]
	return exists
}

// Close closes the underlying reader if it's a Closer.
func (r *CSVReader) Close() error {
	// No-op for now, file closing handled by caller
	return nil
}

// ValidateColumns checks if all required columns are present.
func (r *CSVReader) ValidateColumns(required []string) []string {
	var missing []string
	for _, col := range required {
		if !r.HasColumn(col) {
			missing = append(missing, col)
		}
	}
	return missing
}

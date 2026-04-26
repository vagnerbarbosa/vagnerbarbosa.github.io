// Package models defines the data structures for LinkedIn import entities.
package models

import "fmt"

// ValidationError represents a field validation error.
type ValidationError struct {
	Field   string
	Message string
}

// Error implements the error interface.
func (e ValidationError) Error() string {
	return fmt.Sprintf("validation error for field '%s': %s", e.Field, e.Message)
}

// NewValidationError creates a new validation error.
func NewValidationError(field, message string) error {
	return &ValidationError{
		Field:   field,
		Message: message,
	}
}

// ParseError represents a CSV parsing error.
type ParseError struct {
	Line    int
	Column  string
	Message string
}

// Error implements the error interface.
func (e ParseError) Error() string {
	return fmt.Sprintf("parse error at line %d, column '%s': %s", e.Line, e.Column, e.Message)
}

// NewParseError creates a new parse error.
func NewParseError(line int, column, message string) error {
	return &ParseError{
		Line:    line,
		Column:  column,
		Message: message,
	}
}

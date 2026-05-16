package parser

import (
	"testing"
)

func TestConvertDate(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty", "", ""},
		{"present lowercase", "present", "Presente"},
		{"present uppercase", "PRESENT", "Presente"},
		{"present mixed case", "Present", "Presente"},
		{"present with spaces", "  Present  ", "Presente"},
		{"single date Jan", "Jan 2020", "Jan 2020"},
		{"single date Feb", "Feb 2020", "Fev 2020"},
		{"single date Apr", "Apr 2020", "Abr 2020"},
		{"single date Aug", "Aug 2020", "Ago 2020"},
		{"single date Sep", "Sep 2020", "Set 2020"},
		{"single date Oct", "Oct 2020", "Out 2020"},
		{"single date Dec", "Dec 2020", "Dez 2020"},
		{"year only", "2020", "2020"},
		{"range Jan-Mar", "Jan 2020 - Mar 2022", "Jan 2020 - Mar 2022"},
		{"range Feb-Apr", "Feb 2020 - Apr 2022", "Fev 2020 - Abr 2022"},
		{"range year-year", "2020 - 2022", "2020 - 2022"},
		{"range Present", "Jan 2020 - Present", "Jan 2020 - Presente"},
		{"invalid format", "Invalid Date", "Invalid Date"},
		{"wrong month length", "January 2020", "January 2020"},
		{"wrong year length", "Jan 20", "Jan 20"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertDate(tt.input); got != tt.expected {
				t.Errorf("ConvertDate(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}

func TestValidateDate(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		expected bool
	}{
		{"empty", "", true},
		{"present", "Present", true},
		{"presente", "Presente", true},
		{"single date", "Jan 2020", true},
		{"year only", "2020", true},
		{"range valid", "Jan 2020 - Feb 2021", true},
		{"range present", "Jan 2020 - Present", true},
		{"range invalid", "Jan 2020 - Invalid", false},
		{"invalid date", "Invalid Date", false},
		{"invalid month", "Ja 2020", false},
		{"invalid year", "Jan 20", false},
		{"random text", "hello world", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateDate(tt.input); got != tt.expected {
				t.Errorf("ValidateDate(%q) = %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}

func TestParseDateRange(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		expectStart string
		expectEnd   string
		expectErr   bool
	}{
		{"empty", "", "", "", false},
		{"single date", "Jan 2020", "Jan 2020", "", false},
		{"valid range", "Jan 2020 - Feb 2021", "Jan 2020", "Feb 2021", false},
		{"range with spaces", "  Jan 2020  -  Feb 2021  ", "Jan 2020", "Feb 2021", false},
		{"range with present", "Jan 2020 - Present", "Jan 2020", "Present", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start, end, err := ParseDateRange(tt.input)
			if (err != nil) != tt.expectErr {
				t.Errorf("ParseDateRange(%q) error = %v, expectErr %v", tt.input, err, tt.expectErr)
			}
			if start != tt.expectStart {
				t.Errorf("ParseDateRange(%q) start = %q, want %q", tt.input, start, tt.expectStart)
			}
			if end != tt.expectEnd {
				t.Errorf("ParseDateRange(%q) end = %q, want %q", tt.input, end, tt.expectEnd)
			}
		})
	}
}

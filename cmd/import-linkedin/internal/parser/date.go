// Package parser handles CSV parsing for LinkedIn export files.
package parser

import (
	"fmt"
	"regexp"
	"strings"
)

// monthMap maps English month abbreviations to Portuguese.
var monthMap = map[string]string{
	"Jan": "Jan",
	"Feb": "Fev",
	"Mar": "Mar",
	"Apr": "Abr",
	"May": "Mai",
	"Jun": "Jun",
	"Jul": "Jul",
	"Aug": "Ago",
	"Sep": "Set",
	"Oct": "Out",
	"Nov": "Nov",
	"Dec": "Dez",
}

// ConvertDate converts a LinkedIn date format to Portuguese.
// Supported formats:
//   - "Jan 2020" -> "Jan 2020"
//   - "Present" -> "Presente"
//   - "Jan 2020 - Mar 2022" -> "Jan 2020 - Mar 2022"
func ConvertDate(dateStr string) string {
	if dateStr == "" {
		return ""
	}

	dateStr = strings.TrimSpace(dateStr)

	// Handle "Present"
	if strings.EqualFold(dateStr, "Present") {
		return "Presente"
	}

	// Handle date ranges (e.g., "Jan 2020 - Mar 2022")
	if strings.Contains(dateStr, "-") {
		parts := strings.SplitN(dateStr, "-", 2)
		if len(parts) == 2 {
			start := ConvertDate(strings.TrimSpace(parts[0]))
			end := ConvertDate(strings.TrimSpace(parts[1]))
			return start + " - " + end
		}
	}

	// Convert single date (e.g., "Jan 2020")
	return convertSingleDate(dateStr)
}

// convertSingleDate converts a single date from English to Portuguese.
func convertSingleDate(dateStr string) string {
	dateStr = strings.TrimSpace(dateStr)

	// Regex to match "MMM YYYY" format
	re := regexp.MustCompile(`^([A-Za-z]{3})\s+(\d{4})$`)
	matches := re.FindStringSubmatch(dateStr)

	if len(matches) == 3 {
		month := matches[1]
		year := matches[2]

		if ptMonth, exists := monthMap[month]; exists {
			return fmt.Sprintf("%s %s", ptMonth, year)
		}
	}

	// Regex to match year-only format (e.g., "2010")
	reYear := regexp.MustCompile(`^(\d{4})$`)
	matches = reYear.FindStringSubmatch(dateStr)
	if len(matches) == 2 {
		// Year-only format, return as-is
		return dateStr
	}

	// If not matching expected format, return as-is
	return dateStr
}

// ValidateDate checks if a date string is in a valid format.
func ValidateDate(dateStr string) bool {
	if dateStr == "" {
		return true // Empty dates are valid (optional fields)
	}

	dateStr = strings.TrimSpace(dateStr)

	// Check for "Presente" or "Present"
	if strings.EqualFold(dateStr, "Present") || strings.EqualFold(dateStr, "Presente") {
		return true
	}

	// Check for "MMM YYYY" format (e.g., "Jan 2020")
	re := regexp.MustCompile(`^[A-Za-z]{3}\s+\d{4}$`)
	if re.MatchString(dateStr) {
		return true
	}

	// Check for year-only format (e.g., "2010")
	reYear := regexp.MustCompile(`^\d{4}$`)
	if reYear.MatchString(dateStr) {
		return true
	}

	// Check for date range format
	if strings.Contains(dateStr, "-") {
		parts := strings.SplitN(dateStr, "-", 2)
		if len(parts) == 2 {
			return ValidateDate(strings.TrimSpace(parts[0])) &&
				ValidateDate(strings.TrimSpace(parts[1]))
		}
	}

	return false
}

// ParseDateRange splits a date range into start and end dates.
// Returns (start, end, error).
func ParseDateRange(dateStr string) (string, string, error) {
	dateStr = strings.TrimSpace(dateStr)

	if dateStr == "" {
		return "", "", nil
	}

	// Handle date ranges
	if strings.Contains(dateStr, "-") {
		parts := strings.SplitN(dateStr, "-", 2)
		if len(parts) == 2 {
			start := strings.TrimSpace(parts[0])
			end := strings.TrimSpace(parts[1])
			return start, end, nil
		}
	}

	// Single date is treated as start date with no end
	return dateStr, "", nil
}

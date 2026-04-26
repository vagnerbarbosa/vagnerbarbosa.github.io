// Package transformer handles data transformations for LinkedIn import.
package transformer

import (
	"regexp"
	"strings"
)

// SplitDescription splits a description text into bullet points.
// Uses heuristics to identify logical breaks:
// 1. Double newlines (\n\n) - paragraphs
// 2. Period followed by capital letter - sentence boundaries
// 3. Bullet characters (*, -, •) at line start
func SplitDescription(text string) []string {
	if text == "" {
		return []string{}
	}

	text = strings.TrimSpace(text)

	// Try splitting by double newlines first (paragraphs)
	bullets := splitByParagraphs(text)
	if len(bullets) > 1 {
		return cleanBullets(bullets)
	}

	// Try splitting by bullet characters
	bullets = splitByBullets(text)
	if len(bullets) > 1 {
		return cleanBullets(bullets)
	}

	// Try splitting by sentences (period + space + capital)
	bullets = splitBySentences(text)
	if len(bullets) > 1 {
		return cleanBullets(bullets)
	}

	// If no splits worked, return as single bullet
	return cleanBullets([]string{text})
}

// splitByParagraphs splits text by double newlines.
func splitByParagraphs(text string) []string {
	// Normalize line endings
	text = strings.ReplaceAll(text, "\r\n", "\n")
	// Split by double newlines
	parts := strings.Split(text, "\n\n")
	return parts
}

// splitByBullets splits text by common bullet markers.
func splitByBullets(text string) []string {
	// Match bullet markers at start of line or after newline
	// Supports: *, -, •, ◦, ○, ●, 1., 2., etc.
	re := regexp.MustCompile(`(?:\n|\r\n|^)\s*[*\-•◦○●]\s+|\n\d+\.\s*`)
	parts := re.Split(text, -1)

	// Filter empty parts
	var result []string
	for _, p := range parts {
		if strings.TrimSpace(p) != "" {
			result = append(result, p)
		}
	}

	return result
}

// splitBySentences splits text by sentence boundaries.
// Looks for period followed by space and capital letter.
func splitBySentences(text string) []string {
	// Match period followed by space and capital letter
	re := regexp.MustCompile(`(?m)\.(?:\s+)([A-Z])`)

	// Replace with special marker
	marked := re.ReplaceAllString(text, ".\n$1")

	// Split by newlines
	parts := strings.Split(marked, "\n")

	return parts
}

// cleanBullets trims whitespace from each bullet and removes empty ones.
func cleanBullets(bullets []string) []string {
	var result []string
	for _, b := range bullets {
		b = strings.TrimSpace(b)
		// Remove trailing period if present
		b = strings.TrimSuffix(b, ".")
		if b != "" {
			result = append(result, b)
		}
	}
	return result
}

// JoinBullets joins bullet points back into a single text.
// Used for comparison or display purposes.
func JoinBullets(bullets []string, separator string) string {
	if separator == "" {
		separator = "\n\n"
	}
	return strings.Join(bullets, separator)
}

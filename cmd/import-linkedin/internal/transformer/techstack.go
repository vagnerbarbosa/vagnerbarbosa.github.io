// Package transformer handles data transformations for LinkedIn import.
package transformer

import (
	"regexp"
	"strings"
)

// TechStackResult represents the result of tech stack extraction.
type TechStackResult struct {
	CleanedBullets []string // Bullets without the tech stack
	TechStack      string   // Formatted tech stack or empty
	Found          bool     // Indicates if tech stack was found
	PatternMatched string   // Which pattern was detected (for debugging)
}

// Tech stack patterns to detect (case-insensitive)
// These patterns indicate that a bullet contains tech stack information
var techStackPatterns = []string{
	`(?i)as principais tecnologias e ferramentas utilizadas[:\s]*`,
	`(?i)tecnologias[:\s]*`,
	`(?i)tech stack[:\s]*`,
	`(?i)technologies[:\s]*`,
	`(?i)tech[:\s]*`,
	`(?i)stack[:\s]*`,
	`(?i)ferramentas[:\s]*`,
	`(?i)tools[:\s]*`,
}

// Separators used between technologies
var techSeparators = []string{
	",", ";", "|", "-", "•", "◦", "○", "●", "·",
}

// ExtractTechStack scans bullets for tech stack patterns and extracts technologies.
// It returns the cleaned bullets (without tech stack bullet) and the formatted tech stack.
// If multiple bullets contain patterns, only the last one is used.
func ExtractTechStack(bullets []string) TechStackResult {
	if len(bullets) == 0 {
		return TechStackResult{
			CleanedBullets: []string{},
			TechStack:      "",
			Found:          false,
		}
	}

	// Find the last bullet that matches a tech stack pattern
	lastMatchIndex := -1
	var matchedPattern string
	var extractedTech string

	for i, bullet := range bullets {
		for _, pattern := range techStackPatterns {
			re := regexp.MustCompile(pattern)
			if loc := re.FindStringIndex(bullet); loc != nil {
				// Extract text after the pattern
				afterPattern := bullet[loc[1]:]
				extractedTech = strings.TrimSpace(afterPattern)
				matchedPattern = pattern
				lastMatchIndex = i
				break
			}
		}
	}

	// If no pattern found, return original bullets
	if lastMatchIndex == -1 {
		return TechStackResult{
			CleanedBullets: bullets,
			TechStack:      "",
			Found:          false,
		}
	}

	// Parse technologies and format them
	techStack := parseTechStack(extractedTech)

	// Remove the tech stack bullet from the list
	cleanedBullets := make([]string, 0, len(bullets)-1)
	for i, bullet := range bullets {
		if i != lastMatchIndex {
			cleanedBullets = append(cleanedBullets, bullet)
		}
	}

	return TechStackResult{
		CleanedBullets: cleanedBullets,
		TechStack:      techStack,
		Found:          true,
		PatternMatched: matchedPattern,
	}
}

// parseTechStack parses a tech stack string with various separators
// and returns a formatted string with " • " as separator.
func parseTechStack(text string) string {
	if text == "" {
		return ""
	}

	// First, normalize all separators to a common one (comma)
	normalized := text
	for _, sep := range techSeparators {
		// Don't replace the first separator if it's a hyphen at the start
		if sep == "-" {
			// Replace " - " (space hyphen space) but not at the start
			normalized = strings.ReplaceAll(normalized, " - ", ", ")
			// Replace "-" preceded by space or start
			normalized = regexp.MustCompile(`(^|\s)-\s*`).ReplaceAllString(normalized, "$1, ")
		} else {
			normalized = strings.ReplaceAll(normalized, sep, ",")
		}
	}

	// Split by comma and clean up
	parts := strings.Split(normalized, ",")
	var technologies []string
	for _, part := range parts {
		tech := strings.TrimSpace(part)
		// Remove trailing period if present
		tech = strings.TrimSuffix(tech, ".")
		if tech != "" {
			technologies = append(technologies, tech)
		}
	}

	// Join with " • "
	return strings.Join(technologies, " • ")
}

// Package ui handles user interface components for the LinkedIn import CLI.
package ui

import (
	"fmt"
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/vagnerbarbosa/vagnerbarbosa.github.io/cmd/import-linkedin/internal/comparator"
	"github.com/vagnerbarbosa/vagnerbarbosa.github.io/cmd/import-linkedin/internal/models"
)

// Styles for diff visualization
var (
	// Colors
	addedStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("42"))  // Green
	removedStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("196")) // Red
	modifiedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("220")) // Yellow
	titleStyle    = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("39"))
	summaryStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("245"))
	borderStyle   = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Padding(1)
)

// DiffRenderer handles rendering of diff output.
type DiffRenderer struct {
	useColors bool
}

// NewDiffRenderer creates a new diff renderer.
func NewDiffRenderer(useColors bool) *DiffRenderer {
	return &DiffRenderer{useColors: useColors}
}

// RenderSummary renders a summary of the diff.
func (r *DiffRenderer) RenderSummary(diff *comparator.Diff) string {
	var sb strings.Builder

	sb.WriteString(titleStyle.Render("📊 Resumo de Mudanças"))
	sb.WriteString("\n")
	sb.WriteString(strings.Repeat("─", 40))
	sb.WriteString("\n")

	// Added
	added := diff.CountAdded()
	if r.useColors {
		sb.WriteString(addedStyle.Render(fmt.Sprintf("➕ Novas: %d entradas", added)))
	} else {
		sb.WriteString(fmt.Sprintf("+ Novas: %d entradas", added))
	}
	sb.WriteString("\n")

	// Modified
	modified := diff.CountModified()
	if r.useColors {
		sb.WriteString(modifiedStyle.Render(fmt.Sprintf("✏️  Modificadas: %d entradas", modified)))
	} else {
		sb.WriteString(fmt.Sprintf("~ Modificadas: %d entradas", modified))
	}
	sb.WriteString("\n")

	// Removed
	removed := diff.CountRemoved()
	if r.useColors {
		sb.WriteString(removedStyle.Render(fmt.Sprintf("➖ Removidas: %d entradas", removed)))
	} else {
		sb.WriteString(fmt.Sprintf("- Removidas: %d entradas", removed))
	}
	sb.WriteString("\n")

	return sb.String()
}

// RenderExperiencesDiff renders the experiences diff section.
func (r *DiffRenderer) RenderExperiencesDiff(diff comparator.EntityDiff[models.Experience]) string {
	var sb strings.Builder

	sb.WriteString(titleStyle.Render("💼 Experiências"))
	sb.WriteString("\n\n")

	// Added
	if len(diff.Added) > 0 {
		sb.WriteString(r.formatSectionHeader("Adicionadas", len(diff.Added), "+"))
		for _, exp := range diff.Added {
			sb.WriteString(r.formatAddedItem(fmt.Sprintf("%s @ %s", exp.Title, exp.Company)))
		}
		sb.WriteString("\n")
	}

	// Modified
	if len(diff.Modified) > 0 {
		sb.WriteString(r.formatSectionHeader("Modificadas", len(diff.Modified), "~"))
		for _, pair := range diff.Modified {
			sb.WriteString(r.formatModifiedItem(fmt.Sprintf("%s @ %s", pair.New.Title, pair.New.Company)))
			// Show changed fields
			fields := comparator.GetChangedFields(pair.Old, pair.New)
			if len(fields) > 0 {
				sb.WriteString(fmt.Sprintf("    Campos alterados: %s\n", strings.Join(fields, ", ")))
			}
		}
		sb.WriteString("\n")
	}

	// Removed
	if len(diff.Removed) > 0 {
		sb.WriteString(r.formatSectionHeader("Removidas", len(diff.Removed), "-"))
		for _, exp := range diff.Removed {
			sb.WriteString(r.formatRemovedItem(fmt.Sprintf("%s @ %s", exp.Title, exp.Company)))
		}
		sb.WriteString("\n")
	}

	if len(diff.Added) == 0 && len(diff.Modified) == 0 && len(diff.Removed) == 0 {
		sb.WriteString(summaryStyle.Render("  Nenhuma mudança\n"))
	}

	return sb.String()
}

// RenderEducationDiff renders the education diff section.
func (r *DiffRenderer) RenderEducationDiff(diff comparator.EntityDiff[models.Education]) string {
	var sb strings.Builder

	sb.WriteString(titleStyle.Render("🎓 Educação"))
	sb.WriteString("\n\n")

	// Added
	if len(diff.Added) > 0 {
		sb.WriteString(r.formatSectionHeader("Adicionadas", len(diff.Added), "+"))
		for _, edu := range diff.Added {
			sb.WriteString(r.formatAddedItem(fmt.Sprintf("%s at %s", edu.Degree, edu.Institution)))
		}
		sb.WriteString("\n")
	}

	// Modified
	if len(diff.Modified) > 0 {
		sb.WriteString(r.formatSectionHeader("Modificadas", len(diff.Modified), "~"))
		for _, pair := range diff.Modified {
			sb.WriteString(r.formatModifiedItem(fmt.Sprintf("%s at %s", pair.New.Degree, pair.New.Institution)))
		}
		sb.WriteString("\n")
	}

	// Removed
	if len(diff.Removed) > 0 {
		sb.WriteString(r.formatSectionHeader("Removidas", len(diff.Removed), "-"))
		for _, edu := range diff.Removed {
			sb.WriteString(r.formatRemovedItem(fmt.Sprintf("%s at %s", edu.Degree, edu.Institution)))
		}
		sb.WriteString("\n")
	}

	if len(diff.Added) == 0 && len(diff.Modified) == 0 && len(diff.Removed) == 0 {
		sb.WriteString(summaryStyle.Render("  Nenhuma mudança\n"))
	}

	return sb.String()
}

// RenderCertificationsDiff renders the certifications diff section.
func (r *DiffRenderer) RenderCertificationsDiff(diff comparator.EntityDiff[models.Certification]) string {
	var sb strings.Builder

	sb.WriteString(titleStyle.Render("📜 Certificações"))
	sb.WriteString("\n\n")

	// Added
	if len(diff.Added) > 0 {
		sb.WriteString(r.formatSectionHeader("Adicionadas", len(diff.Added), "+"))
		for _, cert := range diff.Added {
			sb.WriteString(r.formatAddedItem(fmt.Sprintf("%s (%s)", cert.Name, cert.Organization)))
		}
		sb.WriteString("\n")
	}

	// Modified
	if len(diff.Modified) > 0 {
		sb.WriteString(r.formatSectionHeader("Modificadas", len(diff.Modified), "~"))
		for _, pair := range diff.Modified {
			sb.WriteString(r.formatModifiedItem(fmt.Sprintf("%s (%s)", pair.New.Name, pair.New.Organization)))
		}
		sb.WriteString("\n")
	}

	// Removed
	if len(diff.Removed) > 0 {
		sb.WriteString(r.formatSectionHeader("Removidas", len(diff.Removed), "-"))
		for _, cert := range diff.Removed {
			sb.WriteString(r.formatRemovedItem(fmt.Sprintf("%s (%s)", cert.Name, cert.Organization)))
		}
		sb.WriteString("\n")
	}

	if len(diff.Added) == 0 && len(diff.Modified) == 0 && len(diff.Removed) == 0 {
		sb.WriteString(summaryStyle.Render("  Nenhuma mudança\n"))
	}

	return sb.String()
}

// RenderFullDiff renders the complete diff output.
func (r *DiffRenderer) RenderFullDiff(diff *comparator.Diff) string {
	var sb strings.Builder

	sb.WriteString(r.RenderSummary(diff))
	sb.WriteString("\n")

	if len(diff.Experiences.Added) > 0 || len(diff.Experiences.Modified) > 0 || len(diff.Experiences.Removed) > 0 {
		sb.WriteString(r.RenderExperiencesDiff(diff.Experiences))
		sb.WriteString("\n")
	}

	if len(diff.Education.Added) > 0 || len(diff.Education.Modified) > 0 || len(diff.Education.Removed) > 0 {
		sb.WriteString(r.RenderEducationDiff(diff.Education))
		sb.WriteString("\n")
	}

	if len(diff.Certifications.Added) > 0 || len(diff.Certifications.Modified) > 0 || len(diff.Certifications.Removed) > 0 {
		sb.WriteString(r.RenderCertificationsDiff(diff.Certifications))
		sb.WriteString("\n")
	}

	return sb.String()
}

// Helper methods
func (r *DiffRenderer) formatSectionHeader(title string, count int, symbol string) string {
	if r.useColors {
		switch symbol {
		case "+":
			return addedStyle.Render(fmt.Sprintf("%s %s (%d):\n", symbol, title, count))
		case "-":
			return removedStyle.Render(fmt.Sprintf("%s %s (%d):\n", symbol, title, count))
		default:
			return modifiedStyle.Render(fmt.Sprintf("%s %s (%d):\n", symbol, title, count))
		}
	}
	return fmt.Sprintf("%s %s (%d):\n", symbol, title, count)
}

func (r *DiffRenderer) formatAddedItem(item string) string {
	if r.useColors {
		return addedStyle.Render(fmt.Sprintf("  + %s\n", item))
	}
	return fmt.Sprintf("  + %s\n", item)
}

func (r *DiffRenderer) formatRemovedItem(item string) string {
	if r.useColors {
		return removedStyle.Render(fmt.Sprintf("  - %s\n", item))
	}
	return fmt.Sprintf("  - %s\n", item)
}

func (r *DiffRenderer) formatModifiedItem(item string) string {
	if r.useColors {
		return modifiedStyle.Render(fmt.Sprintf("  ~ %s\n", item))
	}
	return fmt.Sprintf("  ~ %s\n", item)
}

// PrintDiff prints the diff to stdout.
func PrintDiff(diff *comparator.Diff, useColors bool) {
	renderer := NewDiffRenderer(useColors)
	fmt.Print(renderer.RenderFullDiff(diff))
}

// PrintSummary prints just the summary to stdout.
func PrintSummary(diff *comparator.Diff, useColors bool) {
	renderer := NewDiffRenderer(useColors)
	fmt.Print(renderer.RenderSummary(diff))
}

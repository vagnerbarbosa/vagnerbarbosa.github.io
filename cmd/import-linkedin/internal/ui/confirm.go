// Package ui handles user interface components for the LinkedIn import CLI.
package ui

import (
	"fmt"

	"charm.land/huh/v2"
	"github.com/vagnerbarbosa/vagnerbarbosa.github.io/cmd/import-linkedin/internal/comparator"
	"github.com/vagnerbarbosa/vagnerbarbosa.github.io/cmd/import-linkedin/internal/models"
)

// ConfirmationResult represents the user's decision.
type ConfirmationResult int

const (
	// ConfirmYes means accept all changes.
	ConfirmYes ConfirmationResult = iota
	// ConfirmNo means reject all changes.
	ConfirmNo
	// ConfirmSelect means user wants to select specific changes.
	ConfirmSelect
	// ConfirmCancel means cancel the operation.
	ConfirmCancel
)

// Confirmation represents a user's choice for a specific change.
type Confirmation struct {
	Change  models.Change
	Accept  bool
}

// ConfirmAll asks the user to confirm all changes at once.
func ConfirmAll(diff *comparator.Diff) (ConfirmationResult, error) {
	if !diff.HasChanges() {
		fmt.Println("✓ Nenhuma mudança detectada. Config.yaml já está atualizado.")
		return ConfirmYes, nil
	}

	// Show summary
	PrintSummary(diff, true)
	fmt.Println()

	var result ConfirmationResult

	options := []huh.Option[ConfirmationResult]{
		huh.NewOption("Sim, aplicar todas as mudanças", ConfirmYes),
		huh.NewOption("Não, rejeitar todas", ConfirmNo),
		huh.NewOption("Selecionar mudanças específicas", ConfirmSelect),
		huh.NewOption("Cancelar", ConfirmCancel),
	}

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[ConfirmationResult]().
				Title("Deseja aplicar estas mudanças?").
				Description("Escolha como deseja prosseguir").
				Options(options...).
				Value(&result),
		),
	)

	if err := form.Run(); err != nil {
		return ConfirmCancel, err
	}

	return result, nil
}

// ConfirmSimple asks a simple yes/no question.
func ConfirmSimple(question string) (bool, error) {
	var confirmed bool

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title(question).
				Affirmative("Sim").
				Negative("Não").
				Value(&confirmed),
		),
	)

	if err := form.Run(); err != nil {
		return false, err
	}

	return confirmed, nil
}

// ConfirmWithSelect allows the user to select specific changes.
func ConfirmWithSelect(diff *comparator.Diff) ([]Confirmation, error) {
	var confirmations []Confirmation

	// Confirm experiences
	if len(diff.Experiences.Added) > 0 {
		fmt.Println("\n💼 Experiências - Adicionar:")
		for _, exp := range diff.Experiences.Added {
			confirmed, err := ConfirmSimple(fmt.Sprintf("Adicionar: %s @ %s?", exp.Title, exp.Company))
			if err != nil {
				return nil, err
			}
			confirmations = append(confirmations, Confirmation{
				Change: models.Change{
					EntityType: "experience",
					EntityID:   exp.ID(),
					NewValue:   exp,
					ChangeType: string(models.ChangeTypeAdded),
				},
				Accept: confirmed,
			})
		}
	}

	if len(diff.Experiences.Modified) > 0 {
		fmt.Println("\n💼 Experiências - Modificar:")
		for _, pair := range diff.Experiences.Modified {
			confirmed, err := ConfirmSimple(fmt.Sprintf("Modificar: %s @ %s?", pair.New.Title, pair.New.Company))
			if err != nil {
				return nil, err
			}
			confirmations = append(confirmations, Confirmation{
				Change: models.Change{
					EntityType: "experience",
					EntityID:   pair.New.ID(),
					OldValue:   pair.Old,
					NewValue:   pair.New,
					ChangeType: string(models.ChangeTypeModified),
				},
				Accept: confirmed,
			})
		}
	}

	// Confirm education
	if len(diff.Education.Added) > 0 {
		fmt.Println("\n🎓 Educação - Adicionar:")
		for _, edu := range diff.Education.Added {
			confirmed, err := ConfirmSimple(fmt.Sprintf("Adicionar: %s at %s?", edu.Degree, edu.Institution))
			if err != nil {
				return nil, err
			}
			confirmations = append(confirmations, Confirmation{
				Change: models.Change{
					EntityType: "education",
					EntityID:   edu.ID(),
					NewValue:   edu,
					ChangeType: string(models.ChangeTypeAdded),
				},
				Accept: confirmed,
			})
		}
	}

	if len(diff.Education.Modified) > 0 {
		fmt.Println("\n🎓 Educação - Modificar:")
		for _, pair := range diff.Education.Modified {
			confirmed, err := ConfirmSimple(fmt.Sprintf("Modificar: %s at %s?", pair.New.Degree, pair.New.Institution))
			if err != nil {
				return nil, err
			}
			confirmations = append(confirmations, Confirmation{
				Change: models.Change{
					EntityType: "education",
					EntityID:   pair.New.ID(),
					OldValue:   pair.Old,
					NewValue:   pair.New,
					ChangeType: string(models.ChangeTypeModified),
				},
				Accept: confirmed,
			})
		}
	}

	// Confirm certifications
	if len(diff.Certifications.Added) > 0 {
		fmt.Println("\n📜 Certificações - Adicionar:")
		for _, cert := range diff.Certifications.Added {
			confirmed, err := ConfirmSimple(fmt.Sprintf("Adicionar: %s (%s)?", cert.Name, cert.Organization))
			if err != nil {
				return nil, err
			}
			confirmations = append(confirmations, Confirmation{
				Change: models.Change{
					EntityType: "certification",
					EntityID:   cert.ID(),
					NewValue:   cert,
					ChangeType: string(models.ChangeTypeAdded),
				},
				Accept: confirmed,
			})
		}
	}

	if len(diff.Certifications.Modified) > 0 {
		fmt.Println("\n📜 Certificações - Modificar:")
		for _, pair := range diff.Certifications.Modified {
			confirmed, err := ConfirmSimple(fmt.Sprintf("Modificar: %s (%s)?", pair.New.Name, pair.New.Organization))
			if err != nil {
				return nil, err
			}
			confirmations = append(confirmations, Confirmation{
				Change: models.Change{
					EntityType: "certification",
					EntityID:   pair.New.ID(),
					OldValue:   pair.Old,
					NewValue:   pair.New,
					ChangeType: string(models.ChangeTypeModified),
				},
				Accept: confirmed,
			})
		}
	}

	return confirmations, nil
}

// ApplyConfirmations applies the confirmed changes to the config.
func ApplyConfirmations(config *models.ConfigPortfolio, confirmations []Confirmation) {
	var (
		experiences    []models.Experience
		education      []models.Education
		certifications []models.Certification
	)

	for _, conf := range confirmations {
		if !conf.Accept {
			continue
		}

		switch conf.Change.EntityType {
		case "experience":
			if exp, ok := conf.Change.NewValue.(models.Experience); ok {
				experiences = append(experiences, exp)
			}
		case "education":
			if edu, ok := conf.Change.NewValue.(models.Education); ok {
				education = append(education, edu)
			}
		case "certification":
			if cert, ok := conf.Change.NewValue.(models.Certification); ok {
				certifications = append(certifications, cert)
			}
		}
	}

	if len(experiences) > 0 {
		config.Content.Experiences = experiences
	}
	if len(education) > 0 {
		config.Content.Education = education
	}
	if len(certifications) > 0 {
		config.Content.Certifications = certifications
	}
}

// PrintConfirmationSummary prints a summary of accepted/rejected changes.
func PrintConfirmationSummary(confirmations []Confirmation) {
	var accepted, rejected int
	for _, conf := range confirmations {
		if conf.Accept {
			accepted++
		} else {
			rejected++
		}
	}

	fmt.Printf("\n✓ %d mudanças aceitas, %d rejeitadas\n", accepted, rejected)
}

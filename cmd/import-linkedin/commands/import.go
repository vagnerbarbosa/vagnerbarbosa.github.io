// Package commands implements the CLI commands for LinkedIn import.
package commands

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/vagnerbarbosa/vagnerbarbosa.github.io/cmd/import-linkedin/internal/comparator"
	"github.com/vagnerbarbosa/vagnerbarbosa.github.io/cmd/import-linkedin/internal/config"
	"github.com/vagnerbarbosa/vagnerbarbosa.github.io/cmd/import-linkedin/internal/models"
	"github.com/vagnerbarbosa/vagnerbarbosa.github.io/cmd/import-linkedin/internal/parser"
	"github.com/vagnerbarbosa/vagnerbarbosa.github.io/cmd/import-linkedin/internal/ui"
)

func RunImport(args []string) error {
	// Parse flags
	fs := flag.NewFlagSet("import", flag.ExitOnError)
	fs.StringVar(&Config.ExperiencesPath, "experiences", Config.ExperiencesPath, "Path to Experiences.csv")
	fs.StringVar(&Config.ExperiencesPath, "e", Config.ExperiencesPath, "Path to Experiences.csv (shorthand)")
	fs.StringVar(&Config.EducationPath, "education", Config.EducationPath, "Path to Education.csv")
	fs.StringVar(&Config.EducationPath, "E", Config.EducationPath, "Path to Education.csv (shorthand)")
	fs.StringVar(&Config.CertificationsPath, "certifications", Config.CertificationsPath, "Path to Certifications.csv")
	fs.StringVar(&Config.CertificationsPath, "c", Config.CertificationsPath, "Path to Certifications.csv (shorthand)")
	fs.StringVar(&Config.ConfigPath, "config", Config.ConfigPath, "Path to config.yaml")
	fs.StringVar(&Config.ConfigPath, "C", Config.ConfigPath, "Path to config.yaml (shorthand)")
	fs.BoolVar(&Config.DryRun, "dry-run", Config.DryRun, "Show diff without applying changes")
	fs.BoolVar(&Config.DryRun, "d", Config.DryRun, "Show diff without applying changes (shorthand)")
	fs.BoolVar(&Config.Yes, "yes", Config.Yes, "Accept all changes without confirmation")
	fs.BoolVar(&Config.Yes, "y", Config.Yes, "Accept all changes without confirmation (shorthand)")
	fs.BoolVar(&Config.Backup, "backup", Config.Backup, "Create backup before modifying")
	fs.BoolVar(&Config.Backup, "b", Config.Backup, "Create backup before modifying (shorthand)")

	if err := fs.Parse(args); err != nil {
		return fmt.Errorf("failed to parse flags: %w", err)
	}

	// Show banner
	fmt.Println("📊 Análise de Dados do LinkedIn")
	fmt.Println(strings.Repeat("═", 40))
	fmt.Println()

	// Parse CSV files
	var (
		experiences    []models.Experience
		education      []models.Education
		certifications []models.Certification
		parseErrors    []error
	)

	// Parse experiences
	if Config.ExperiencesPath != "" {
		if _, err := os.Stat(Config.ExperiencesPath); err == nil {
			expParser, err := parser.NewExperienceParser(Config.ExperiencesPath)
			if err != nil {
				parseErrors = append(parseErrors, fmt.Errorf("experiences: %w", err))
			} else {
				experiences, err = expParser.ParseAll()
				if err != nil {
					parseErrors = append(parseErrors, fmt.Errorf("experiences: %w", err))
				}
				fmt.Printf("✓ Experiências: %d encontradas\n", len(experiences))
			}
		} else {
			fmt.Printf("⚠ Experiências: arquivo não encontrado (%s)\n", Config.ExperiencesPath)
		}
	}

	// Parse education
	if Config.EducationPath != "" {
		if _, err := os.Stat(Config.EducationPath); err == nil {
			eduParser, err := parser.NewEducationParser(Config.EducationPath)
			if err != nil {
				parseErrors = append(parseErrors, fmt.Errorf("education: %w", err))
			} else {
				education, err = eduParser.ParseAll()
				if err != nil {
					parseErrors = append(parseErrors, fmt.Errorf("education: %w", err))
				}
				fmt.Printf("✓ Educação: %d encontradas\n", len(education))
			}
		} else {
			fmt.Printf("⚠ Educação: arquivo não encontrado (%s)\n", Config.EducationPath)
		}
	}

	// Parse certifications
	if Config.CertificationsPath != "" {
		if _, err := os.Stat(Config.CertificationsPath); err == nil {
			certParser, err := parser.NewCertificationParser(Config.CertificationsPath)
			if err != nil {
				parseErrors = append(parseErrors, fmt.Errorf("certifications: %w", err))
			} else {
				certifications, err = certParser.ParseAll()
				if err != nil {
					parseErrors = append(parseErrors, fmt.Errorf("certifications: %w", err))
				}
				fmt.Printf("✓ Certificações: %d encontradas\n", len(certifications))
			}
		} else {
			fmt.Printf("⚠ Certificações: arquivo não encontrado (%s)\n", Config.CertificationsPath)
		}
	}

	if len(parseErrors) > 0 {
		fmt.Println()
		fmt.Println("⚠ Erros encontrados durante o parsing:")
		for _, err := range parseErrors {
			fmt.Printf("  - %v\n", err)
		}
	}

	fmt.Println()

	// Read current config
	var currentConfig *models.ConfigPortfolio
	if config.FileExists(Config.ConfigPath) {
		var err error
		currentConfig, err = config.ReadYAML(Config.ConfigPath)
		if err != nil {
			fmt.Printf("⚠ Criando novo config.yaml (arquivo atual não pôde ser lido: %v)\n", err)
			currentConfig = models.NewConfigPortfolio()
		}
	} else {
		fmt.Println("⚠ Criando novo config.yaml (arquivo não existe)")
		currentConfig = models.NewConfigPortfolio()
	}

	// Compare data
	fmt.Println("📋 Comparando com config.yaml atual...")
	fmt.Println()

	diff := comparator.CompareAll(
		experiences, currentConfig.Content.Experiences,
		education, currentConfig.Content.Education,
		certifications, currentConfig.Content.Certifications,
	)

	// Show diff
	ui.PrintDiff(diff, true)

	// If dry-run, exit here
	if Config.DryRun {
		fmt.Println("\n📋 Modo dry-run: nenhuma alteração foi aplicada.")
		return nil
	}

	// If no changes, exit
	if !diff.HasChanges() {
		fmt.Println("\n✓ Nenhuma mudança detectada. Config.yaml já está atualizado.")
		return nil
	}

	// Confirmation
	var confirmed bool
	if Config.Yes {
		confirmed = true
		fmt.Println("\n✓ Modo automático: aceitando todas as mudanças.")
	} else {
		result, err := ui.ConfirmAll(diff)
		if err != nil {
			return fmt.Errorf("confirmation failed: %w", err)
		}

		switch result {
		case ui.ConfirmYes:
			confirmed = true
		case ui.ConfirmNo:
			fmt.Println("\n✓ Mudanças rejeitadas. Nenhuma alteração foi aplicada.")
			return nil
		case ui.ConfirmSelect:
			confirmations, err := ui.ConfirmWithSelect(diff)
			if err != nil {
				return fmt.Errorf("selective confirmation failed: %w", err)
			}
			ui.ApplyConfirmations(currentConfig, confirmations)
			ui.PrintConfirmationSummary(confirmations)
			confirmed = false
		case ui.ConfirmCancel:
			fmt.Println("\n✓ Operação cancelada.")
			return nil
		}
	}

	// Apply changes
	if confirmed {
		if err := comparator.ApplyChanges(diff, currentConfig); err != nil {
			return fmt.Errorf("failed to apply changes: %w", err)
		}
	}

	// Create backup if enabled
	if Config.Backup && config.FileExists(Config.ConfigPath) {
		backupPath, err := config.CreateDefaultBackup(Config.ConfigPath)
		if err != nil {
			fmt.Printf("⚠ Falha ao criar backup: %v\n", err)
		} else {
			fmt.Printf("✓ Backup criado: %s\n", backupPath)
		}
	}

	// Write updated config
	if err := config.WriteYAML(Config.ConfigPath, currentConfig); err != nil {
		return fmt.Errorf("failed to write config: %w", err)
	}

	// Print summary
	result := comparator.GetImportResult(diff)
	comparator.PrintImportResult(result)
	fmt.Printf("\n✓ Configuração salva em: %s\n", Config.ConfigPath)

	return nil
}

// Package commands implements the CLI commands for LinkedIn import.
package commands

import (
	"flag"
	"fmt"
	"os"

	"github.com/vagnerbarbosa/vagnerbarbosa.github.io/cmd/import-linkedin/internal/parser"
)

func runValidate(args []string) error {
	// Parse flags
	fs := flag.NewFlagSet("validate", flag.ExitOnError)
	fs.StringVar(&Config.ExperiencesPath, "experiences", Config.ExperiencesPath, "Path to Experiences.csv")
	fs.StringVar(&Config.ExperiencesPath, "e", Config.ExperiencesPath, "Path to Experiences.csv (shorthand)")
	fs.StringVar(&Config.EducationPath, "education", Config.EducationPath, "Path to Education.csv")
	fs.StringVar(&Config.EducationPath, "E", Config.EducationPath, "Path to Education.csv (shorthand)")
	fs.StringVar(&Config.CertificationsPath, "certifications", Config.CertificationsPath, "Path to Certifications.csv")
	fs.StringVar(&Config.CertificationsPath, "c", Config.CertificationsPath, "Path to Certifications.csv (shorthand)")

	if err := fs.Parse(args); err != nil {
		return fmt.Errorf("failed to parse flags: %w", err)
	}

	fmt.Println("🔍 Validando arquivos CSV...")
	fmt.Println()

	var hasErrors bool

	// Validate experiences
	if Config.ExperiencesPath != "" {
		if _, err := os.Stat(Config.ExperiencesPath); err == nil {
			expParser, err := parser.NewExperienceParser(Config.ExperiencesPath)
			if err != nil {
				fmt.Printf("✗ Experiences.csv: %v\n", err)
				hasErrors = true
			} else {
				errors := expParser.Validate()
				if len(errors) > 0 {
					fmt.Printf("✗ Experiences.csv:\n")
					for _, e := range errors {
						fmt.Printf("  - %v\n", e)
					}
					hasErrors = true
				} else {
					// Parse to count entries
					expParser2, _ := parser.NewExperienceParser(Config.ExperiencesPath)
					experiences, _ := expParser2.ParseAll()
					fmt.Printf("✓ Experiences.csv: %d entradas válidas\n", len(experiences))
				}
			}
		} else {
			fmt.Printf("⚠ Experiences.csv: arquivo não encontrado\n")
		}
	}

	// Validate education
	if Config.EducationPath != "" {
		if _, err := os.Stat(Config.EducationPath); err == nil {
			eduParser, err := parser.NewEducationParser(Config.EducationPath)
			if err != nil {
				fmt.Printf("✗ Education.csv: %v\n", err)
				hasErrors = true
			} else {
				errors := eduParser.Validate()
				if len(errors) > 0 {
					fmt.Printf("✗ Education.csv:\n")
					for _, e := range errors {
						fmt.Printf("  - %v\n", e)
					}
					hasErrors = true
				} else {
					// Parse to count entries
					eduParser2, _ := parser.NewEducationParser(Config.EducationPath)
					education, _ := eduParser2.ParseAll()
					fmt.Printf("✓ Education.csv: %d entradas válidas\n", len(education))
				}
			}
		} else {
			fmt.Printf("⚠ Education.csv: arquivo não encontrado\n")
		}
	}

	// Validate certifications
	if Config.CertificationsPath != "" {
		if _, err := os.Stat(Config.CertificationsPath); err == nil {
			certParser, err := parser.NewCertificationParser(Config.CertificationsPath)
			if err != nil {
				fmt.Printf("✗ Certifications.csv: %v\n", err)
				hasErrors = true
			} else {
				errors := certParser.Validate()
				if len(errors) > 0 {
					fmt.Printf("✗ Certifications.csv:\n")
					for _, e := range errors {
						fmt.Printf("  - %v\n", e)
					}
					hasErrors = true
				} else {
					// Parse to count entries
					certParser2, _ := parser.NewCertificationParser(Config.CertificationsPath)
					certs, _ := certParser2.ParseAll()
					fmt.Printf("✓ Certifications.csv: %d entradas válidas\n", len(certs))
				}
			}
		} else {
			fmt.Printf("⚠ Certifications.csv: arquivo não encontrado\n")
		}
	}

	fmt.Println()

	if hasErrors {
		return fmt.Errorf("validação falhou")
	}

	fmt.Println("✓ Todos os arquivos são válidos!")
	return nil
}

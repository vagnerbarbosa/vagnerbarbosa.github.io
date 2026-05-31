// Package commands implements the CLI commands for LinkedIn import.
package commands

import (
	"fmt"
	"os"
)

// Exit codes
const (
	ExitSuccess           = 0
	ExitErrorGeneric      = 1
	ExitErrorCSV          = 2
	ExitErrorYAML         = 3
	ExitErrorNotFound     = 4
	ExitErrorPermission   = 5
	ExitErrorInterrupted  = 130
)

// Config holds CLI configuration.
var Config = struct {
	ExperiencesPath    string
	EducationPath      string
	CertificationsPath string
	ConfigPath         string
	DryRun             bool
	Yes                bool
	Backup             bool
}{
	ExperiencesPath:    "Experiences.csv",
	EducationPath:      "Education.csv",
	CertificationsPath: "Certifications.csv",
	ConfigPath:         "config.yaml",
	DryRun:             false,
	Yes:                false,
	Backup:             true,
}

// Command is the interface for all CLI commands.
type Command interface {
	Name() string
	Description() string
	Run(args []string) error
}

// Execute runs the appropriate command based on arguments.
func Execute(args []string) int {
	if len(args) == 0 {
		args = append(args, "import")
	}

	cmd := args[0]

	switch cmd {
	case "import":
		if err := RunImport(args[1:]); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			return ExitErrorGeneric
		}
		return ExitSuccess
	case "validate":
		if err := runValidate(args[1:]); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			return ExitErrorCSV
		}
		return ExitSuccess
	case "version":
		printVersion()
		return ExitSuccess
	case "help", "--help", "-h":
		printHelp()
		return ExitSuccess
	default:
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", cmd)
		printHelp()
		return ExitErrorGeneric
	}
}

// printVersion prints the version information.
func printVersion() {
	fmt.Println("linkedin-import version 1.0.0")
	fmt.Println("Go version: go1.26.2")
}

// printHelp prints the help message.
func printHelp() {
	fmt.Println(`LinkedIn Import CLI

Usage:
  linkedin-import <command> [flags]

Commands:
  import     Import data from LinkedIn CSV files
  validate   Validate CSV files without importing
  version    Show version information
  help       Show this help message

Use "linkedin-import <command> --help" for more information about a command.`)
}

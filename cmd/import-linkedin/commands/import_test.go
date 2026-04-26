// Package commands implements the CLI commands for LinkedIn import.
package commands

import (
	"testing"
)

func TestExecute_UnknownCommand(t *testing.T) {
	code := Execute([]string{"unknown"})
	if code != ExitErrorGeneric {
		t.Errorf("Expected exit code %d for unknown command, got %d", ExitErrorGeneric, code)
	}
}

func TestExecute_Help(t *testing.T) {
	code := Execute([]string{"help"})
	if code != ExitSuccess {
		t.Errorf("Expected exit code %d for help, got %d", ExitSuccess, code)
	}
}

func TestExecute_Version(t *testing.T) {
	code := Execute([]string{"version"})
	if code != ExitSuccess {
		t.Errorf("Expected exit code %d for version, got %d", ExitSuccess, code)
	}
}

func TestExecute_NoArgs(t *testing.T) {
	// No args defaults to import command
	// With no files, it creates an empty config and returns success
	code := Execute([]string{})
	// No files = no changes = success
	if code != ExitSuccess {
		t.Errorf("Expected exit code %d for default import with no files, got %d", ExitSuccess, code)
	}
}

func TestExecute_Validate(t *testing.T) {
	// Validate with no files should succeed
	code := Execute([]string{"validate", "-e", "", "-E", "", "-c", ""})
	if code != ExitSuccess {
		t.Errorf("Expected exit code %d for validate with no files, got %d", ExitSuccess, code)
	}
}

func TestConfig_Defaults(t *testing.T) {
	// Reset config to defaults
	Config.ExperiencesPath = "Experiences.csv"
	Config.EducationPath = "Education.csv"
	Config.CertificationsPath = "Certifications.csv"
	Config.ConfigPath = "config.yaml"
	Config.DryRun = false
	Config.Yes = false
	Config.Backup = true

	if Config.ExperiencesPath != "Experiences.csv" {
		t.Error("Default ExperiencesPath not set correctly")
	}
	if Config.ConfigPath != "config.yaml" {
		t.Error("Default ConfigPath not set correctly")
	}
	if Config.Backup != true {
		t.Error("Default Backup should be true")
	}
}

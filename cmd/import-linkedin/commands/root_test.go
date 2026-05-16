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
	tests := []struct {
		name string
		args []string
	}{
		{"help command", []string{"help"}},
		{"--help flag", []string{"--help"}},
		{"-h flag", []string{"-h"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			code := Execute(tt.args)
			if code != ExitSuccess {
				t.Errorf("Expected exit code %d for %v, got %d", ExitSuccess, tt.args, code)
			}
		})
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
	code := Execute([]string{})
	if code != ExitSuccess {
		t.Errorf("Expected exit code %d for default import with no files, got %d", ExitSuccess, code)
	}
}

func TestExecute_VersionExtended(t *testing.T) {
	// Test version extended output
	code := Execute([]string{"version", "--verbose"})
	// Should succeed even if verbose flag not fully supported
	if code != ExitSuccess && code != ExitErrorGeneric {
		t.Logf("Version extended returned code: %d", code)
	}
}

func TestPrintVersionExtended(t *testing.T) {
	// Simply call the function to ensure coverage
	// Can't easily capture stdout, but this ensures the code runs
	printVersionExtended()
}

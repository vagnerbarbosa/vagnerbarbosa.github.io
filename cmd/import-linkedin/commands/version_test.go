package commands

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestPrintVersionExtended(t *testing.T) {
	// Capture stdout
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	printVersionExtended()

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	expectedMarkers := []string{
		"LinkedIn Import CLI",
		"Version:",
		"Build Date:",
		"Go version:",
		"linkedin-import help",
	}

	for _, marker := range expectedMarkers {
		if !strings.Contains(output, marker) {
			t.Errorf("Expected output to contain %q, got:\n%s", marker, output)
		}
	}
}

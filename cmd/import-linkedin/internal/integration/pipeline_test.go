//go:build integration

package integration

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/vagnerbarbosa/vagnerbarbosa.github.io/cmd/import-linkedin/commands"
)

func TestPipelineHappyPath(t *testing.T) {
	runPipelineTest(t, "happy_path", "experiences.csv")
}

func TestPipelinePartialData(t *testing.T) {
	runPipelineTest(t, "partial_data", "partial_data.csv")
}

func TestPipelineParsingError(t *testing.T) {
	runPipelineTest(t, "parsing_error", "parsing_error.csv")
}

func runPipelineTest(t *testing.T, goldenName, csvName string) {
	tmpDir := t.TempDir()
	configPath := filepathToYaml(tmpDir, "config.yaml")
	csvPath := filepathToCSV(csvName)
	goldenPath := filepathToGolden(goldenName + ".yaml")

	commands.Config.ConfigPath = configPath
	commands.Config.ExperiencesPath = csvPath
	commands.Config.EducationPath = ""
	commands.Config.CertificationsPath = ""
	commands.Config.Yes = true
	commands.Config.DryRun = false

	err := commands.RunImport([]string{})
	if err != nil {
		t.Fatalf("Import failed: %v", err)
	}

	currentOutput, err := os.ReadFile(configPath)
	if err != nil {
		if os.IsNotExist(err) && goldenName == "parsing_error" {
			// For parsing errors, we expect no config file to be created
			// Verify that the golden file is also empty or represents an empty config
			goldenOutput, err := os.ReadFile(goldenPath)
			if err != nil {
				t.Fatalf("failed to read golden file: %v", err)
			}
			if len(goldenOutput) > 0 {
				t.Errorf("Expected empty config for parsing error, but found content")
			}
			return
		}
		t.Fatalf("failed to read resulting config: %v", err)
	}

	goldenOutput, err := os.ReadFile(goldenPath)
	if err != nil {
		t.Fatalf("failed to read golden file: %v", err)
	}

	equal, err := CompareYAML(currentOutput, goldenOutput)
	if err != nil {
		t.Fatalf("Comparison error: %v", err)
	}

	if !equal {
		t.Errorf("Output does not match Golden File for %s\nActual: %s\nExpected: %s", goldenName, string(currentOutput), string(goldenOutput))
	}
}

func filepathToCSV(name string) string {
	cwd, _ := os.Getwd()
	return filepath.Join(cwd, "../../testdata/input_csv/", name)
}

func filepathToGolden(name string) string {
	cwd, _ := os.Getwd()
	return filepath.Join(cwd, "../../testdata/golden/", name)
}

func filepathToYaml(dir, name string) string {
	return filepath.Join(dir, name)
}

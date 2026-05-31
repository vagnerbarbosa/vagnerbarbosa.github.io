package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/vagnerbarbosa/vagnerbarbosa.github.io/cmd/import-linkedin/commands"
)

// E2ETestRunner orchestrates the full flow from import to HTML generation.
type E2ETestRunner struct {
	TempDir    string
	ConfigPath string
	PublicDir  string
}

func NewE2ETestRunner(t *testing.T) *E2ETestRunner {
	tmpDir := t.TempDir()

	// Use relative paths from the package directory (cmd/generator) to reach the root
	root := "../../"
	copyDir(filepath.Join(root, "templates"), filepath.Join(tmpDir, "templates"))
	copyDir(filepath.Join(root, "assets"), filepath.Join(tmpDir, "assets"))

	configPath := filepath.Join(tmpDir, "config.yaml")
	publicDir := filepath.Join(tmpDir, "public")

	return &E2ETestRunner{
		TempDir:    tmpDir,
		ConfigPath: configPath,
		PublicDir:  publicDir,
	}
}

func (r *E2ETestRunner) RunImport(expPath, eduPath, certPath string) error {
	commands.Config.ConfigPath = r.ConfigPath
	commands.Config.ExperiencesPath = expPath
	commands.Config.EducationPath = eduPath
	commands.Config.CertificationsPath = certPath
	commands.Config.Yes = true
	commands.Config.DryRun = false

	return commands.RunImport([]string{})
}

func (r *E2ETestRunner) RunGenerator() error {
	originalCwd, err := os.Getwd()
	if err != nil {
		return err
	}
	defer os.Chdir(originalCwd)

	if err := os.Chdir(r.TempDir); err != nil {
		return err
	}

	return Run()
}

func TestE2E_FullPipeline(t *testing.T) {
	runner := NewE2ETestRunner(t)

	// 1. Setup test data paths
	// The test runs in cmd/generator, so we go up two levels to reach the project root.
	root := "../../"
	e2eDataDir := filepath.Join(root, "cmd/import-linkedin/testdata/e2e")
	expPath := filepath.Join(e2eDataDir, "experiences.csv")
	eduPath := filepath.Join(e2eDataDir, "education.csv")
	certPath := filepath.Join(e2eDataDir, "certifications.csv")

	// 2. Measure total execution time
	start := time.Now()

	// 3. Run Import
	t.Log("Running import pipeline...")
	if err := runner.RunImport(expPath, eduPath, certPath); err != nil {
		t.Fatalf("Import failed: %v", err)
	}

	// 4. Run Generator
	t.Log("Running site generator...")
	if err := runner.RunGenerator(); err != nil {
		t.Fatalf("Generator failed: %v", err)
	}

	duration := time.Since(start)
	t.Logf("Full pipeline completed in %v", duration)

	if duration > 15*time.Second {
		t.Errorf("E2E pipeline took too long: %v (limit: 15s)", duration)
	}

	// 5. Validate Output
	indexPath := filepath.Join(runner.PublicDir, "index.html")
	if _, err := os.Stat(indexPath); os.IsNotExist(err) {
		t.Fatalf("index.html was not generated at %s", indexPath)
	}

	htmlContent, err := os.ReadFile(indexPath)
	if err != nil {
		t.Fatalf("failed to read index.html: %v", err)
	}
	content := string(htmlContent)

	// Validate markers for different sections
	testCases := []struct {
		name    string
		markers []string
	}{
		{"Experience", []string{"Senior Software Engineer", "TechCorp", "Fullstack Developer", "WebStart"}},
		{"Education", []string{"University of Technology", "B.S. Computer Science"}},
		{"Certifications", []string{"AWS Certified Solutions Architect", "Amazon Web Services"}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			for _, marker := range tc.markers {
				if !strings.Contains(content, marker) {
					t.Errorf("Expected marker %q not found in generated HTML", marker)
				}
			}
		})
	}

	// Validate SEO assets
	seoFiles := []string{"sitemap.xml", "robots.txt", "site.webmanifest"}
	for _, file := range seoFiles {
		path := filepath.Join(runner.PublicDir, file)
		if _, err := os.Stat(path); os.IsNotExist(err) {
			t.Errorf("SEO file %s was not generated", file)
		}
	}
}

func copyDir(src, dst string) error {
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		rel, _ := filepath.Rel(src, path)
		target := filepath.Join(dst, rel)
		if info.IsDir() {
			return os.MkdirAll(target, 0755)
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		return os.WriteFile(target, data, 0644)
	})
}

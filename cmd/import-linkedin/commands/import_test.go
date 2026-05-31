// Package commands implements the CLI commands for LinkedIn import.
package commands

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

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

func TestRunImport_MissingFiles(t *testing.T) {
	tmpDir := t.TempDir()

	// Save original config
	origExp := Config.ExperiencesPath
	origEdu := Config.EducationPath
	origCert := Config.CertificationsPath
	origCfg := Config.ConfigPath
	origDry := Config.DryRun
	origYes := Config.Yes
	origBkp := Config.Backup

	defer func() {
		Config.ExperiencesPath = origExp
		Config.EducationPath = origEdu
		Config.CertificationsPath = origCert
		Config.ConfigPath = origCfg
		Config.DryRun = origDry
		Config.Yes = origYes
		Config.Backup = origBkp
	}()

	// Set paths to non-existent files in tmpDir
	Config.ExperiencesPath = filepath.Join(tmpDir, "missing_exp.csv")
	Config.EducationPath = filepath.Join(tmpDir, "missing_edu.csv")
	Config.CertificationsPath = filepath.Join(tmpDir, "missing_cert.csv")
	Config.ConfigPath = filepath.Join(tmpDir, "config.yaml")
	Config.Yes = true // Avoid interactive prompt

	err := RunImport([]string{})
	if err != nil {
		t.Errorf("Expected no error when CSV files are missing, got: %v", err)
	}
}

func TestRunImport_DryRun(t *testing.T) {
	tmpDir := t.TempDir()

	// Save original config
	origExp := Config.ExperiencesPath
	origEdu := Config.EducationPath
	origCert := Config.CertificationsPath
	origCfg := Config.ConfigPath
	origDry := Config.DryRun
	origYes := Config.Yes
	origBkp := Config.Backup

	defer func() {
		Config.ExperiencesPath = origExp
		Config.EducationPath = origEdu
		Config.CertificationsPath = origCert
		Config.ConfigPath = origCfg
		Config.DryRun = origDry
		Config.Yes = origYes
		Config.Backup = origBkp
	}()

	configPath := filepath.Join(tmpDir, "config.yaml")
	initialContent := "initial content"
	err := os.WriteFile(configPath, []byte(initialContent), 0644)
	if err != nil {
		t.Fatal(err)
	}

	Config.ExperiencesPath = filepath.Join(tmpDir, "exp.csv")
	Config.EducationPath = filepath.Join(tmpDir, "edu.csv")
	Config.CertificationsPath = filepath.Join(tmpDir, "cert.csv")
	Config.ConfigPath = configPath
	Config.DryRun = true
	Config.Yes = true

	// Create empty CSV files to avoid "file not found" warnings,
	// but since they are empty, no changes will be detected.
	os.WriteFile(Config.ExperiencesPath, []byte(""), 0644)
	os.WriteFile(Config.EducationPath, []byte(""), 0644)
	os.WriteFile(Config.CertificationsPath, []byte(""), 0644)

	err = RunImport([]string{"--dry-run"})
	if err != nil {
		t.Errorf("Expected no error during dry-run, got: %v", err)
	}

	content, err := os.ReadFile(configPath)
	if err != nil {
		t.Fatal(err)
	}
	if string(content) != initialContent {
		t.Errorf("Expected config file to remain unchanged during dry-run, but it was modified")
	}
}

func TestRunImport_WithValidFiles(t *testing.T) {
	tmpDir := t.TempDir()

	// Save original config
	origExp := Config.ExperiencesPath
	origEdu := Config.EducationPath
	origCert := Config.CertificationsPath
	origCfg := Config.ConfigPath
	origDry := Config.DryRun
	origYes := Config.Yes
	origBkp := Config.Backup

	defer func() {
		Config.ExperiencesPath = origExp
		Config.EducationPath = origEdu
		Config.CertificationsPath = origCert
		Config.ConfigPath = origCfg
		Config.DryRun = origDry
		Config.Yes = origYes
		Config.Backup = origBkp
	}()

	// Create valid CSV files
	expPath := filepath.Join(tmpDir, "exp.csv")
	expContent := "Company Name,Title,Started On,Finished On,Location,Description\nAcme Corp,Engineer,Jan 2020,Mar 2022,NY,Work\n"
	os.WriteFile(expPath, []byte(expContent), 0644)

	eduPath := filepath.Join(tmpDir, "edu.csv")
	eduContent := "School Name,Degree Name,Field Of Study,Start Date,Finished On\nMIT,BSc,CS,Jan 2015,Jan 2019\n"
	os.WriteFile(eduPath, []byte(eduContent), 0644)

	certPath := filepath.Join(tmpDir, "cert.csv")
	certContent := "Name,Authority,Started On\nAWS Cert,Amazon,Jan 2020\n"
	os.WriteFile(certPath, []byte(certContent), 0644)

	configPath := filepath.Join(tmpDir, "config.yaml")
	os.WriteFile(configPath, []byte("site:\n  title: \"Test\"\ncontent:\n  experiences: []\n  education: []\n  certifications: []\n"), 0644)

	Config.ExperiencesPath = expPath
	Config.EducationPath = eduPath
	Config.CertificationsPath = certPath
	Config.ConfigPath = configPath
	Config.DryRun = false
	Config.Yes = true
	Config.Backup = true

	err := RunImport([]string{})
	if err != nil {
		t.Errorf("Expected no error with valid files, got: %v", err)
	}

	// Verify config.yaml was updated
	content, err := os.ReadFile(configPath)
	if err != nil {
		t.Fatalf("Failed to read config.yaml: %v", err)
	}
	if !strings.Contains(string(content), "Acme Corp") {
		t.Errorf("Expected config.yaml to contain 'Acme Corp', got:\n%s", string(content))
	}
}

func TestRunImport_WithBackup(t *testing.T) {
	tmpDir := t.TempDir()

	// Save original config
	origExp := Config.ExperiencesPath
	origEdu := Config.EducationPath
	origCert := Config.CertificationsPath
	origCfg := Config.ConfigPath
	origDry := Config.DryRun
	origYes := Config.Yes
	origBkp := Config.Backup

	defer func() {
		Config.ExperiencesPath = origExp
		Config.EducationPath = origEdu
		Config.CertificationsPath = origCert
		Config.ConfigPath = origCfg
		Config.DryRun = origDry
		Config.Yes = origYes
		Config.Backup = origBkp
	}()

	// Create existing config file
	configPath := filepath.Join(tmpDir, "config.yaml")
	os.WriteFile(configPath, []byte("existing: true\n"), 0644)

	// Create valid CSV file
	expPath := filepath.Join(tmpDir, "exp.csv")
	expContent := "Company Name,Title,Started On\nAcme Corp,Engineer,Jan 2020\nTech Inc,Dev,Mar 2021\n"
	os.WriteFile(expPath, []byte(expContent), 0644)

	Config.ExperiencesPath = expPath
	Config.EducationPath = ""
	Config.CertificationsPath = ""
	Config.ConfigPath = configPath
	Config.DryRun = false
	Config.Yes = true
	Config.Backup = true

	err := RunImport([]string{})
	if err != nil {
		t.Errorf("Expected no error with backup enabled, got: %v", err)
	}
}

func TestRunImport_ParseErrors(t *testing.T) {
	tmpDir := t.TempDir()

	// Save original config
	origExp := Config.ExperiencesPath
	origEdu := Config.EducationPath
	origCert := Config.CertificationsPath
	origCfg := Config.ConfigPath
	origDry := Config.DryRun
	origYes := Config.Yes

	defer func() {
		Config.ExperiencesPath = origExp
		Config.EducationPath = origEdu
		Config.CertificationsPath = origCert
		Config.ConfigPath = origCfg
		Config.DryRun = origDry
		Config.Yes = origYes
	}()

	// Create invalid CSV files (missing required columns)
	expPath := filepath.Join(tmpDir, "exp.csv")
	os.WriteFile(expPath, []byte("Invalid,Header\nValue1,Value2\n"), 0644)

	eduPath := filepath.Join(tmpDir, "edu.csv")
	os.WriteFile(eduPath, []byte("Invalid,Header\nValue1,Value2\n"), 0644)

	certPath := filepath.Join(tmpDir, "cert.csv")
	os.WriteFile(certPath, []byte("Invalid,Header\nValue1,Value2\n"), 0644)

	Config.ExperiencesPath = expPath
	Config.EducationPath = eduPath
	Config.CertificationsPath = certPath
	Config.ConfigPath = filepath.Join(tmpDir, "config.yaml")
	Config.DryRun = true
	Config.Yes = true

	// Should handle parse errors gracefully and return nil in dry-run mode
	err := RunImport([]string{})
	if err != nil {
		t.Errorf("Expected no error in dry-run mode even with parse errors, got: %v", err)
	}
}

func TestRunImport_InvalidYAML(t *testing.T) {
	tmpDir := t.TempDir()

	origCfg := Config.ConfigPath
	defer func() { Config.ConfigPath = origCfg }()

	configPath := filepath.Join(tmpDir, "config.yaml")
	os.WriteFile(configPath, []byte("invalid: yaml: :"), 0644)

	Config.ConfigPath = configPath
	Config.Yes = true

	err := RunImport([]string{})
	if err != nil {
		t.Errorf("Expected no error when YAML is invalid (should fallback to new config), got: %v", err)
	}
}

func TestRunImport_NoBackup(t *testing.T) {
	tmpDir := t.TempDir()

	origBkp := Config.Backup
	origCfg := Config.ConfigPath
	defer func() {
		Config.Backup = origBkp
		Config.ConfigPath = origCfg
	}()

	configPath := filepath.Join(tmpDir, "config.yaml")
	os.WriteFile(configPath, []byte("site:\n  title: \"Test\"\n"), 0644)

	Config.ConfigPath = configPath
	Config.Backup = false
	Config.Yes = true

	err := RunImport([]string{})
	if err != nil {
		t.Errorf("Expected no error with backup disabled, got: %v", err)
	}

	// Verify no backup file was created in tmpDir
	files, _ := os.ReadDir(tmpDir)
	for _, f := range files {
		if strings.Contains(f.Name(), ".bak") {
			t.Errorf("Backup file %s should not have been created", f.Name())
		}
	}
}

func TestRunImport_ShorthandFlags(t *testing.T) {
	tmpDir := t.TempDir()

	origExp := Config.ExperiencesPath
	origEdu := Config.EducationPath
	origCert := Config.CertificationsPath
	origCfg := Config.ConfigPath
	origDry := Config.DryRun
	origYes := Config.Yes
	origBkp := Config.Backup

	defer func() {
		Config.ExperiencesPath = origExp
		Config.EducationPath = origEdu
		Config.CertificationsPath = origCert
		Config.ConfigPath = origCfg
		Config.DryRun = origDry
		Config.Yes = origYes
		Config.Backup = origBkp
	}()

	expPath := filepath.Join(tmpDir, "exp.csv")
	eduPath := filepath.Join(tmpDir, "edu.csv")
	certPath := filepath.Join(tmpDir, "cert.csv")
	configPath := filepath.Join(tmpDir, "config.yaml")

	os.WriteFile(expPath, []byte(""), 0644)
	os.WriteFile(eduPath, []byte(""), 0644)
	os.WriteFile(certPath, []byte(""), 0644)
	os.WriteFile(configPath, []byte("site:\n  title: \"Test\"\n"), 0644)

	args := []string{"-e", expPath, "-E", eduPath, "-c", certPath, "-C", configPath, "-d", "-y", "-b"}
	err := RunImport(args)
	if err != nil {
		t.Errorf("Expected no error with shorthand flags, got: %v", err)
	}

	if Config.ExperiencesPath != expPath || Config.EducationPath != eduPath || Config.CertificationsPath != certPath || Config.ConfigPath != configPath {
		t.Errorf("Shorthand flags did not set Config paths correctly")
	}
	if !Config.DryRun || !Config.Yes || !Config.Backup {
		t.Errorf("Shorthand flags did not set boolean config correctly")
	}
}

func TestRunImport_CSVParseError(t *testing.T) {
	tmpDir := t.TempDir()

	origExp := Config.ExperiencesPath
	origEdu := Config.EducationPath
	origCert := Config.CertificationsPath
	origCfg := Config.ConfigPath
	defer func() {
		Config.ExperiencesPath = origExp
		Config.EducationPath = origEdu
		Config.CertificationsPath = origCert
		Config.ConfigPath = origCfg
	}()

	expPath := filepath.Join(tmpDir, "exp.csv")
	// Create a CSV with a malformed row (wrong number of fields)
	// Header has 6 fields, but second row has 1.
	os.WriteFile(expPath, []byte("Company Name,Title,Started On,Finished On,Location,Description\nMalformedRow\n"), 0644)

	Config.ExperiencesPath = expPath
	Config.EducationPath = ""
	Config.CertificationsPath = ""
	Config.ConfigPath = filepath.Join(tmpDir, "config.yaml")
	Config.Yes = true

	err := RunImport([]string{})
	if err != nil {
		t.Errorf("Expected no error in RunImport when CSV parsing fails (should report warning), got: %v", err)
	}
}

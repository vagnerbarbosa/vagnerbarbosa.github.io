package commands

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRunValidate_MalformedCSV(t *testing.T) {
	// Save original config
	origExp := Config.ExperiencesPath
	origEdu := Config.EducationPath
	origCert := Config.CertificationsPath

	defer func() {
		Config.ExperiencesPath = origExp
		Config.EducationPath = origEdu
		Config.CertificationsPath = origCert
	}()

	// Use a malformed CSV for experiences
	malformedPath := filepath.Join("..", "testdata", "malformed", "missing_columns.csv")
	Config.ExperiencesPath = malformedPath
	Config.EducationPath = "" // Skip education
	Config.CertificationsPath = "" // Skip certifications

	err := runValidate([]string{})
	if err == nil {
		t.Error("Expected validation to fail for malformed CSV, but it succeeded")
	}
}

func TestRunValidate_ValidCSV(t *testing.T) {
	tmpDir := t.TempDir()

	// Save original config
	origExp := Config.ExperiencesPath
	origEdu := Config.EducationPath
	origCert := Config.CertificationsPath

	defer func() {
		Config.ExperiencesPath = origExp
		Config.EducationPath = origEdu
		Config.CertificationsPath = origCert
	}()

	// Create a minimal valid CSV for experiences
	expPath := filepath.Join(tmpDir, "valid_exp.csv")
	content := "Company Name,Title,Started On,Finished On,Location,Description\nCompany A,Title A,Jan 2020,Jan 2021,City,Desc\n"
	if err := os.WriteFile(expPath, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	Config.ExperiencesPath = expPath
	Config.EducationPath = ""
	Config.CertificationsPath = ""

	err := runValidate([]string{})
	if err != nil {
		t.Errorf("Expected validation to succeed for valid CSV, got: %v", err)
	}
}

func TestRunValidate_WithEducation(t *testing.T) {
	tmpDir := t.TempDir()

	// Save original config
	origExp := Config.ExperiencesPath
	origEdu := Config.EducationPath
	origCert := Config.CertificationsPath

	defer func() {
		Config.ExperiencesPath = origExp
		Config.EducationPath = origEdu
		Config.CertificationsPath = origCert
	}()

	// Create valid CSV for education
	eduPath := filepath.Join(tmpDir, "edu.csv")
	content := "School Name,Degree Name,Field Of Study,Start Date\nMIT,BSc,CS,Jan 2020\n"
	if err := os.WriteFile(eduPath, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	Config.ExperiencesPath = ""
	Config.EducationPath = eduPath
	Config.CertificationsPath = ""

	err := runValidate([]string{})
	if err != nil {
		t.Errorf("Expected validation to succeed for valid education CSV, got: %v", err)
	}
}

func TestRunValidate_WithCertifications(t *testing.T) {
	tmpDir := t.TempDir()

	// Save original config
	origExp := Config.ExperiencesPath
	origEdu := Config.EducationPath
	origCert := Config.CertificationsPath

	defer func() {
		Config.ExperiencesPath = origExp
		Config.EducationPath = origEdu
		Config.CertificationsPath = origCert
	}()

	// Create valid CSV for certifications
	certPath := filepath.Join(tmpDir, "cert.csv")
	content := "Name,Authority,Started On\nAWS Cert,Amazon,Jan 2020\n"
	if err := os.WriteFile(certPath, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	Config.ExperiencesPath = ""
	Config.EducationPath = ""
	Config.CertificationsPath = certPath

	err := runValidate([]string{})
	if err != nil {
		t.Errorf("Expected validation to succeed for valid certification CSV, got: %v", err)
	}
}

func TestRunValidate_WithParseError(t *testing.T) {
	tmpDir := t.TempDir()

	// Save original config
	origExp := Config.ExperiencesPath
	origEdu := Config.EducationPath
	origCert := Config.CertificationsPath

	defer func() {
		Config.ExperiencesPath = origExp
		Config.EducationPath = origEdu
		Config.CertificationsPath = origCert
	}()

	// Create invalid CSV (missing required column)
	expPath := filepath.Join(tmpDir, "exp.csv")
	content := "Invalid,Header\nValue1,Value2\n"
	if err := os.WriteFile(expPath, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	Config.ExperiencesPath = expPath
	Config.EducationPath = ""
	Config.CertificationsPath = ""

	err := runValidate([]string{})
	if err == nil {
		t.Error("Expected validation to fail for CSV with parse error")
	}
}

func TestRunValidate_WithValidationErrors(t *testing.T) {
	tmpDir := t.TempDir()

	// Save original config
	origExp := Config.ExperiencesPath
	origEdu := Config.EducationPath
	origCert := Config.CertificationsPath

	defer func() {
		Config.ExperiencesPath = origExp
		Config.EducationPath = origEdu
		Config.CertificationsPath = origCert
	}()

	// Create CSV with validation errors (empty required fields)
	expPath := filepath.Join(tmpDir, "exp.csv")
	content := "Company Name,Title,Started On\n,Engineer,Jan 2020\n" // Missing Company Name
	if err := os.WriteFile(expPath, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	Config.ExperiencesPath = expPath
	Config.EducationPath = ""
	Config.CertificationsPath = ""

	err := runValidate([]string{})
	if err == nil {
		t.Error("Expected validation to fail for CSV with validation errors")
	}
}

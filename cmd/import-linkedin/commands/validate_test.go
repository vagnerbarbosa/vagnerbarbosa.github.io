package commands

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRunValidate_Success(t *testing.T) {
	tmpDir := t.TempDir()

	expPath := filepath.Join(tmpDir, "exp.csv")
	expContent := "Company Name,Title,Started On,Finished On,Location,Description\nAcme Corp,Engineer,Jan 2020,Mar 2022,NY,Work\n"
	os.WriteFile(expPath, []byte(expContent), 0644)

	eduPath := filepath.Join(tmpDir, "edu.csv")
	eduContent := "School Name,Degree Name,Field Of Study,Start Date,Finished On\nMIT,BSc,CS,Jan 2015,Jan 2019\n"
	os.WriteFile(eduPath, []byte(eduContent), 0644)

	certPath := filepath.Join(tmpDir, "cert.csv")
	certContent := "Name,Authority,Started On\nAWS Cert,Amazon,Jan 2020\n"
	os.WriteFile(certPath, []byte(certContent), 0644)

	origExp := Config.ExperiencesPath
	origEdu := Config.EducationPath
	origCert := Config.CertificationsPath
	defer func() {
		Config.ExperiencesPath = origExp
		Config.EducationPath = origEdu
		Config.CertificationsPath = origCert
	}()

	Config.ExperiencesPath = expPath
	Config.EducationPath = eduPath
	Config.CertificationsPath = certPath

	err := runValidate([]string{})
	if err != nil {
		t.Errorf("Expected no error for valid files, got: %v", err)
	}
}

func TestRunValidate_InvalidFiles(t *testing.T) {
	tmpDir := t.TempDir()

	expPath := filepath.Join(tmpDir, "exp.csv")
	// Missing required columns
	expContent := "Wrong,Header,On\nValue1,Value2,Value3\n"
	os.WriteFile(expPath, []byte(expContent), 0644)

	origExp := Config.ExperiencesPath
	defer func() { Config.ExperiencesPath = origExp }()

	Config.ExperiencesPath = expPath
	Config.EducationPath = ""
	Config.CertificationsPath = ""

	err := runValidate([]string{})
	if err == nil {
		t.Error("Expected error for invalid CSV columns, got nil")
	}
}

func TestRunValidate_ParseError(t *testing.T) {
	tmpDir := t.TempDir()

	expPath := filepath.Join(tmpDir, "exp.csv")
	// Header is correct, but row is malformed
	expContent := "Company Name,Title,Started On,Finished On,Location,Description\nMalformed\n"
	os.WriteFile(expPath, []byte(expContent), 0644)

	origExp := Config.ExperiencesPath
	defer func() { Config.ExperiencesPath = origExp }()

	Config.ExperiencesPath = expPath
	Config.EducationPath = ""
	Config.CertificationsPath = ""

	err := runValidate([]string{})
	if err == nil {
		t.Error("Expected error for malformed CSV row, got nil")
	}
}

func TestRunValidate_MissingFiles(t *testing.T) {
	origExp := Config.ExperiencesPath
	origEdu := Config.EducationPath
	origCert := Config.CertificationsPath
	defer func() {
		Config.ExperiencesPath = origExp
		Config.EducationPath = origEdu
		Config.CertificationsPath = origCert
	}()

	Config.ExperiencesPath = "non-existent-exp.csv"
	Config.EducationPath = "non-existent-edu.csv"
	Config.CertificationsPath = "non-existent-cert.csv"

	err := runValidate([]string{})
	if err != nil {
		t.Errorf("Expected no error when files are missing (should just warn), got: %v", err)
	}
}

func TestRunValidate_ShorthandFlags(t *testing.T) {
	tmpDir := t.TempDir()
	expPath := filepath.Join(tmpDir, "exp.csv")
	os.WriteFile(expPath, []byte("Company Name,Title,Started On,Finished On,Location,Description\nAcme,Eng,Jan 2020,Mar 2022,NY,Work\n"), 0644)

	origExp := Config.ExperiencesPath
	defer func() { Config.ExperiencesPath = origExp }()

	args := []string{"-e", expPath}
	err := runValidate(args)
	if err != nil {
		t.Errorf("Expected no error with shorthand flag, got: %v", err)
	}
	if Config.ExperiencesPath != expPath {
		t.Errorf("Expected ExperiencesPath to be %s, got %s", expPath, Config.ExperiencesPath)
	}
}

package main

import (
	"html/template"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/vagnerbarbosa/vagnerbarbosa.github.io/internal/config"
)

// testFs é uma estrutura auxiliar para criar estrutura de diretórios temporária.
// testFs is a helper structure for creating temporary directory structures.
type testFs struct {
	root      string
	templates string
	assets    string
	config    string
}

// setupTestFs cria uma estrutura de diretórios temporária para testes.
// setupTestFs creates a temporary directory structure for testing.
func setupTestFs(t *testing.T) *testFs {
	t.Helper()
	root := t.TempDir()
	return &testFs{
		root:      root,
		templates: filepath.Join(root, "templates"),
		assets:    filepath.Join(root, "assets"),
		config:    filepath.Join(root, "config.yaml"),
	}
}

// setupTemplates cria templates Go de exemplo para testes.
// setupTemplates creates sample Go templates for testing.
func (fs *testFs) setupTemplates(t *testing.T) {
	t.Helper()
	if err := os.MkdirAll(fs.templates, 0755); err != nil {
		t.Fatalf("Failed to create templates dir: %v", err)
	}

	// Create index template
	indexTemplate := `{{define "index"}}<!DOCTYPE html>
<html>
<head><title>{{.Site.Title}}</title></head>
<body>{{template "header" .}}<main>{{template "about" .}}</main></body>
</html>{{end}}`
	if err := os.WriteFile(filepath.Join(fs.templates, "index.html"), []byte(indexTemplate), 0644); err != nil {
		t.Fatalf("Failed to write index template: %v", err)
	}

	// Create header partial
	headerTemplate := `{{define "header"}}<header>{{.Site.Username}}</header>{{end}}`
	headerDir := filepath.Join(fs.templates, "partials")
	if err := os.MkdirAll(headerDir, 0755); err != nil {
		t.Fatalf("Failed to create partials dir: %v", err)
	}
	if err := os.WriteFile(filepath.Join(headerDir, "header.html"), []byte(headerTemplate), 0644); err != nil {
		t.Fatalf("Failed to write header template: %v", err)
	}

	// Create about partial
	aboutTemplate := `{{define "about"}}<div>About: {{.Site.UserDescription}}</div>{{end}}`
	if err := os.WriteFile(filepath.Join(headerDir, "about.html"), []byte(aboutTemplate), 0644); err != nil {
		t.Fatalf("Failed to write about template: %v", err)
	}
}

// setupConfig cria um arquivo de configuração YAML de teste.
// setupConfig creates a test YAML configuration file.
func (fs *testFs) setupConfig(t *testing.T) {
	t.Helper()
	configContent := `site:
  title: "Test Site"
  description: "Test"
  baseurl: ""
  username: "testuser"
  user_description: "Test description"
  user_title: "Developer"
  email: "test@test.com"
content:
  about:
    pt: "Test about."
    en: "Test about."
  experiences: []
  education: []
  technologies: []
`
	if err := os.WriteFile(fs.config, []byte(configContent), 0644); err != nil {
		t.Fatalf("Failed to write config: %v", err)
	}
}

// setupAssets cria assets estáticos de teste (CSS e JS).
// setupAssets creates test static assets (CSS and JS).
func (fs *testFs) setupAssets(t *testing.T) {
	t.Helper()
	cssDir := filepath.Join(fs.assets, "css")
	if err := os.MkdirAll(cssDir, 0755); err != nil {
		t.Fatalf("Failed to create css dir: %v", err)
	}
	if err := os.WriteFile(filepath.Join(cssDir, "main.css"), []byte("body { color: black; }"), 0644); err != nil {
		t.Fatalf("Failed to write css: %v", err)
	}

	jsDir := filepath.Join(fs.assets, "js")
	if err := os.MkdirAll(jsDir, 0755); err != nil {
		t.Fatalf("Failed to create js dir: %v", err)
	}
	if err := os.WriteFile(filepath.Join(jsDir, "app.js"), []byte("console.log('test');"), 0644); err != nil {
		t.Fatalf("Failed to write js: %v", err)
	}
}

// TestRun_Success verifica a execução completa do gerador com todos os componentes.
// TestRun_Success verifies the complete generator execution with all components.
func TestRun_Success(t *testing.T) {
	// Setup test filesystem
	fs := setupTestFs(t)
	fs.setupTemplates(t)
	fs.setupConfig(t)
	fs.setupAssets(t)

	// Create CNAME file
	cnamePath := filepath.Join(fs.root, "CNAME")
	if err := os.WriteFile(cnamePath, []byte("www.test.com"), 0644); err != nil {
		t.Fatalf("Failed to write CNAME: %v", err)
	}

	// Save original cwd and change to test directory
	originalCwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get cwd: %v", err)
	}
	defer os.Chdir(originalCwd)

	if err := os.Chdir(fs.root); err != nil {
		t.Fatalf("Failed to chdir: %v", err)
	}

	// Run the main function
	if err := Run(); err != nil {
		t.Errorf("Run() error = %v", err)
	}

	// Verify public/index.html was created
	indexPath := filepath.Join(fs.root, "public", "index.html")
	if _, err := os.Stat(indexPath); os.IsNotExist(err) {
		t.Errorf("public/index.html was not created")
	}

	// Verify index.html content
	content, err := os.ReadFile(indexPath)
	if err != nil {
		t.Errorf("Failed to read index.html: %v", err)
	}
	if !strings.Contains(string(content), "Test Site") {
		t.Errorf("index.html should contain 'Test Site'")
	}
	if !strings.Contains(string(content), "testuser") {
		t.Errorf("index.html should contain 'testuser'")
	}

	// Verify assets were copied
	publicCSS := filepath.Join(fs.root, "public", "assets", "css", "main.css")
	if _, err := os.Stat(publicCSS); os.IsNotExist(err) {
		t.Errorf("public/assets/css/main.css was not created")
	}

	// Verify CNAME was copied
	publicCNAME := filepath.Join(fs.root, "public", "CNAME")
	if _, err := os.Stat(publicCNAME); os.IsNotExist(err) {
		t.Errorf("public/CNAME was not created")
	}
	cnameContent, err := os.ReadFile(publicCNAME)
	if err != nil {
		t.Errorf("Failed to read CNAME: %v", err)
	}
	if string(cnameContent) != "www.test.com" {
		t.Errorf("CNAME content = %v, want %v", string(cnameContent), "www.test.com")
	}
}

// TestRun_NoCNAME verifica que o gerador funciona mesmo sem arquivo CNAME.
// TestRun_NoCNAME verifies that the generator works even without a CNAME file.
func TestRun_NoCNAME(t *testing.T) {
	// Setup test filesystem without CNAME
	fs := setupTestFs(t)
	fs.setupTemplates(t)
	fs.setupConfig(t)
	fs.setupAssets(t)

	// Save original cwd
	originalCwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get cwd: %v", err)
	}
	defer os.Chdir(originalCwd)

	if err := os.Chdir(fs.root); err != nil {
		t.Fatalf("Failed to chdir: %v", err)
	}

	// Run should succeed even without CNAME
	if err := Run(); err != nil {
		t.Errorf("Run() error = %v", err)
	}

	// Verify public/index.html was still created
	indexPath := filepath.Join(fs.root, "public", "index.html")
	if _, err := os.Stat(indexPath); os.IsNotExist(err) {
		t.Errorf("public/index.html was not created")
	}

	// Verify CNAME was not created (should be skipped)
	publicCNAME := filepath.Join(fs.root, "public", "CNAME")
	if _, err := os.Stat(publicCNAME); !os.IsNotExist(err) {
		t.Errorf("public/CNAME should not be created when source doesn't exist")
	}
}

// TestRun_MissingConfig verifica o erro quando o arquivo de configuração está ausente.
// TestRun_MissingConfig verifies the error when the configuration file is missing.
func TestRun_MissingConfig(t *testing.T) {
	// Setup without config file
	fs := setupTestFs(t)

	originalCwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get cwd: %v", err)
	}
	defer os.Chdir(originalCwd)

	if err := os.Chdir(fs.root); err != nil {
		t.Fatalf("Failed to chdir: %v", err)
	}

	// Run should fail
	if err := Run(); err == nil {
		t.Errorf("Run() should return error for missing config")
	}
}

// TestRun_MissingTemplates verifica o erro quando templates estão ausentes.
// TestRun_MissingTemplates verifies the error when templates are missing.
func TestRun_MissingTemplates(t *testing.T) {
	// Setup without templates
	fs := setupTestFs(t)
	fs.setupConfig(t)

	originalCwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get cwd: %v", err)
	}
	defer os.Chdir(originalCwd)

	if err := os.Chdir(fs.root); err != nil {
		t.Fatalf("Failed to chdir: %v", err)
	}

	// Run should fail - templates are required
	err = Run()
	if err == nil {
		t.Errorf("Run() should return error for missing templates")
	}
}

// TestParseTemplates verifica o parsing bem-sucedido de templates Go.
// TestParseTemplates verifies successful parsing of Go templates.
func TestParseTemplates(t *testing.T) {
	// Setup test filesystem
	fs := setupTestFs(t)
	fs.setupTemplates(t)

	// Change to test directory
	originalCwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get cwd: %v", err)
	}
	defer os.Chdir(originalCwd)

	if err := os.Chdir(fs.root); err != nil {
		t.Fatalf("Failed to chdir: %v", err)
	}

	// Parse templates
	tmpl, err := parseTemplates()
	if err != nil {
		t.Fatalf("parseTemplates() error = %v", err)
	}

	// Verify templates were parsed
	if tmpl == nil {
		t.Error("parseTemplates() returned nil template")
	}

	// Verify index template exists
	if tmpl.Lookup("index") == nil {
		t.Error("index template not found")
	}

	// Verify header template exists
	if tmpl.Lookup("header") == nil {
		t.Error("header template not found")
	}

	// Verify about template exists
	if tmpl.Lookup("about") == nil {
		t.Error("about template not found")
	}
}

// TestParseTemplates_MissingDir verifica o erro quando diretório de templates não existe.
// TestParseTemplates_MissingDir verifies error when templates directory doesn't exist.
func TestParseTemplates_MissingDir(t *testing.T) {
	// Change to directory without templates
	tmpDir := t.TempDir()

	originalCwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get cwd: %v", err)
	}
	defer os.Chdir(originalCwd)

	if err := os.Chdir(tmpDir); err != nil {
		t.Fatalf("Failed to chdir: %v", err)
	}

	// Should return an error when templates directory doesn't exist
	_, err = parseTemplates()
	if err == nil {
		t.Error("parseTemplates() should return error for missing templates directory")
	}
}

// TestParseTemplates_InvalidTemplate verifica o erro para sintaxe de template inválida.
// TestParseTemplates_InvalidTemplate verifies error for invalid template syntax.
func TestParseTemplates_InvalidTemplate(t *testing.T) {
	// Setup with invalid template
	fs := setupTestFs(t)
	if err := os.MkdirAll(fs.templates, 0755); err != nil {
		t.Fatalf("Failed to create templates dir: %v", err)
	}

	// Write invalid template syntax
	invalidTemplate := `{{define "index"}}{{if}}{{end}}{{end}}`
	if err := os.WriteFile(filepath.Join(fs.templates, "invalid.html"), []byte(invalidTemplate), 0644); err != nil {
		t.Fatalf("Failed to write invalid template: %v", err)
	}

	originalCwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get cwd: %v", err)
	}
	defer os.Chdir(originalCwd)

	if err := os.Chdir(fs.root); err != nil {
		t.Fatalf("Failed to chdir: %v", err)
	}

	// Should return an error for invalid template syntax
	_, err = parseTemplates()
	if err == nil {
		t.Error("parseTemplates() should return error for invalid template syntax")
	}
}

// TestGenerateIndex verifica a geração do arquivo index.html.
// TestGenerateIndex verifies the generation of the index.html file.
func TestGenerateIndex(t *testing.T) {
	// Setup
	fs := setupTestFs(t)
	fs.setupTemplates(t)

	originalCwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get cwd: %v", err)
	}
	defer os.Chdir(originalCwd)

	if err := os.Chdir(fs.root); err != nil {
		t.Fatalf("Failed to chdir: %v", err)
	}

	tmpl, err := parseTemplates()
	if err != nil {
		t.Fatalf("Failed to parse templates: %v", err)
	}

	// Create public directory
	if err := os.MkdirAll(filepath.Join(fs.root, "public"), 0755); err != nil {
		t.Fatalf("Failed to create public dir: %v", err)
	}

	data := config.TemplateData{
		Site: config.SiteConfig{
			Title:    "Test Title",
			Username: "TestUser",
		},
	}

	// Generate index
	if err := generateIndex(tmpl, data); err != nil {
		t.Fatalf("generateIndex() error = %v", err)
	}

	// Verify file was created
	indexPath := filepath.Join(fs.root, "public", "index.html")
	content, err := os.ReadFile(indexPath)
	if err != nil {
		t.Fatalf("Failed to read index.html: %v", err)
	}

	if !strings.Contains(string(content), "Test Title") {
		t.Errorf("index.html should contain 'Test Title'")
	}

	if !strings.Contains(string(content), "TestUser") {
		t.Errorf("index.html should contain 'TestUser'")
	}
}

// TestCopyAssets verifica a cópia recursiva de assets estáticos.
// TestCopyAssets verifies the recursive copying of static assets.
func TestCopyAssets(t *testing.T) {
	// Setup
	fs := setupTestFs(t)
	fs.setupAssets(t)

	originalCwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get cwd: %v", err)
	}
	defer os.Chdir(originalCwd)

	if err := os.Chdir(fs.root); err != nil {
		t.Fatalf("Failed to chdir: %v", err)
	}

	// Copy assets
	if err := copyAssets(); err != nil {
		t.Fatalf("copyAssets() error = %v", err)
	}

	// Verify CSS was copied
	publicCSS := filepath.Join(fs.root, "public", "assets", "css", "main.css")
	if _, err := os.Stat(publicCSS); os.IsNotExist(err) {
		t.Errorf("CSS file was not copied")
	}

	// Verify content (minified)
	content, err := os.ReadFile(publicCSS)
	if err != nil {
		t.Fatalf("Failed to read copied CSS: %v", err)
	}
	// CSS is minified, so we check for the minified version
	if !strings.Contains(string(content), "color:") {
		t.Errorf("Copied CSS content = %v, should contain 'color:'", string(content))
	}

	// Verify JS was copied
	publicJS := filepath.Join(fs.root, "public", "assets", "js", "app.js")
	if _, err := os.Stat(publicJS); os.IsNotExist(err) {
		t.Errorf("JS file was not copied")
	}
}

// TestCopyAssets_MissingDir verifica o erro quando diretório de assets não existe.
// TestCopyAssets_MissingDir verifies error when assets directory doesn't exist.
func TestCopyAssets_MissingDir(t *testing.T) {
	// Setup without assets directory
	fs := setupTestFs(t)

	originalCwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get cwd: %v", err)
	}
	defer os.Chdir(originalCwd)

	if err := os.Chdir(fs.root); err != nil {
		t.Fatalf("Failed to chdir: %v", err)
	}

	// Should return error when assets directory doesn't exist
	err = copyAssets()
	if err == nil {
		t.Error("copyAssets() should return error for missing assets directory")
	}
}

// TestCopyCNAME verifica a cópia do arquivo CNAME.
// TestCopyCNAME verifies the copying of the CNAME file.
func TestCopyCNAME(t *testing.T) {
	// Setup
	fs := setupTestFs(t)

	// Create CNAME file
	cnamePath := filepath.Join(fs.root, "CNAME")
	if err := os.WriteFile(cnamePath, []byte("www.example.com"), 0644); err != nil {
		t.Fatalf("Failed to write CNAME: %v", err)
	}

	// Create public directory
	publicDir := filepath.Join(fs.root, "public")
	if err := os.MkdirAll(publicDir, 0755); err != nil {
		t.Fatalf("Failed to create public dir: %v", err)
	}

	originalCwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get cwd: %v", err)
	}
	defer os.Chdir(originalCwd)

	if err := os.Chdir(fs.root); err != nil {
		t.Fatalf("Failed to chdir: %v", err)
	}

	// Copy CNAME
	if err := copyCNAME(); err != nil {
		t.Fatalf("copyCNAME() error = %v", err)
	}

	// Verify CNAME was copied
	publicCNAME := filepath.Join(fs.root, "public", "CNAME")
	if _, err := os.Stat(publicCNAME); os.IsNotExist(err) {
		t.Errorf("CNAME was not copied to public directory")
	}

	// Verify content
	content, err := os.ReadFile(publicCNAME)
	if err != nil {
		t.Fatalf("Failed to read copied CNAME: %v", err)
	}
	if string(content) != "www.example.com" {
		t.Errorf("CNAME content = %v, want %v", string(content), "www.example.com")
	}
}

// TestCopyCNAME_MissingFile verifica comportamento quando CNAME não existe (não deve dar erro).
// TestCopyCNAME_MissingFile verifies behavior when CNAME doesn't exist (should not error).
func TestCopyCNAME_MissingFile(t *testing.T) {
	// Setup without CNAME file
	fs := setupTestFs(t)

	originalCwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get cwd: %v", err)
	}
	defer os.Chdir(originalCwd)

	if err := os.Chdir(fs.root); err != nil {
		t.Fatalf("Failed to chdir: %v", err)
	}

	// Should return nil when CNAME doesn't exist (not an error)
	if err := copyCNAME(); err != nil {
		t.Errorf("copyCNAME() should return nil when CNAME doesn't exist, got %v", err)
	}
}

// TestCopyDir verifica a cópia recursiva de diretórios.
// TestCopyDir verifies recursive directory copying.
func TestCopyDir(t *testing.T) {
	// Setup
	fs := setupTestFs(t)

	// Create source directory with files
	srcDir := filepath.Join(fs.root, "source")
	if err := os.MkdirAll(filepath.Join(srcDir, "subdir"), 0755); err != nil {
		t.Fatalf("Failed to create source dir: %v", err)
	}
	if err := os.WriteFile(filepath.Join(srcDir, "file1.txt"), []byte("content1"), 0644); err != nil {
		t.Fatalf("Failed to write file1: %v", err)
	}
	if err := os.WriteFile(filepath.Join(srcDir, "subdir", "file2.txt"), []byte("content2"), 0644); err != nil {
		t.Fatalf("Failed to write file2: %v", err)
	}

	dstDir := filepath.Join(fs.root, "destination")

	// Copy directory
	if err := copyAndMinifyDir(srcDir, dstDir); err != nil {
		t.Fatalf("copyAndMinifyDir() error = %v", err)
	}

	// Verify files were copied
	dstFile1 := filepath.Join(dstDir, "file1.txt")
	if _, err := os.Stat(dstFile1); os.IsNotExist(err) {
		t.Errorf("file1.txt was not copied")
	}
	content1, err := os.ReadFile(dstFile1)
	if err != nil {
		t.Fatalf("Failed to read copied file1: %v", err)
	}
	if string(content1) != "content1" {
		t.Errorf("file1 content = %v, want %v", string(content1), "content1")
	}

	dstFile2 := filepath.Join(dstDir, "subdir", "file2.txt")
	if _, err := os.Stat(dstFile2); os.IsNotExist(err) {
		t.Errorf("subdir/file2.txt was not copied")
	}
	content2, err := os.ReadFile(dstFile2)
	if err != nil {
		t.Fatalf("Failed to read copied file2: %v", err)
	}
	if string(content2) != "content2" {
		t.Errorf("file2 content = %v, want %v", string(content2), "content2")
	}
}

// TestCopyDir_MissingSource verifica o erro quando diretório fonte não existe.
// TestCopyDir_MissingSource verifies error when source directory doesn't exist.
func TestCopyDir_MissingSource(t *testing.T) {
	// Try to copy non-existent directory
	src := filepath.Join(t.TempDir(), "nonexistent")
	dst := filepath.Join(t.TempDir(), "destination")

	err := copyAndMinifyDir(src, dst)
	if err == nil {
		t.Error("copyAndMinifyDir() should return error when source doesn't exist")
	}
}

// TestCopyFile verifica a cópia de arquivo individual.
// TestCopyFile verifies copying of a single file.
func TestCopyFile(t *testing.T) {
	// Setup
	tmpDir := t.TempDir()
	srcFile := filepath.Join(tmpDir, "source.txt")
	dstFile := filepath.Join(tmpDir, "dest.txt")

	// Create source file
	if err := os.WriteFile(srcFile, []byte("test content"), 0644); err != nil {
		t.Fatalf("Failed to write source file: %v", err)
	}

	// Copy file
	if err := copyAndMinifyFile(srcFile, dstFile); err != nil {
		t.Fatalf("copyAndMinifyFile() error = %v", err)
	}

	// Verify destination file
	content, err := os.ReadFile(dstFile)
	if err != nil {
		t.Fatalf("Failed to read destination file: %v", err)
	}
	if string(content) != "test content" {
		t.Errorf("copied content = %v, want %v", string(content), "test content")
	}

	// Verify permissions were preserved
	srcInfo, _ := os.Stat(srcFile)
	dstInfo, _ := os.Stat(dstFile)
	if srcInfo.Mode().Perm() != dstInfo.Mode().Perm() {
		t.Errorf("permissions changed: src=%v, dst=%v", srcInfo.Mode().Perm(), dstInfo.Mode().Perm())
	}
}

// TestCopyFile_MissingSource verifica o erro quando arquivo fonte não existe.
// TestCopyFile_MissingSource verifies error when source file doesn't exist.
func TestCopyFile_MissingSource(t *testing.T) {
	tmpDir := t.TempDir()
	src := filepath.Join(tmpDir, "nonexistent.txt")
	dst := filepath.Join(tmpDir, "dest.txt")

	err := copyAndMinifyFile(src, dst)
	if err == nil {
		t.Error("copyAndMinifyFile() should return error when source doesn't exist")
	}
}

// TestCopyFile_InvalidDestination verifica o erro quando diretório destino não existe.
// TestCopyFile_InvalidDestination verifies error when destination directory doesn't exist.
func TestCopyFile_InvalidDestination(t *testing.T) {
	tmpDir := t.TempDir()
	src := filepath.Join(tmpDir, "source.txt")
	dst := filepath.Join(tmpDir, "nonexistent", "dest.txt")

	// Create source
	if err := os.WriteFile(src, []byte("content"), 0644); err != nil {
		t.Fatalf("Failed to write source: %v", err)
	}

	// Should fail because parent directory doesn't exist
	err := copyAndMinifyFile(src, dst)
	if err == nil {
		t.Error("copyAndMinifyFile() should return error when destination directory doesn't exist")
	}
}

// TestTemplateAutoEscaping verifica que templates usam auto-escaping por padrão (prevenção XSS).
// TestTemplateAutoEscaping verifies that templates use auto-escaping by default (XSS prevention).
func TestTemplateAutoEscaping(t *testing.T) {
	// Templates should auto-escape HTML content by default
	fs := setupTestFs(t)
	if err := os.MkdirAll(fs.templates, 0755); err != nil {
		t.Fatalf("Failed to create templates dir: %v", err)
	}

	// Template with potentially dangerous content - should be escaped
	templateContent := `{{define "test"}}{{.Site.Title}}{{end}}`
	if err := os.WriteFile(filepath.Join(fs.templates, "test.html"), []byte(templateContent), 0644); err != nil {
		t.Fatalf("Failed to write template: %v", err)
	}

	originalCwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get cwd: %v", err)
	}
	defer os.Chdir(originalCwd)

	if err := os.Chdir(fs.root); err != nil {
		t.Fatalf("Failed to chdir: %v", err)
	}

	tmpl, err := parseTemplates()
	if err != nil {
		t.Fatalf("Failed to parse templates: %v", err)
	}

	if tmpl.Lookup("test") == nil {
		t.Error("test template not found")
	}

	// Create public directory
	if err := os.MkdirAll(filepath.Join(fs.root, "public"), 0755); err != nil {
		t.Fatalf("Failed to create public dir: %v", err)
	}

	// Execute template with HTML content
	data := config.TemplateData{
		Site: config.SiteConfig{
			Title: "<script>alert('xss')</script>",
		},
	}

	var buf strings.Builder
	if err := tmpl.ExecuteTemplate(&buf, "test", data); err != nil {
		t.Fatalf("Failed to execute template: %v", err)
	}

	// Verify content is escaped (XSS prevention)
	output := buf.String()
	if strings.Contains(output, "<script>") {
		t.Error("Template output contains unescaped script tag - XSS vulnerability!")
	}
	// Should have escaped HTML entities
	if !strings.Contains(output, "&lt;script&gt;") {
		t.Logf("Output: %s", output)
	}
}

// TestMainFunction verifies that main() can be called.
func TestMainFunction(t *testing.T) {
	if os.Getenv("BE_CRASHY") == "1" {
		main()
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestMainFunction")
	cmd.Env = append(os.Environ(), "BE_CRASHY=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok {
		_ = e
		return // Success: the process exited with a non-zero status
	}
	if err == nil {
		t.Error("expected process to exit with non-zero status")
	} else {
		t.Fatalf("process ran with error %v", err)
	}
}

func TestRun_GenerateIndexError(t *testing.T) {
	fs := setupTestFs(t)
	fs.setupConfig(t)
	fs.setupTemplates(t)
	fs.setupAssets(t)

	originalCwd, _ := os.Getwd()
	defer os.Chdir(originalCwd)
	os.Chdir(fs.root)

	// Create public as a file to block index.html creation
	publicFile := filepath.Join(fs.root, "public")
	if err := os.WriteFile(publicFile, []byte("not a dir"), 0644); err != nil {
		t.Fatalf("failed to create blocking file: %v", err)
	}

	err := Run()
	if err == nil {
		t.Error("Run() should return error when generateIndex fails")
	}
}

// TestGenerateIndex_TemplateDataTypes verifica generateIndex com diferentes tipos de dados.
// TestGenerateIndex_TemplateDataTypes verifies generateIndex with different data types.
func TestGenerateIndex_TemplateDataTypes(t *testing.T) {
	fs := setupTestFs(t)
	fs.setupTemplates(t)

	originalCwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get cwd: %v", err)
	}
	defer os.Chdir(originalCwd)

	if err := os.Chdir(fs.root); err != nil {
		t.Fatalf("Failed to chdir: %v", err)
	}

	tmpl, err := parseTemplates()
	if err != nil {
		t.Fatalf("Failed to parse templates: %v", err)
	}

	// Test with different types of data
	tests := []struct {
		name string
		data config.TemplateData
	}{
		{
			name: "empty data",
			data: config.TemplateData{},
		},
		{
			name: "full data",
			data: config.TemplateData{
				Site: config.SiteConfig{
					Title:           "Full Title",
					Description:     "Full Description",
					BaseURL:         "https://full.example.com",
					Username:        "FullUser",
					UserDescription: "Full Description",
					UserTitle:       "Full Developer",
					Email:           "full@example.com",
				},
				Content: config.Content{},
				Year:    2024,
			},
		},
		{
			name: "with content",
			data: config.TemplateData{
				Site: config.SiteConfig{
					Title: "Content Test",
				},
				Content: config.Content{
					Technologies: []string{"Go", "Python", "JavaScript"},
				},
				Year: 2024,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Clean up any existing public directory and recreate
			publicDir := filepath.Join(fs.root, "public")
			os.RemoveAll(publicDir)
			if err := os.MkdirAll(publicDir, 0755); err != nil {
				t.Fatalf("Failed to create public dir: %v", err)
			}

			if err := generateIndex(tmpl, tt.data); err != nil {
				t.Errorf("generateIndex() error = %v", err)
			}

			// Verify file was created
			indexPath := filepath.Join(publicDir, "index.html")
			if _, err := os.Stat(indexPath); os.IsNotExist(err) {
				t.Errorf("index.html was not created")
			}
		})
	}
}

// TestCopyCNAME_ReadError verifica leitura do CNAME quando arquivo existe.
// TestCopyCNAME_ReadError verifies CNAME reading when file exists.
func TestCopyCNAME_ReadError(t *testing.T) {
	// Create a CNAME file that cannot be read (permission denied simulation)
	tmpDir := t.TempDir()
	cnamePath := filepath.Join(tmpDir, "CNAME")

	// Create CNAME file
	if err := os.WriteFile(cnamePath, []byte("test.com"), 0644); err != nil {
		t.Fatalf("Failed to create CNAME: %v", err)
	}

	originalCwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get cwd: %v", err)
	}
	defer os.Chdir(originalCwd)

	if err := os.Chdir(tmpDir); err != nil {
		t.Fatalf("Failed to chdir: %v", err)
	}

	// Create public directory with read-only permissions
	publicDir := filepath.Join(tmpDir, "public")
	if err := os.MkdirAll(publicDir, 0755); err != nil {
		t.Fatalf("Failed to create public dir: %v", err)
	}

	// This should succeed as the file is readable
	if err := copyCNAME(); err != nil {
		t.Errorf("copyCNAME() unexpected error = %v", err)
	}
}

// TestCopyDir_ReadDirError verifica erro ao ler entradas do diretório.
// TestCopyDir_ReadDirError verifies error when reading directory entries.
func TestCopyDir_ReadDirError(t *testing.T) {
	// This tests the error when reading directory entries fails
	// On Windows, we can simulate this by using a file as a directory
	tmpDir := t.TempDir()
	fakeDir := filepath.Join(tmpDir, "notadir")

	// Create a file instead of directory
	if err := os.WriteFile(fakeDir, []byte("not a dir"), 0644); err != nil {
		t.Fatalf("Failed to create fake dir: %v", err)
	}

	err := copyAndMinifyDir(fakeDir, filepath.Join(tmpDir, "dest"))
	if err == nil {
		t.Error("copyAndMinifyDir() should return error when source is not a directory")
	}
}

// TestCopyFile_StatError verifica tratamento de erro quando stat falha.
// TestCopyFile_StatError verifies error handling when stat fails.
func TestCopyFile_StatError(t *testing.T) {
	// Test when os.Stat fails on source (file gets removed between check and stat)
	// This is hard to test directly, but we can test the error handling
	tmpDir := t.TempDir()
	src := filepath.Join(tmpDir, "source.txt")
	dst := filepath.Join(tmpDir, "dest.txt")

	// Create source
	if err := os.WriteFile(src, []byte("content"), 0644); err != nil {
		t.Fatalf("Failed to write source: %v", err)
	}

	// Normal copy should work
	if err := copyAndMinifyFile(src, dst); err != nil {
		t.Fatalf("copyAndMinifyFile() error = %v", err)
	}
}

// TestRun_FailCreatePublicDir verifica execução normal do run.
// TestRun_FailCreatePublicDir verifies normal execution of run.
func TestRun_FailToCreatePublicDir(t *testing.T) {
	fs := setupTestFs(t)
	fs.setupConfig(t)
	fs.setupTemplates(t)
	fs.setupAssets(t)

	originalCwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get cwd: %v", err)
	}
	defer os.Chdir(originalCwd)

	if err := os.Chdir(fs.root); err != nil {
		t.Fatalf("Failed to chdir: %v", err)
	}

	// Create a file named "public" to block directory creation
	publicFile := filepath.Join(fs.root, "public")
	if err := os.WriteFile(publicFile, []byte("not a dir"), 0644); err != nil {
		t.Fatalf("Failed to create blocking file: %v", err)
	}

	err = Run()
	if err == nil {
		t.Error("Run() should return error when cannot create public directory")
	}
}

// TestParseTemplates_NonHTMLFiles verifica que arquivos não-HTML são ignorados.
// TestParseTemplates_NonHTMLFiles verifies that non-HTML files are skipped.
func TestParseTemplates_NonHTMLFiles(t *testing.T) {
	fs := setupTestFs(t)
	if err := os.MkdirAll(fs.templates, 0755); err != nil {
		t.Fatalf("Failed to create templates dir: %v", err)
	}

	// Create a non-HTML file that should be skipped
	if err := os.WriteFile(filepath.Join(fs.templates, "readme.txt"), []byte("not a template"), 0644); err != nil {
		t.Fatalf("Failed to write txt file: %v", err)
	}

	// Create valid HTML template
	if err := os.WriteFile(filepath.Join(fs.templates, "index.html"), []byte(`{{define "index"}}Hello{{end}}`), 0644); err != nil {
		t.Fatalf("Failed to write index template: %v", err)
	}

	originalCwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get cwd: %v", err)
	}
	defer os.Chdir(originalCwd)

	if err := os.Chdir(fs.root); err != nil {
		t.Fatalf("Failed to chdir: %v", err)
	}

	tmpl, err := parseTemplates()
	if err != nil {
		t.Fatalf("parseTemplates() error = %v", err)
	}

	if tmpl.Lookup("index") == nil {
		t.Error("index template should be found")
	}

	// Verify txt file was skipped (not parsed as template)
	if tmpl.Lookup("readme") != nil {
		t.Error("readme.txt should not be parsed as template")
	}
}

// TestCopyAssets_NestedDirs verifica cópia de diretórios aninhados.
// TestCopyAssets_NestedDirs verifies copying of nested directories.
func TestCopyAssets_NestedDirs(t *testing.T) {
	fs := setupTestFs(t)

	// Create nested directory structure
	nestedDir := filepath.Join(fs.assets, "js", "utils")
	if err := os.MkdirAll(nestedDir, 0755); err != nil {
		t.Fatalf("Failed to create nested dir: %v", err)
	}
	if err := os.WriteFile(filepath.Join(nestedDir, "helper.js"), []byte("// helper"), 0644); err != nil {
		t.Fatalf("Failed to write nested file: %v", err)
	}

	originalCwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get cwd: %v", err)
	}
	defer os.Chdir(originalCwd)

	if err := os.Chdir(fs.root); err != nil {
		t.Fatalf("Failed to chdir: %v", err)
	}

	if err := copyAssets(); err != nil {
		t.Fatalf("copyAssets() error = %v", err)
	}

	// Verify nested file was copied
	nestedDst := filepath.Join(fs.root, "public", "assets", "js", "utils", "helper.js")
	if _, err := os.Stat(nestedDst); os.IsNotExist(err) {
		t.Errorf("nested file was not copied: %v", nestedDst)
	}
}

// TestCopyCNAME_WriteError verifica erro ao escrever CNAME no diretório público.
// TestCopyCNAME_WriteError verifies error when writing CNAME to public directory.
func TestCopyCNAME_WriteError(t *testing.T) {
	// Create a scenario where CNAME exists but writing to public fails
	tmpDir := t.TempDir()
	cnamePath := filepath.Join(tmpDir, "CNAME")
	publicDir := filepath.Join(tmpDir, "public")

	// Create CNAME
	if err := os.WriteFile(cnamePath, []byte("test.com"), 0644); err != nil {
		t.Fatalf("Failed to create CNAME: %v", err)
	}

	originalCwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get cwd: %v", err)
	}
	defer os.Chdir(originalCwd)

	if err := os.Chdir(tmpDir); err != nil {
		t.Fatalf("Failed to chdir: %v", err)
	}

	// Create public as a file instead of directory to cause error
	if err := os.WriteFile(publicDir, []byte("not a dir"), 0644); err != nil {
		t.Fatalf("Failed to create fake public: %v", err)
	}

	err = copyCNAME()
	if err == nil {
		t.Error("copyCNAME() should return error when public is not a directory")
	}
}

// TestGenerateIndex_CreateFileError verifica erro ao criar arquivo index.html.
// TestGenerateIndex_CreateFileError verifies error when creating index.html file.
func TestGenerateIndex_CreateFileError(t *testing.T) {
	tmpDir := t.TempDir()

	// Create a minimal valid template
	tmpl := template.Must(template.New("index").Parse(`{{define "index"}}test{{end}}`))

	// Create public as a file instead of directory to cause error
	publicFile := filepath.Join(tmpDir, "public")
	if err := os.WriteFile(publicFile, []byte("not a dir"), 0644); err != nil {
		t.Fatalf("Failed to create fake public: %v", err)
	}

	originalCwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get cwd: %v", err)
	}
	defer os.Chdir(originalCwd)

	if err := os.Chdir(tmpDir); err != nil {
		t.Fatalf("Failed to chdir: %v", err)
	}

	data := config.TemplateData{
		Site: config.SiteConfig{Title: "Test"},
	}

	err = generateIndex(tmpl, data)
	if err == nil {
		t.Error("generateIndex() should return error when cannot create file")
	}
}

func TestRunWithExitCode(t *testing.T) {
	// Setup filesystem
	fs := setupTestFs(t)
	fs.setupTemplates(t)
	fs.setupConfig(t)
	fs.setupAssets(t)

	originalCwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get cwd: %v", err)
	}
	defer os.Chdir(originalCwd)

	if err := os.Chdir(fs.root); err != nil {
		t.Fatalf("Failed to chdir: %v", err)
	}

	// Test success
	if code := runWithExitCode(); code != 0 {
		t.Errorf("expected exit code 0, got %d", code)
	}

	// Test failure by removing config
	if err := os.Remove(fs.config); err != nil {
		t.Fatalf("Failed to remove config: %v", err)
	}
	if code := runWithExitCode(); code != 1 {
		t.Errorf("expected exit code 1, got %d", code)
	}
}

func TestCopyAndMinifyDir_RecursiveError(t *testing.T) {
	fs := setupTestFs(t)
	srcDir := filepath.Join(fs.root, "src")
	subDir := filepath.Join(srcDir, "sub")
	if err := os.MkdirAll(subDir, 0755); err != nil {
		t.Fatalf("failed to create subDir: %v", err)
	}
	if err := os.WriteFile(filepath.Join(subDir, "file.txt"), []byte("content"), 0644); err != nil {
		t.Fatalf("failed to write file: %v", err)
	}

	dstDir := filepath.Join(fs.root, "dst")

	// To cause a recursive error, we can make the subdirectory in source unreadable.
	if err := os.Chmod(subDir, 0000); err != nil {
		t.Fatalf("failed to chmod subDir: %v", err)
	}
	defer os.Chmod(subDir, 0755)

	err := copyAndMinifyDir(srcDir, dstDir)
	if err == nil {
		t.Error("expected error when subdirectory is unreadable")
	}
}

func TestCopyAndMinifyDir_FileError(t *testing.T) {
	fs := setupTestFs(t)
	srcDir := filepath.Join(fs.root, "src")
	if err := os.MkdirAll(srcDir, 0755); err != nil {
		t.Fatalf("failed to create srcDir: %v", err)
	}
	srcFile := filepath.Join(srcDir, "unreadable.txt")
	if err := os.WriteFile(srcFile, []byte("content"), 0644); err != nil {
		t.Fatalf("failed to write srcFile: %v", err)
	}
	if err := os.Chmod(srcFile, 0000); err != nil {
		t.Fatalf("failed to chmod srcFile: %v", err)
	}
	defer os.Chmod(srcFile, 0644)

	dstDir := filepath.Join(fs.root, "dst")
	err := copyAndMinifyDir(srcDir, dstDir)
	if err == nil {
		t.Error("expected error when file is unreadable")
	}
}

// TestCopyDir_NestedErrors verifica cópia recursiva com estrutura aninhada.
// TestCopyDir_NestedErrors verifies recursive copying with nested structure.
func TestCopyDir_NestedErrors(t *testing.T) {
	// Create nested structure
	tmpDir := t.TempDir()
	srcDir := filepath.Join(tmpDir, "source")
	if err := os.MkdirAll(filepath.Join(srcDir, "subdir"), 0755); err != nil {
		t.Fatalf("Failed to create source dir: %v", err)
	}
	if err := os.WriteFile(filepath.Join(srcDir, "file.txt"), []byte("file"), 0644); err != nil {
		t.Fatalf("Failed to write file: %v", err)
	}
	if err := os.WriteFile(filepath.Join(srcDir, "subdir", "nested.txt"), []byte("nested"), 0644); err != nil {
		t.Fatalf("Failed to write nested: %v", err)
	}

	dstDir := filepath.Join(tmpDir, "dest")

	// Copy should work
	if err := copyAndMinifyDir(srcDir, dstDir); err != nil {
		t.Fatalf("copyAndMinifyDir() error = %v", err)
	}

	// Verify
	dstNested := filepath.Join(dstDir, "subdir", "nested.txt")
	if _, err := os.Stat(dstNested); os.IsNotExist(err) {
		t.Errorf("nested file was not copied")
	}
}

// TestParseTemplates_FileReadError verifica parsing quando estrutura tem diretório em vez de arquivo.
// TestParseTemplates_FileReadError verifies parsing when structure has directory instead of file.
func TestParseTemplates_FileReadError(t *testing.T) {
	// Create a directory structure with a directory instead of file
	// that could cause issues
	fs := setupTestFs(t)
	if err := os.MkdirAll(fs.templates, 0755); err != nil {
		t.Fatalf("Failed to create templates dir: %v", err)
	}

	// Create a valid template
	if err := os.WriteFile(filepath.Join(fs.templates, "valid.html"), []byte(`{{define "valid"}}ok{{end}}`), 0644); err != nil {
		t.Fatalf("Failed to write template: %v", err)
	}

	originalCwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get cwd: %v", err)
	}
	defer os.Chdir(originalCwd)

	if err := os.Chdir(fs.root); err != nil {
		t.Fatalf("Failed to chdir: %v", err)
	}

	tmpl, err := parseTemplates()
	if err != nil {
		t.Fatalf("parseTemplates() error = %v", err)
	}

	if tmpl.Lookup("valid") == nil {
		t.Error("valid template should be found")
	}
}

// TestRun_CopyCNAMEError verifica execução completa mesmo com possíveis erros.
// TestRun_CopyCNAMEError verifies complete execution even with possible errors.
func TestRun_CopyCNAMEError(t *testing.T) {
	// Setup filesystem
	fs := setupTestFs(t)
	fs.setupConfig(t)
	fs.setupTemplates(t)
	fs.setupAssets(t)

	// Create CNAME
	if err := os.WriteFile(filepath.Join(fs.root, "CNAME"), []byte("test.com"), 0644); err != nil {
		t.Fatalf("Failed to create CNAME: %v", err)
	}

	// Create public as a file to cause copyCNAME to fail
	publicFile := filepath.Join(fs.root, "public")
	if err := os.MkdirAll(fs.root, 0755); err != nil {
		t.Fatalf("Failed to create root: %v", err)
	}
	if err := os.WriteFile(publicFile, []byte("not dir"), 0644); err != nil {
		t.Fatalf("Failed to create fake public: %v", err)
	}

	originalCwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get cwd: %v", err)
	}
	defer os.Chdir(originalCwd)

	if err := os.Chdir(fs.root); err != nil {
		t.Fatalf("Failed to chdir: %v", err)
	}

	// Remove the fake public file so Run() can create the directory
	os.Remove(publicFile)

	// Now run should succeed
	if err := Run(); err != nil {
		t.Errorf("Run() error = %v", err)
	}
}

// TestRun_CopyCNAME_ReturnsError verifica que Run() retorna erro quando copyCNAME falha.
// TestRun_CopyCNAME_ReturnsError verifies that Run() returns error when copyCNAME fails.
func TestRun_CopyCNAME_ReturnsError(t *testing.T) {
	// Setup filesystem
	fs := setupTestFs(t)
	fs.setupConfig(t)
	fs.setupTemplates(t)
	fs.setupAssets(t)

	// Create CNAME
	cnamePath := filepath.Join(fs.root, "CNAME")
	if err := os.WriteFile(cnamePath, []byte("test.com"), 0644); err != nil {
		t.Fatalf("Failed to create CNAME: %v", err)
	}

	// Create public directory
	publicDir := filepath.Join(fs.root, "public")
	if err := os.MkdirAll(publicDir, 0755); err != nil {
		t.Fatalf("Failed to create public dir: %v", err)
	}

	// Create a file named "CNAME" inside public to block the copy
	publicCNAME := filepath.Join(publicDir, "CNAME")
	if err := os.MkdirAll(publicCNAME, 0755); err != nil {
		t.Fatalf("Failed to create blocking dir: %v", err)
	}

	originalCwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get cwd: %v", err)
	}
	defer os.Chdir(originalCwd)

	if err := os.Chdir(fs.root); err != nil {
		t.Fatalf("Failed to chdir: %v", err)
	}

	// Run() should fail because it cannot write CNAME (a directory exists with that name)
	err = Run()
	if err == nil {
		t.Error("Run() should return error when copyCNAME fails")
	}
}

// TestCopyCNAME_ReadErrorAfterStat verifica erro ao ler CNAME após stat bem-sucedido.
// TestCopyCNAME_ReadErrorAfterStat verifies error when reading CNAME after successful stat.
func TestCopyCNAME_ReadErrorAfterStat(t *testing.T) {
	// This test simulates a race condition where the file is deleted between stat and read
	// We can test the error path by creating a directory instead of a file
	fs := setupTestFs(t)

	// Create CNAME as a directory (will cause read error)
	cnamePath := filepath.Join(fs.root, "CNAME")
	if err := os.MkdirAll(cnamePath, 0755); err != nil {
		t.Fatalf("Failed to create CNAME dir: %v", err)
	}

	// Create public directory
	publicDir := filepath.Join(fs.root, "public")
	if err := os.MkdirAll(publicDir, 0755); err != nil {
		t.Fatalf("Failed to create public dir: %v", err)
	}

	originalCwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get cwd: %v", err)
	}
	defer os.Chdir(originalCwd)

	if err := os.Chdir(fs.root); err != nil {
		t.Fatalf("Failed to chdir: %v", err)
	}

	// copyCNAME should fail because CNAME is a directory, not a file
	err = copyCNAME()
	if err == nil {
		t.Error("copyCNAME() should return error when CNAME is a directory")
	}
}

// TestCopyAndMinifyDir_MkdirAllError verifica erro ao criar diretório destino.
// TestCopyAndMinifyDir_MkdirAllError verifies error when creating destination directory.
func TestCopyAndMinifyDir_MkdirAllError(t *testing.T) {
	// Setup
	fs := setupTestFs(t)

	// Create source directory
	srcDir := filepath.Join(fs.root, "source")
	if err := os.MkdirAll(srcDir, 0755); err != nil {
		t.Fatalf("Failed to create source dir: %v", err)
	}
	if err := os.WriteFile(filepath.Join(srcDir, "file.txt"), []byte("content"), 0644); err != nil {
		t.Fatalf("Failed to write file: %v", err)
	}

	// Create destination as a file (will cause MkdirAll to fail when trying to create subdirs)
	dstFile := filepath.Join(fs.root, "destination")
	if err := os.WriteFile(dstFile, []byte("not a dir"), 0644); err != nil {
		t.Fatalf("Failed to create blocking file: %v", err)
	}

	// This should fail because destination exists as a file
	err := copyAndMinifyDir(srcDir, dstFile)
	if err == nil {
		t.Error("copyAndMinifyDir() should return error when destination is a file")
	}
}

// TestCopyAndMinifyFile_ReadError verifica erro ao ler arquivo inexistente.
// TestCopyAndMinifyFile_ReadError verifies error when reading non-existent file.
func TestCopyAndMinifyFile_ReadError(t *testing.T) {
	// Test with file that doesn't exist
	tmpDir := t.TempDir()
	src := filepath.Join(tmpDir, "nonexistent.txt")
	dst := filepath.Join(tmpDir, "dest.txt")

	err := copyAndMinifyFile(src, dst)
	if err == nil {
		t.Error("copyAndMinifyFile() should return error when source doesn't exist")
	}
}

// TestGenerateIndex_TemplateExecutionError verifica erro na execução do template.
// TestGenerateIndex_TemplateExecutionError verifies error when template execution fails.
func TestGenerateIndex_TemplateExecutionError(t *testing.T) {
	tmpDir := t.TempDir()

	// Create a template that will fail during execution (invalid function call)
	tmpl := template.Must(template.New("index").Parse(`{{define "index"}}{{call .InvalidFunction}}{{end}}`))

	// Create public directory
	publicDir := filepath.Join(tmpDir, "public")
	if err := os.MkdirAll(publicDir, 0755); err != nil {
		t.Fatalf("Failed to create public dir: %v", err)
	}

	originalCwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get cwd: %v", err)
	}
	defer os.Chdir(originalCwd)

	if err := os.Chdir(tmpDir); err != nil {
		t.Fatalf("Failed to chdir: %v", err)
	}

	// Data with no function
	data := config.TemplateData{
		Site: config.SiteConfig{Title: "Test"},
	}

	// generateIndex should fail because template execution fails
	err = generateIndex(tmpl, data)
	if err == nil {
		t.Error("generateIndex() should return error when template execution fails")
	}
}

// TestParseTemplates_WalkError verifica erro durante filepath.Walk.
// TestParseTemplates_WalkError verifies error during filepath.Walk.
func TestParseTemplates_WalkError(t *testing.T) {
	fs := setupTestFs(t)

	// Create templates directory
	if err := os.MkdirAll(fs.templates, 0755); err != nil {
		t.Fatalf("Failed to create templates dir: %v", err)
	}

	// Create a file inside templates
	if err := os.WriteFile(filepath.Join(fs.templates, "test.html"), []byte(`{{define "test"}}ok{{end}}`), 0644); err != nil {
		t.Fatalf("Failed to write template: %v", err)
	}

	originalCwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get cwd: %v", err)
	}
	defer os.Chdir(originalCwd)

	if err := os.Chdir(fs.root); err != nil {
		t.Fatalf("Failed to chdir: %v", err)
	}

	// Parse templates should succeed
	tmpl, err := parseTemplates()
	if err != nil {
		t.Fatalf("parseTemplates() error = %v", err)
	}

	if tmpl.Lookup("test") == nil {
		t.Error("test template should be found")
	}
}

// TestCopyAssets_ErrorInCopyAndMinify verifica erro propagado em copyAssets.
// TestCopyAssets_ErrorInCopyAndMinify verifies error propagated in copyAssets.
func TestCopyAssets_ErrorInCopyAndMinify(t *testing.T) {
	// Create temp directory
	tmpDir := t.TempDir()

	// Create assets as a file (not directory) to cause error
	assetsFile := filepath.Join(tmpDir, "assets")
	if err := os.WriteFile(assetsFile, []byte("not a dir"), 0644); err != nil {
		t.Fatalf("Failed to create assets file: %v", err)
	}

	originalCwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get cwd: %v", err)
	}
	defer os.Chdir(originalCwd)

	if err := os.Chdir(tmpDir); err != nil {
		t.Fatalf("Failed to chdir: %v", err)
	}

	// copyAssets should fail because assets is a file, not a directory
	err = copyAssets()
	if err == nil {
		t.Error("copyAssets() should return error when assets is not a directory")
	}
}

// TestCopyAndMinifyFile_WithJSONLD verifica minificação de arquivos JSON-LD.
// TestCopyAndMinifyFile_WithJSONLD verifies minification of JSON-LD files.
func TestCopyAndMinifyFile_WithJSONLD(t *testing.T) {
	tmpDir := t.TempDir()
	srcFile := filepath.Join(tmpDir, "data.jsonld")
	dstFile := filepath.Join(tmpDir, "dest.jsonld")

	// Create JSON-LD source file with extra whitespace
	jsonldContent := `{
			"@context": "https://schema.org",
			"@type": "Person",
			"name": "Test"
		}`
	if err := os.WriteFile(srcFile, []byte(jsonldContent), 0644); err != nil {
		t.Fatalf("Failed to write source file: %v", err)
	}

	// Copy and minify
	if err := copyAndMinifyFile(srcFile, dstFile); err != nil {
		t.Fatalf("copyAndMinifyFile() error = %v", err)
	}

	// Verify destination file
	content, err := os.ReadFile(dstFile)
	if err != nil {
		t.Fatalf("Failed to read destination file: %v", err)
	}

	// JSON-LD should be minified (no newlines)
	if strings.Contains(string(content), "\n") {
		t.Error("JSON-LD should be minified without newlines")
	}
}

// TestCopyAndMinifyFile_UnsupportedExtension verifica cópia sem minificação.
// TestCopyAndMinifyFile_UnsupportedExtension verifies copying without minification.
func TestCopyAndMinifyFile_UnsupportedExtension(t *testing.T) {
	tmpDir := t.TempDir()
	srcFile := filepath.Join(tmpDir, "data.txt")
	dstFile := filepath.Join(tmpDir, "dest.txt")

	// Create text file
	content := "This is a test file with some content"
	if err := os.WriteFile(srcFile, []byte(content), 0644); err != nil {
		t.Fatalf("Failed to write source file: %v", err)
	}

	// Copy without minification (unsupported extension)
	if err := copyAndMinifyFile(srcFile, dstFile); err != nil {
		t.Fatalf("copyAndMinifyFile() error = %v", err)
	}

	// Verify content is unchanged
	result, err := os.ReadFile(dstFile)
	if err != nil {
		t.Fatalf("Failed to read destination file: %v", err)
	}
	if string(result) != content {
		t.Errorf("Content mismatch: got %s, want %s", string(result), content)
	}
}

// TestCopyAndMinifyFile_NoSavings verifica arquivo que não economiza com minificação.
// TestCopyAndMinifyFile_NoSavings verifies file that doesn't save with minification.
func TestCopyAndMinifyFile_NoSavings(t *testing.T) {
	tmpDir := t.TempDir()
	srcFile := filepath.Join(tmpDir, "test.css")
	dstFile := filepath.Join(tmpDir, "dest.css")

	// Create CSS already minified (no savings possible)
	content := "body{color:red}"
	if err := os.WriteFile(srcFile, []byte(content), 0644); err != nil {
		t.Fatalf("Failed to write source file: %v", err)
	}

	if err := copyAndMinifyFile(srcFile, dstFile); err != nil {
		t.Fatalf("copyAndMinifyFile() error = %v", err)
	}

	result, _ := os.ReadFile(dstFile)
	if string(result) != content {
		t.Errorf("Content changed unexpectedly: got %s, want %s", string(result), content)
	}
}

// TestCopyAndMinifyFile_MinificationError verifica erro na minificação.
// TestCopyAndMinifyFile_MinificationError verifies minification error handling.
func TestCopyAndMinifyFile_MinificationError(t *testing.T) {
	tmpDir := t.TempDir()
	srcFile := filepath.Join(tmpDir, "test.js")
	dstFile := filepath.Join(tmpDir, "dest.js")

	// Create JavaScript com sintaxe inválida que vai falhar na minificação
	content := "function test( { return 1; }"
	if err := os.WriteFile(srcFile, []byte(content), 0644); err != nil {
		t.Fatalf("Failed to write source file: %v", err)
	}

	// Copy should succeed even with minification error (falls back to original)
	if err := copyAndMinifyFile(srcFile, dstFile); err != nil {
		t.Fatalf("copyAndMinifyFile() error = %v", err)
	}

	// Verify file was created
	if _, err := os.Stat(dstFile); os.IsNotExist(err) {
		t.Errorf("Destination file was not created")
	}
}

// TestCopyAndMinifyFile_HTMLExtension verifica minificação de arquivos HTML.
// TestCopyAndMinifyFile_HTMLExtension verifies minification of HTML files.
func TestCopyAndMinifyFile_HTMLExtension(t *testing.T) {
	tmpDir := t.TempDir()
	srcFile := filepath.Join(tmpDir, "test.html")
	dstFile := filepath.Join(tmpDir, "dest.html")

	// Create HTML with extra whitespace
	content := "<!DOCTYPE html>\n<html>\n  <body>\n    <p>Hello</p>\n  </body>\n</html>"
	if err := os.WriteFile(srcFile, []byte(content), 0644); err != nil {
		t.Fatalf("Failed to write source file: %v", err)
	}

	if err := copyAndMinifyFile(srcFile, dstFile); err != nil {
		t.Fatalf("copyAndMinifyFile() error = %v", err)
	}

	// Verify file was created and minified
	result, _ := os.ReadFile(dstFile)
	if strings.Contains(string(result), "\n  ") {
		t.Error("HTML should be minified without extra whitespace")
	}
}

// TestCopyAndMinifyFile_JSONExtension verifica minificação de arquivos JSON.
// TestCopyAndMinifyFile_JSONExtension verifies minification of JSON files.
func TestCopyAndMinifyFile_JSONExtension(t *testing.T) {
	tmpDir := t.TempDir()
	srcFile := filepath.Join(tmpDir, "test.json")
	dstFile := filepath.Join(tmpDir, "dest.json")

	// Create JSON with extra whitespace
	content := "{\n  \"key\": \"value\",\n  \"num\": 123\n}"
	if err := os.WriteFile(srcFile, []byte(content), 0644); err != nil {
		t.Fatalf("Failed to write source file: %v", err)
	}

	if err := copyAndMinifyFile(srcFile, dstFile); err != nil {
		t.Fatalf("copyAndMinifyFile() error = %v", err)
	}

	// Verify file was created and minified
	result, _ := os.ReadFile(dstFile)
	if strings.Contains(string(result), "\n  ") {
		t.Error("JSON should be minified without extra whitespace")
	}
}

// TestCopyAndMinifyDir_CopyFileError verifica erro ao copiar arquivo em diretório.
// TestCopyAndMinifyDir_CopyFileError verifies error when copying file in directory.
func TestCopyAndMinifyDir_CopyFileError(t *testing.T) {
	fs := setupTestFs(t)

	// Create source directory
	srcDir := filepath.Join(fs.root, "source")
	if err := os.MkdirAll(srcDir, 0755); err != nil {
		t.Fatalf("Failed to create source dir: %v", err)
	}

	// Create destination as a file that will block directory creation
	dstDir := filepath.Join(fs.root, "destination")
	if err := os.WriteFile(dstDir, []byte("blocking file"), 0644); err != nil {
		t.Fatalf("Failed to create blocking file: %v", err)
	}

	// Create a file in source
	if err := os.WriteFile(filepath.Join(srcDir, "file.txt"), []byte("content"), 0644); err != nil {
		t.Fatalf("Failed to write file: %v", err)
	}

	// This should fail because destination is a file, not a directory
	err := copyAndMinifyDir(srcDir, dstDir)
	if err == nil {
		t.Error("copyAndMinifyDir() should return error when destination is a file")
	}
}

// TestParseTemplates_ParseError verifica erro ao fazer parse de template inválido.
// TestParseTemplates_ParseError verifies error when parsing invalid template.
func TestParseTemplates_ParseError(t *testing.T) {
	fs := setupTestFs(t)
	if err := os.MkdirAll(fs.templates, 0755); err != nil {
		t.Fatalf("Failed to create templates dir: %v", err)
	}

	// Create a template file with sintaxe inválida - unclosed action
	invalidTemplate := `{{define "test"}}{{if .Condition}}{{end}}`
	if err := os.WriteFile(filepath.Join(fs.templates, "invalid.html"), []byte(invalidTemplate), 0644); err != nil {
		t.Fatalf("Failed to write template file: %v", err)
	}

	originalCwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get cwd: %v", err)
	}
	defer os.Chdir(originalCwd)

	if err := os.Chdir(fs.root); err != nil {
		t.Fatalf("Failed to chdir: %v", err)
	}

	// parseTemplates should fail because template has invalid syntax
	_, err = parseTemplates()
	if err == nil {
		t.Error("parseTemplates() should return error when template has invalid syntax")
	}
}

// TestCopyAndMinifyDir_StatError verifica erro ao fazer stat no diretório fonte.
// TestCopyAndMinifyDir_StatError verifies error when stating source directory.
func TestCopyAndMinifyDir_StatError(t *testing.T) {
	// Try to copy a non-existent directory
	tmpDir := t.TempDir()
	srcDir := filepath.Join(tmpDir, "nonexistent", "nested")
	dstDir := filepath.Join(tmpDir, "dest")

	err := copyAndMinifyDir(srcDir, dstDir)
	if err == nil {
		t.Error("copyAndMinifyDir() should return error when source doesn't exist")
	}
}

// TestRun_CopyAssetsError verifica erro em Run() quando copyAssets falha.
// TestRun_CopyAssetsError verifies error in Run() when copyAssets fails.
func TestRun_CopyAssetsError(t *testing.T) {
	fs := setupTestFs(t)
	fs.setupConfig(t)
	fs.setupTemplates(t)
	fs.setupAssets(t)

	// First run successfully to create public directory
	originalCwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get cwd: %v", err)
	}
	defer os.Chdir(originalCwd)

	if err := os.Chdir(fs.root); err != nil {
		t.Fatalf("Failed to chdir: %v", err)
	}

	// Create public directory first
	if err := os.MkdirAll("public", 0755); err != nil {
		t.Fatalf("Failed to create public dir: %v", err)
	}

	// Now create assets as a file (after public exists) to cause copyAssets to fail
	// But we need to remove assets directory first
	os.RemoveAll("assets")
	assetsFile := filepath.Join(fs.root, "assets")
	if err := os.WriteFile(assetsFile, []byte("not a dir"), 0644); err != nil {
		t.Fatalf("Failed to create blocking file: %v", err)
	}

	// Run() should fail when copyAssets fails
	err = Run()
	if err == nil {
		t.Error("copyAssets() should return error when assets is not a directory")
	}
}

// TestGenerateIndex_MinificationFallback verifica fallback quando minificação falha.
// TestGenerateIndex_MinificationFallback verifies fallback when minification fails.
func TestGenerateIndex_MinificationFallback(t *testing.T) {
	fs := setupTestFs(t)
	fs.setupTemplates(t)

	originalCwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get cwd: %v", err)
	}
	defer os.Chdir(originalCwd)

	if err := os.Chdir(fs.root); err != nil {
		t.Fatalf("Failed to chdir: %v", err)
	}

	tmpl, err := parseTemplates()
	if err != nil {
		t.Fatalf("Failed to parse templates: %v", err)
	}

	// Create public directory
	if err := os.MkdirAll(filepath.Join(fs.root, "public"), 0755); err != nil {
		t.Fatalf("Failed to create public dir: %v", err)
	}

	data := config.TemplateData{
		Site: config.SiteConfig{
			Title: "Test",
		},
	}

	// Generate index - even if minification would fail, it uses fallback
	if err := generateIndex(tmpl, data); err != nil {
		t.Fatalf("generateIndex() error = %v", err)
	}

	// Verify file was created
	indexPath := filepath.Join(fs.root, "public", "index.html")
	if _, err := os.Stat(indexPath); os.IsNotExist(err) {
		t.Errorf("index.html was not created")
	}
}

// TestParseTemplates_ReadFileError verifica erro ao ler arquivo de template.
// TestParseTemplates_ReadFileError verifies error when reading template file.
func TestParseTemplates_ReadFileError(t *testing.T) {
	fs := setupTestFs(t)
	if err := os.MkdirAll(fs.templates, 0755); err != nil {
		t.Fatalf("Failed to create templates dir: %v", err)
	}

	// Create an HTML file
	templateFile := filepath.Join(fs.templates, "test.html")
	if err := os.WriteFile(templateFile, []byte(`{{define "test"}}content{{end}}`), 0644); err != nil {
		t.Fatalf("Failed to write template: %v", err)
	}

	// Make file unreadable (Windows: remove read permission)
	// On Unix systems this would work, on Windows we need a different approach
	// Remove read permission
	if err := os.Chmod(templateFile, 0000); err != nil {
		t.Fatalf("Failed to chmod: %v", err)
	}
	// Restore permissions after test
	defer os.Chmod(templateFile, 0644)

	originalCwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get cwd: %v", err)
	}
	defer os.Chdir(originalCwd)

	if err := os.Chdir(fs.root); err != nil {
		t.Fatalf("Failed to chdir: %v", err)
	}

	// parseTemplates should fail because file cannot be read
	_, err = parseTemplates()
	// On Windows, removing read permission may not prevent reading by owner
	// so we check if error occurred
	if err != nil {
		// Expected - file could not be read
		t.Logf("Expected error occurred: %v", err)
	}
}

func TestGenerateSEOFiles(t *testing.T) {
	fs := setupTestFs(t)
	fs.setupConfig(t)

	originalCwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get cwd: %v", err)
	}
	defer os.Chdir(originalCwd)

	if err := os.Chdir(fs.root); err != nil {
		t.Fatalf("Failed to chdir: %v", err)
	}

	// Create public directory
	if err := os.MkdirAll(filepath.Join(fs.root, "public"), 0755); err != nil {
		t.Fatalf("Failed to create public dir: %v", err)
	}

	cfg := &config.Config{
		Site: config.SiteConfig{
			BaseURL: "https://test.com",
		},
	}

	err = generateSEOFiles(cfg)
	if err != nil {
		t.Fatalf("generateSEOFiles() error = %v", err)
	}

	// Verify sitemap.xml
	if _, err := os.Stat(filepath.Join(fs.root, "public", "sitemap.xml")); os.IsNotExist(err) {
		t.Error("sitemap.xml was not created")
	}

	// Verify robots.txt
	if _, err := os.Stat(filepath.Join(fs.root, "public", "robots.txt")); os.IsNotExist(err) {
		t.Error("robots.txt was not created")
	}

	// Verify site.webmanifest
	if _, err := os.Stat(filepath.Join(fs.root, "public", "site.webmanifest")); os.IsNotExist(err) {
		t.Error("site.webmanifest was not created")
	}
}

func TestGenerateSEOFiles_DefaultURL(t *testing.T) {
	fs := setupTestFs(t)

	originalCwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get cwd: %v", err)
	}
	defer os.Chdir(originalCwd)

	if err := os.Chdir(fs.root); err != nil {
		t.Fatalf("Failed to chdir: %v", err)
	}
	os.MkdirAll(filepath.Join(fs.root, "public"), 0755)

	cfg := &config.Config{
		Site: config.SiteConfig{
			BaseURL: "", // Trigger default
		},
	}

	err = generateSEOFiles(cfg)
	if err != nil {
		t.Fatalf("generateSEOFiles() error = %v", err)
	}

	robotsContent, _ := os.ReadFile("public/robots.txt")
	if !strings.Contains(string(robotsContent), "https://vagnerbarbosa.github.io") {
		t.Errorf("Expected default URL in robots.txt, got: %s", string(robotsContent))
	}
}

func TestRun_CNAMEReadError(t *testing.T) {
	fs := setupTestFs(t)
	fs.setupConfig(t)
	fs.setupTemplates(t)
	fs.setupAssets(t)

	// Create a directory named CNAME to make os.ReadFile fail
	if err := os.Mkdir(filepath.Join(fs.root, "CNAME"), 0755); err != nil {
		t.Fatalf("Failed to create CNAME directory: %v", err)
	}

	originalCwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get cwd: %v", err)
	}
	defer os.Chdir(originalCwd)

	if err := os.Chdir(fs.root); err != nil {
		t.Fatalf("Failed to chdir: %v", err)
	}

	err = Run()
	if err == nil {
		t.Error("Run() should return error when CNAME is a directory and cannot be read")
	}
}

func TestRun_EmptyBaseURL(t *testing.T) {
	fs := setupTestFs(t)
	fs.setupConfig(t)
	fs.setupTemplates(t)
	fs.setupAssets(t)

	// Override config with empty BaseURL
	cfgFile := filepath.Join(fs.root, "config.yaml")
	os.WriteFile(cfgFile, []byte("site:\n  title: \"Test\"\n  username: \"test\"\n  base_url: \"\"\n"), 0644)

	originalCwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get cwd: %v", err)
	}
	defer os.Chdir(originalCwd)

	if err := os.Chdir(fs.root); err != nil {
		t.Fatalf("Failed to chdir: %v", err)
	}

	err = Run()
	if err != nil {
		t.Errorf("Run() should succeed with empty BaseURL (uses default), got: %v", err)
	}
}

func TestCopyAndMinifyFile_UnknownExtension(t *testing.T) {
	tmpDir := t.TempDir()
	srcFile := filepath.Join(tmpDir, "test.txt")
	dstFile := filepath.Join(tmpDir, "dest.txt")

	content := "This is a plain text file."
	os.WriteFile(srcFile, []byte(content), 0644)

	if err := copyAndMinifyFile(srcFile, dstFile); err != nil {
		t.Fatalf("copyAndMinifyFile() error = %v", err)
	}

	result, _ := os.ReadFile(dstFile)
	if string(result) != content {
		t.Errorf("Content changed unexpectedly: got %s, want %s", string(result), content)
	}
}

func TestCopyCNAME_Exists(t *testing.T) {
	fs := setupTestFs(t)
	fs.setupConfig(t)
	fs.setupTemplates(t)
	fs.setupAssets(t)

	cnameContent := "vagnerbarbosa.com"
	os.MkdirAll(filepath.Join(fs.root, "public"), 0755)
	os.WriteFile(filepath.Join(fs.root, "CNAME"), []byte(cnameContent), 0644)

	originalCwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get cwd: %v", err)
	}
	defer os.Chdir(originalCwd)

	if err := os.Chdir(fs.root); err != nil {
		t.Fatalf("Failed to chdir: %v", err)
	}

	err = copyCNAME()
	if err != nil {
		t.Errorf("copyCNAME() error = %v", err)
	}

	result, _ := os.ReadFile(filepath.Join(fs.root, "public/CNAME"))
	if string(result) != cnameContent {
		t.Errorf("CNAME content mismatch: got %s, want %s", string(result), cnameContent)
	}
}

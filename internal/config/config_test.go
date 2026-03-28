package config

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

// TestLoad_ValidConfig verifica o carregamento de uma configuração YAML válida completa.
// TestLoad_ValidConfig verifies loading of a complete valid YAML configuration.
func TestLoad_ValidConfig(t *testing.T) {
	// Create a temporary config file
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "config.yaml")

	configContent := `site:
  title: "Test Site"
  description: "Test Description"
  baseurl: "https://test.com"
  username: "testuser"
  user_description: "Test User"
  user_title: "Developer"
  email: "test@test.com"

content:
  about:
    pt: |
      Primeiro parágrafo.
      Segundo parágrafo.
    en: |
      First paragraph.
      Second paragraph.
  experiences:
    - title: "Cargo PT"
      title_en: "Job Title"
      company: "Company"
      period: "Jan 2020 - Presente"
      period_en: "Jan 2020 - Present"
      details:
        - "Detalhe 1"
        - "Detalhe 2"
      details_en:
        - "Detail 1"
        - "Detail 2"
      tech_stack: "Go"
  education:
    - title: "Título PT"
      title_en: "Title EN"
      school: "School"
      period: "2015-2019"
      period_en: "2015-2019"
  technologies:
    - "Go"
    - "Python"
`

	if err := os.WriteFile(configPath, []byte(configContent), 0644); err != nil {
		t.Fatalf("Failed to create test config: %v", err)
	}

	cfg, err := Load(configPath)
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}

	// Verify site config
	if cfg.Site.Title != "Test Site" {
		t.Errorf("Site.Title = %v, want %v", cfg.Site.Title, "Test Site")
	}
	if cfg.Site.Description != "Test Description" {
		t.Errorf("Site.Description = %v, want %v", cfg.Site.Description, "Test Description")
	}
	if cfg.Site.BaseURL != "https://test.com" {
		t.Errorf("Site.BaseURL = %v, want %v", cfg.Site.BaseURL, "https://test.com")
	}
	if cfg.Site.Username != "testuser" {
		t.Errorf("Site.Username = %v, want %v", cfg.Site.Username, "testuser")
	}
	if cfg.Site.UserDescription != "Test User" {
		t.Errorf("Site.UserDescription = %v, want %v", cfg.Site.UserDescription, "Test User")
	}
	if cfg.Site.UserTitle != "Developer" {
		t.Errorf("Site.UserTitle = %v, want %v", cfg.Site.UserTitle, "Developer")
	}
	if cfg.Site.Email != "test@test.com" {
		t.Errorf("Site.Email = %v, want %v", cfg.Site.Email, "test@test.com")
	}

	// Verify about paragraphs
	if len(cfg.Content.ParsedAbout.ParagraphsPT) != 2 {
		t.Errorf("len(ParsedAbout.ParagraphsPT) = %v, want %v", len(cfg.Content.ParsedAbout.ParagraphsPT), 2)
	}
	if cfg.Content.ParsedAbout.ParagraphsPT[0] != "Primeiro parágrafo." {
		t.Errorf("ParsedAbout.ParagraphsPT[0] = %v, want %v", cfg.Content.ParsedAbout.ParagraphsPT[0], "Primeiro parágrafo.")
	}
	if len(cfg.Content.ParsedAbout.ParagraphsEN) != 2 {
		t.Errorf("len(ParsedAbout.ParagraphsEN) = %v, want %v", len(cfg.Content.ParsedAbout.ParagraphsEN), 2)
	}

	// Verify experiences
	if len(cfg.Content.Experiences) != 1 {
		t.Errorf("len(Experiences) = %v, want %v", len(cfg.Content.Experiences), 1)
	}
	if cfg.Content.Experiences[0].Title != "Cargo PT" {
		t.Errorf("Experiences[0].Title = %v, want %v", cfg.Content.Experiences[0].Title, "Cargo PT")
	}
	if cfg.Content.Experiences[0].TitleEN != "Job Title" {
		t.Errorf("Experiences[0].TitleEN = %v, want %v", cfg.Content.Experiences[0].TitleEN, "Job Title")
	}
	if cfg.Content.Experiences[0].Company != "Company" {
		t.Errorf("Experiences[0].Company = %v, want %v", cfg.Content.Experiences[0].Company, "Company")
	}
	if cfg.Content.Experiences[0].Period != "Jan 2020 - Presente" {
		t.Errorf("Experiences[0].Period = %v, want %v", cfg.Content.Experiences[0].Period, "Jan 2020 - Presente")
	}
	if cfg.Content.Experiences[0].PeriodEN != "Jan 2020 - Present" {
		t.Errorf("Experiences[0].PeriodEN = %v, want %v", cfg.Content.Experiences[0].PeriodEN, "Jan 2020 - Present")
	}
	if len(cfg.Content.Experiences[0].Details) != 2 {
		t.Errorf("len(Experiences[0].Details) = %v, want %v", len(cfg.Content.Experiences[0].Details), 2)
	}
	if cfg.Content.Experiences[0].Details[0] != "Detalhe 1" {
		t.Errorf("Experiences[0].Details[0] = %v, want %v", cfg.Content.Experiences[0].Details[0], "Detalhe 1")
	}
	if len(cfg.Content.Experiences[0].DetailsEN) != 2 {
		t.Errorf("len(Experiences[0].DetailsEN) = %v, want %v", len(cfg.Content.Experiences[0].DetailsEN), 2)
	}
	if cfg.Content.Experiences[0].TechStack != "Go" {
		t.Errorf("Experiences[0].TechStack = %v, want %v", cfg.Content.Experiences[0].TechStack, "Go")
	}

	// Verify education
	if len(cfg.Content.Education) != 1 {
		t.Errorf("len(Education) = %v, want %v", len(cfg.Content.Education), 1)
	}
	if cfg.Content.Education[0].Title != "Título PT" {
		t.Errorf("Education[0].Title = %v, want %v", cfg.Content.Education[0].Title, "Título PT")
	}
	if cfg.Content.Education[0].TitleEN != "Title EN" {
		t.Errorf("Education[0].TitleEN = %v, want %v", cfg.Content.Education[0].TitleEN, "Title EN")
	}
	if cfg.Content.Education[0].School != "School" {
		t.Errorf("Education[0].School = %v, want %v", cfg.Content.Education[0].School, "School")
	}
	if cfg.Content.Education[0].Period != "2015-2019" {
		t.Errorf("Education[0].Period = %v, want %v", cfg.Content.Education[0].Period, "2015-2019")
	}
	if cfg.Content.Education[0].PeriodEN != "2015-2019" {
		t.Errorf("Education[0].PeriodEN = %v, want %v", cfg.Content.Education[0].PeriodEN, "2015-2019")
	}

	// Verify technologies
	if len(cfg.Content.Technologies) != 2 {
		t.Errorf("len(Technologies) = %v, want %v", len(cfg.Content.Technologies), 2)
	}
	if cfg.Content.Technologies[0] != "Go" {
		t.Errorf("Technologies[0] = %v, want %v", cfg.Content.Technologies[0], "Go")
	}
	if cfg.Content.Technologies[1] != "Python" {
		t.Errorf("Technologies[1] = %v, want %v", cfg.Content.Technologies[1], "Python")
	}
}

// TestLoad_FileNotFound verifica o comportamento quando o arquivo de configuração não existe.
// TestLoad_FileNotFound verifies the behavior when the configuration file does not exist.
func TestLoad_FileNotFound(t *testing.T) {
	_, err := Load("/nonexistent/path/config.yaml")
	if err == nil {
		t.Fatal("Load() expected error for non-existent file, got nil")
	}
	if !strings.Contains(err.Error(), "failed to read config file") {
		t.Errorf("Load() error = %v, should contain 'failed to read config file'", err)
	}
}

// TestLoad_InvalidYAML verifica o tratamento de erro para conteúdo YAML inválido.
// TestLoad_InvalidYAML verifies error handling for invalid YAML content.
func TestLoad_InvalidYAML(t *testing.T) {
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "config.yaml")

	// Invalid YAML
	invalidYAML := `site: [invalid yaml content: ::::`

	if err := os.WriteFile(configPath, []byte(invalidYAML), 0644); err != nil {
		t.Fatalf("Failed to create test config: %v", err)
	}

	_, err := Load(configPath)
	if err == nil {
		t.Fatal("Load() expected error for invalid YAML, got nil")
	}
	if !strings.Contains(err.Error(), "failed to parse config file") {
		t.Errorf("Load() error = %v, should contain 'failed to parse config file'", err)
	}
}

// TestLoad_EmptyParagraphs verifica o comportamento com conteúdo 'about' vazio.
// TestLoad_EmptyParagraphs verifies behavior with empty 'about' content.
func TestLoad_EmptyParagraphs(t *testing.T) {
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "config.yaml")

	// Config with empty lines in about
	configContent := `site:
  title: "Test"
  description: "Test"
  baseurl: ""
  username: "test"
  user_description: ""
  user_title: ""
  email: ""
content:
  about:
    pt: ""
    en: ""
  experiences: []
  education: []
  technologies: []
`

	if err := os.WriteFile(configPath, []byte(configContent), 0644); err != nil {
		t.Fatalf("Failed to create test config: %v", err)
	}

	cfg, err := Load(configPath)
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}

	// Empty about should result in empty paragraphs
	if len(cfg.Content.ParsedAbout.ParagraphsPT) != 0 {
		t.Errorf("len(ParsedAbout.ParagraphsPT) = %v, want %v", len(cfg.Content.ParsedAbout.ParagraphsPT), 0)
	}
	if len(cfg.Content.ParsedAbout.ParagraphsEN) != 0 {
		t.Errorf("len(ParsedAbout.ParagraphsEN) = %v, want %v", len(cfg.Content.ParsedAbout.ParagraphsEN), 0)
	}
}

// TestLoad_AboutWithEmptyLines verifica que linhas vazias entre parágrafos são ignoradas.
// TestLoad_AboutWithEmptyLines verifies that empty lines between paragraphs are skipped.
func TestLoad_AboutWithEmptyLines(t *testing.T) {
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "config.yaml")

	// Config with empty lines between paragraphs
	configContent := `site:
  title: "Test"
  description: "Test"
  baseurl: ""
  username: "test"
  user_description: ""
  user_title: ""
  email: ""
content:
  about:
    pt: |
      Parágrafo 1.

      Parágrafo 2.


      Parágrafo 3.
    en: ""
  experiences: []
  education: []
  technologies: []
`

	if err := os.WriteFile(configPath, []byte(configContent), 0644); err != nil {
		t.Fatalf("Failed to create test config: %v", err)
	}

	cfg, err := Load(configPath)
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}

	// Should skip empty lines
	if len(cfg.Content.ParsedAbout.ParagraphsPT) != 3 {
		t.Errorf("len(ParsedAbout.ParagraphsPT) = %v, want %v", len(cfg.Content.ParsedAbout.ParagraphsPT), 3)
	}
	if cfg.Content.ParsedAbout.ParagraphsPT[0] != "Parágrafo 1." {
		t.Errorf("ParsedAbout.ParagraphsPT[0] = %v, want %v", cfg.Content.ParsedAbout.ParagraphsPT[0], "Parágrafo 1.")
	}
	if cfg.Content.ParsedAbout.ParagraphsPT[1] != "Parágrafo 2." {
		t.Errorf("ParsedAbout.ParagraphsPT[1] = %v, want %v", cfg.Content.ParsedAbout.ParagraphsPT[1], "Parágrafo 2.")
	}
	if cfg.Content.ParsedAbout.ParagraphsPT[2] != "Parágrafo 3." {
		t.Errorf("ParsedAbout.ParagraphsPT[2] = %v, want %v", cfg.Content.ParsedAbout.ParagraphsPT[2], "Parágrafo 3.")
	}
}

// TestSplitParagraphs testa a função splitParagraphs com múltiplos cenários de entrada.
// TestSplitParagraphs tests the splitParagraphs function with multiple input scenarios.
func TestSplitParagraphs(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "multiple paragraphs",
			input:    "Line 1.\nLine 2.\nLine 3.",
			expected: []string{"Line 1.", "Line 2.", "Line 3."},
		},
		{
			name:     "empty lines",
			input:    "Line 1.\n\nLine 2.\n\n\nLine 3.",
			expected: []string{"Line 1.", "Line 2.", "Line 3."},
		},
		{
			name:     "whitespace trimming",
			input:    "  Line 1.  \n  Line 2.  ",
			expected: []string{"Line 1.", "Line 2."},
		},
		{
			name:     "empty string",
			input:    "",
			expected: []string{},
		},
		{
			name:     "only whitespace",
			input:    "   \n   \n   ",
			expected: []string{},
		},
		{
			name:     "single line",
			input:    "Single line.",
			expected: []string{"Single line."},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := splitParagraphs(tt.input)
			if len(result) != len(tt.expected) {
				t.Errorf("splitParagraphs() = %v, want %v", result, tt.expected)
				return
			}
			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("splitParagraphs()[%d] = %v, want %v", i, result[i], tt.expected[i])
				}
			}
		})
	}
}

// TestCurrentYear verifica se a função CurrentYear retorna o ano atual corretamente.
// TestCurrentYear verifies that CurrentYear returns the current year correctly.
func TestCurrentYear(t *testing.T) {
	expected := time.Now().Year()
	result := CurrentYear()
	if result != expected {
		t.Errorf("CurrentYear() = %v, want %v", result, expected)
	}
}

// TestToTemplateData verifica a conversão de Config para TemplateData com todos os campos.
// TestToTemplateData verifies the conversion from Config to TemplateData with all fields.
func TestToTemplateData(t *testing.T) {
	cfg := &Config{
		Site: SiteConfig{
			Title:           "Test Title",
			Description:     "Test Desc",
			BaseURL:         "https://test.com",
			Username:        "testuser",
			UserDescription: "Test User",
			UserTitle:       "Developer",
			Email:           "test@test.com",
		},
		Content: Content{
			About: AboutContent{
				PT: "Parágrafo PT 1.\n\nParágrafo PT 2.",
				EN: "Paragraph EN 1.\n\nParagraph EN 2.",
			},
			Experiences: []Experience{
				{
					Title:     "Dev",
					TitleEN:     "Developer",
					Company:   "Test Co",
					Period:    "2020-2021",
					PeriodEN:  "2020-2021",
					Details:   []string{"Detail 1"},
					DetailsEN: []string{"Detail EN 1"},
					TechStack: "Go",
				},
			},
			Education: []Education{
				{
					Title:    "Degree",
					TitleEN:  "Degree",
					School:   "University",
					Period:   "2015-2019",
					PeriodEN: "2015-2019",
				},
			},
			Technologies: []string{"Go", "Python"},
		},
	}

	data := cfg.ToTemplateData()

	// Verify Site data is copied
	if data.Site.Title != cfg.Site.Title {
		t.Errorf("TemplateData.Site.Title = %v, want %v", data.Site.Title, cfg.Site.Title)
	}

	// Verify Content data is copied
	if data.Content.About.PT != cfg.Content.About.PT {
		t.Errorf("TemplateData.Content.About.PT = %v, want %v", data.Content.About.PT, cfg.Content.About.PT)
	}

	// Verify ParsedAbout is populated
	if len(data.Content.ParsedAbout.ParagraphsPT) != 2 {
		t.Errorf("len(TemplateData.Content.ParsedAbout.ParagraphsPT) = %v, want %v", len(data.Content.ParsedAbout.ParagraphsPT), 2)
	}
	if data.Content.ParsedAbout.ParagraphsPT[0] != "Parágrafo PT 1." {
		t.Errorf("ParsedAbout.ParagraphsPT[0] = %v, want %v", data.Content.ParsedAbout.ParagraphsPT[0], "Parágrafo PT 1.")
	}
	if len(data.Content.ParsedAbout.ParagraphsEN) != 2 {
		t.Errorf("len(TemplateData.Content.ParsedAbout.ParagraphsEN) = %v, want %v", len(data.Content.ParsedAbout.ParagraphsEN), 2)
	}
	if data.Content.ParsedAbout.ParagraphsEN[0] != "Paragraph EN 1." {
		t.Errorf("ParsedAbout.ParagraphsEN[0] = %v, want %v", data.Content.ParsedAbout.ParagraphsEN[0], "Paragraph EN 1.")
	}

	// Verify Year is set
	expectedYear := time.Now().Year()
	if data.Year != expectedYear {
		t.Errorf("TemplateData.Year = %v, want %v", data.Year, expectedYear)
	}

	// Verify experiences are copied
	if len(data.Content.Experiences) != 1 {
		t.Errorf("len(TemplateData.Content.Experiences) = %v, want %v", len(data.Content.Experiences), 1)
	}
	if data.Content.Experiences[0].Company != "Test Co" {
		t.Errorf("Experiences[0].Company = %v, want %v", data.Content.Experiences[0].Company, "Test Co")
	}

	// Verify education is copied
	if len(data.Content.Education) != 1 {
		t.Errorf("len(TemplateData.Content.Education) = %v, want %v", len(data.Content.Education), 1)
	}
	if data.Content.Education[0].School != "University" {
		t.Errorf("Education[0].School = %v, want %v", data.Content.Education[0].School, "University")
	}

	// Verify technologies are copied
	if len(data.Content.Technologies) != 2 {
		t.Errorf("len(TemplateData.Content.Technologies) = %v, want %v", len(data.Content.Technologies), 2)
	}
	if data.Content.Technologies[0] != "Go" {
		t.Errorf("Technologies[0] = %v, want %v", data.Content.Technologies[0], "Go")
	}
}

// TestToTemplateData_EmptyAbout verifica o comportamento com about vazio (graceful handling).
// TestToTemplateData_EmptyAbout verifies behavior with empty about (graceful handling).
func TestToTemplateData_EmptyAbout(t *testing.T) {
	cfg := &Config{
		Site: SiteConfig{
			Title: "Test",
		},
		Content: Content{
			About: AboutContent{
				PT: "",
				EN: "",
			},
		},
	}

	data := cfg.ToTemplateData()

	// Should handle empty about gracefully
	if len(data.Content.ParsedAbout.ParagraphsPT) != 0 {
		t.Errorf("len(ParsedAbout.ParagraphsPT) = %v, want %v", len(data.Content.ParsedAbout.ParagraphsPT), 0)
	}
	if len(data.Content.ParsedAbout.ParagraphsEN) != 0 {
		t.Errorf("len(ParsedAbout.ParagraphsEN) = %v, want %v", len(data.Content.ParsedAbout.ParagraphsEN), 0)
	}
}

// TestConfigStructs verifica a criação e acesso a todas as structs do pacote config.
// TestConfigStructs verifies the creation and access to all structs in the config package.
func TestConfigStructs(t *testing.T) {
	// Test Experience struct
	exp := Experience{
		Title:     "Dev",
		TitleEN:   "Developer",
		Company:   "Company",
		Period:    "2020-2021",
		PeriodEN:  "2020-2021",
		Details:   []string{"detail"},
		DetailsEN: []string{"detail"},
		TechStack: "Go",
	}
	if exp.Title != "Dev" {
		t.Errorf("Experience.Title = %v, want %v", exp.Title, "Dev")
	}

	// Test Education struct
	edu := Education{
		Title:    "Degree",
		TitleEN:  "Degree",
		School:   "School",
		Period:   "2015-2019",
		PeriodEN: "2015-2019",
	}
	if edu.School != "School" {
		t.Errorf("Education.School = %v, want %v", edu.School, "School")
	}

	// Test AboutContent struct
	about := AboutContent{
		PT: "Portuguese",
		EN: "English",
	}
	if about.PT != "Portuguese" {
		t.Errorf("AboutContent.PT = %v, want %v", about.PT, "Portuguese")
	}

	// Test ParsedAbout struct
	parsed := ParsedAbout{
		ParagraphsPT: []string{"PT"},
		ParagraphsEN: []string{"EN"},
	}
	if len(parsed.ParagraphsPT) != 1 {
		t.Errorf("len(ParsedAbout.ParagraphsPT) = %v, want %v", len(parsed.ParagraphsPT), 1)
	}

	// Test SiteConfig struct
	site := SiteConfig{
		Title:           "Title",
		Description:     "Desc",
		BaseURL:         "https://example.com",
		Username:        "user",
		UserDescription: "description",
		UserTitle:       "title",
		Email:           "email@example.com",
	}
	if site.Email != "email@example.com" {
		t.Errorf("SiteConfig.Email = %v, want %v", site.Email, "email@example.com")
	}

	// Test Content struct
	content := Content{
		About:        about,
		ParsedAbout:  parsed,
		Experiences:  []Experience{exp},
		Education:    []Education{edu},
		Technologies: []string{"Go"},
	}
	if len(content.Experiences) != 1 {
		t.Errorf("len(Content.Experiences) = %v, want %v", len(content.Experiences), 1)
	}
	if len(content.Technologies) != 1 {
		t.Errorf("len(Content.Technologies) = %v, want %v", len(content.Technologies), 1)
	}

	// Test Config struct
	cfg := Config{
		Site:    site,
		Content: content,
	}
	if cfg.Site.Title != "Title" {
		t.Errorf("Config.Site.Title = %v, want %v", cfg.Site.Title, "Title")
	}

	// Test TemplateData struct
	td := TemplateData{
		Site:    site,
		Content: content,
		Year:    2024,
	}
	if td.Year != 2024 {
		t.Errorf("TemplateData.Year = %v, want %v", td.Year, 2024)
	}
}

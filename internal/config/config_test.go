package config

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

// TestLoad_ValidConfig verifica o carregamento de uma configuração YAML válida completa.
func TestLoad_ValidConfig(t *testing.T) {
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
      start_date: "Jan 2020"
      end_date: "Presente"
      start_date_en: "Jan 2020"
      end_date_en: "Present"
      description:
        - "Detalhe 1"
        - "Detalhe 2"
      description_en:
        - "Detail 1"
        - "Detail 2"
      tech_stack: "Go"
  education:
    - degree: "Título PT"
      degree_en: "Title EN"
      institution: "School"
      start_date: "2015-2019"
      start_date_en: "2015-2019"
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

	if len(cfg.Content.ParsedAbout.ParagraphsPT) != 2 {
		t.Errorf("len(ParsedAbout.ParagraphsPT) = %v, want %v", len(cfg.Content.ParsedAbout.ParagraphsPT), 2)
	}
	if cfg.Content.ParsedAbout.ParagraphsPT[0] != "Primeiro parágrafo." {
		t.Errorf("ParsedAbout.ParagraphsPT[0] = %v, want %v", cfg.Content.ParsedAbout.ParagraphsPT[0], "Primeiro parágrafo.")
	}
	if len(cfg.Content.ParsedAbout.ParagraphsEN) != 2 {
		t.Errorf("len(ParsedAbout.ParagraphsEN) = %v, want %v", len(cfg.Content.ParsedAbout.ParagraphsEN), 2)
	}

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
	if cfg.Content.Experiences[0].StartDate != "Jan 2020" {
		t.Errorf("Experiences[0].StartDate = %v, want %v", cfg.Content.Experiences[0].StartDate, "Jan 2020")
	}
	if cfg.Content.Experiences[0].EndDate != "Presente" {
		t.Errorf("Experiences[0].EndDate = %v, want %v", cfg.Content.Experiences[0].EndDate, "Presente")
	}
	if cfg.Content.Experiences[0].StartDateEN != "Jan 2020" {
		t.Errorf("Experiences[0].StartDateEN = %v, want %v", cfg.Content.Experiences[0].StartDateEN, "Jan 2020")
	}
	if cfg.Content.Experiences[0].EndDateEN != "Present" {
		t.Errorf("Experiences[0].EndDateEN = %v, want %v", cfg.Content.Experiences[0].EndDateEN, "Present")
	}
	if len(cfg.Content.Experiences[0].Description) != 2 {
		t.Errorf("len(Experiences[0].Description) = %v, want %v", len(cfg.Content.Experiences[0].Description), 2)
	}
	if cfg.Content.Experiences[0].Description[0] != "Detalhe 1" {
		t.Errorf("Experiences[0].Description[0] = %v, want %v", cfg.Content.Experiences[0].Description[0], "Detalhe 1")
	}
	if len(cfg.Content.Experiences[0].DescriptionEN) != 2 {
		t.Errorf("len(Experiences[0].DescriptionEN) = %v, want %v", len(cfg.Content.Experiences[0].DescriptionEN), 2)
	}
	if cfg.Content.Experiences[0].TechStack != "Go" {
		t.Errorf("Experiences[0].TechStack = %v, want %v", cfg.Content.Experiences[0].TechStack, "Go")
	}

	if len(cfg.Content.Education) != 1 {
		t.Errorf("len(Education) = %v, want %v", len(cfg.Content.Education), 1)
	}
	if cfg.Content.Education[0].Degree != "Título PT" {
		t.Errorf("Education[0].Degree = %v, want %v", cfg.Content.Education[0].Degree, "Título PT")
	}
	if cfg.Content.Education[0].DegreeEN != "Title EN" {
		t.Errorf("Education[0].DegreeEN = %v, want %v", cfg.Content.Education[0].DegreeEN, "Title EN")
	}
	if cfg.Content.Education[0].Institution != "School" {
		t.Errorf("Education[0].Institution = %v, want %v", cfg.Content.Education[0].Institution, "School")
	}
	if cfg.Content.Education[0].StartDate != "2015-2019" {
		t.Errorf("Education[0].StartDate = %v, want %v", cfg.Content.Education[0].StartDate, "2015-2019")
	}
	if cfg.Content.Education[0].StartDateEN != "2015-2019" {
		t.Errorf("Education[0].StartDateEN = %v, want %v", cfg.Content.Education[0].StartDateEN, "2015-2019")
	}

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

func TestLoad_FileNotFound(t *testing.T) {
	_, err := Load("/nonexistent/path/config.yaml")
	if err == nil {
		t.Fatal("Load() expected error for non-existent file, got nil")
	}
	if !strings.Contains(err.Error(), "failed to read config file") {
		t.Errorf("Load() error = %v, should contain 'failed to read config file'", err)
	}
}

func TestLoad_InvalidYAML(t *testing.T) {
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "config.yaml")

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

func TestLoad_EmptyParagraphs(t *testing.T) {
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "config.yaml")

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

	if len(cfg.Content.ParsedAbout.ParagraphsPT) != 0 {
		t.Errorf("len(ParsedAbout.ParagraphsPT) = %v, want %v", len(cfg.Content.ParsedAbout.ParagraphsPT), 0)
	}
	if len(cfg.Content.ParsedAbout.ParagraphsEN) != 0 {
		t.Errorf("len(ParsedAbout.ParagraphsEN) = %v, want %v", len(cfg.Content.ParsedAbout.ParagraphsEN), 0)
	}
}

func TestLoad_AboutWithEmptyLines(t *testing.T) {
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "config.yaml")

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

func TestCurrentYear(t *testing.T) {
	expected := time.Now().Year()
	result := CurrentYear()
	if result != expected {
		t.Errorf("CurrentYear() = %v, want %v", result, expected)
	}
}

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
					Title:         "Dev",
					TitleEN:       "Developer",
					Company:       "Test Co",
					StartDate:     "2020-2021",
					StartDateEN:   "2020-2021",
					Description:   []string{"Detail 1"},
					DescriptionEN: []string{"Detail EN 1"},
					TechStack:     "Go",
				},
			},
			Education: []Education{
				{
					Degree:      "Degree",
					DegreeEN:    "Degree",
					Institution: "University",
					StartDate:   "2015-2019",
					StartDateEN: "2015-2019",
				},
			},
			Technologies: []string{"Go", "Python"},
		},
	}

	data := cfg.ToTemplateData()

	if data.Site.Title != cfg.Site.Title {
		t.Errorf("TemplateData.Site.Title = %v, want %v", data.Site.Title, cfg.Site.Title)
	}
	if data.Content.About.PT != cfg.Content.About.PT {
		t.Errorf("TemplateData.Content.About.PT = %v, want %v", data.Content.About.PT, cfg.Content.About.PT)
	}
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
	expectedYear := time.Now().Year()
	if data.Year != expectedYear {
		t.Errorf("TemplateData.Year = %v, want %v", data.Year, expectedYear)
	}
	if len(data.Content.Experiences) != 1 {
		t.Errorf("len(TemplateData.Content.Experiences) = %v, want %v", len(data.Content.Experiences), 1)
	}
	if data.Content.Experiences[0].Company != "Test Co" {
		t.Errorf("Experiences[0].Company = %v, want %v", data.Content.Experiences[0].Company, "Test Co")
	}
	if len(data.Content.Education) != 1 {
		t.Errorf("len(TemplateData.Content.Education) = %v, want %v", len(data.Content.Education), 1)
	}
	if data.Content.Education[0].Institution != "University" {
		t.Errorf("Education[0].Institution = %v, want %v", data.Content.Education[0].Institution, "University")
	}
	if len(data.Content.Technologies) != 2 {
		t.Errorf("len(TemplateData.Content.Technologies) = %v, want %v", len(data.Content.Technologies), 2)
	}
	if data.Content.Technologies[0] != "Go" {
		t.Errorf("Technologies[0] = %v, want %v", data.Content.Technologies[0], "Go")
	}
}

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

	if len(data.Content.ParsedAbout.ParagraphsPT) != 0 {
		t.Errorf("len(ParsedAbout.ParagraphsPT) = %v, want %v", len(data.Content.ParsedAbout.ParagraphsPT), 0)
	}
	if len(data.Content.ParsedAbout.ParagraphsEN) != 0 {
		t.Errorf("len(ParsedAbout.ParagraphsEN) = %v, want %v", len(data.Content.ParsedAbout.ParagraphsEN), 0)
	}
}

func TestConfigStructs(t *testing.T) {
	exp := Experience{
		Title:         "Dev",
		TitleEN:       "Developer",
		Company:       "Company",
		StartDate:     "2020-2021",
		StartDateEN:   "2020-2021",
		Description:   []string{"detail"},
		DescriptionEN: []string{"detail"},
		TechStack:     "Go",
	}
	if exp.Title != "Dev" {
		t.Errorf("Experience.Title = %v, want %v", exp.Title, "Dev")
	}

	edu := Education{
		Degree:      "Degree",
		DegreeEN:    "Degree",
		Institution: "School",
		StartDate:   "2015-2019",
		StartDateEN: "2015-2019",
	}
	if edu.Institution != "School" {
		t.Errorf("Education.Institution = %v, want %v", edu.Institution, "School")
	}

	about := AboutContent{
		PT: "Portuguese",
		EN: "English",
	}
	if about.PT != "Portuguese" {
		t.Errorf("AboutContent.PT = %v, want %v", about.PT, "Portuguese")
	}

	parsed := ParsedAbout{
		ParagraphsPT: []string{"PT"},
		ParagraphsEN: []string{"EN"},
	}
	if len(parsed.ParagraphsPT) != 1 {
		t.Errorf("len(ParsedAbout.ParagraphsPT) = %v, want %v", len(parsed.ParagraphsPT), 1)
	}

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

	cfg := Config{
		Site:    site,
		Content: content,
	}
	if cfg.Site.Title != "Title" {
		t.Errorf("Config.Site.Title = %v, want %v", cfg.Site.Title, "Title")
	}

	td := TemplateData{
		Site:    site,
		Content: content,
		Year:    2024,
	}
	if td.Year != 2024 {
		t.Errorf("TemplateData.Year = %v, want %v", td.Year, 2024)
	}
}

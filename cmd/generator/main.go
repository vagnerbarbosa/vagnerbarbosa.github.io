// Generator builds the static site from templates and config
package main

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"github.com/vagnerbarbosa/vagnerbarbosa.github.io/internal/config"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	// Load configuration
	cfg, err := config.Load("config.yaml")
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	// Create output directory
	if err := os.MkdirAll("public", 0755); err != nil {
		return fmt.Errorf("failed to create public directory: %w", err)
	}

	// Parse templates
	tmpl, err := parseTemplates()
	if err != nil {
		return fmt.Errorf("failed to parse templates: %w", err)
	}

	// Generate index.html
	data := cfg.ToTemplateData()
	if err := generateIndex(tmpl, data); err != nil {
		return fmt.Errorf("failed to generate index: %w", err)
	}

	// Copy static assets
	if err := copyAssets(); err != nil {
		return fmt.Errorf("failed to copy assets: %w", err)
	}

	// Copy CNAME if exists
	if err := copyCNAME(); err != nil {
		return fmt.Errorf("failed to copy CNAME: %w", err)
	}

	fmt.Println("Site generated successfully in public/")
	return nil
}

func parseTemplates() (*template.Template, error) {
	// Create template with custom functions
	funcMap := template.FuncMap{
		"safeHTML": func(s string) template.HTML {
			return template.HTML(s)
		},
	}

	tmpl := template.New("").Funcs(funcMap)

	// Walk templates directory
	err := filepath.Walk("templates", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if !strings.HasSuffix(path, ".html") {
			return nil
		}

		content, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("failed to read template %s: %w", path, err)
		}

		_, err = tmpl.Parse(string(content))
		if err != nil {
			return fmt.Errorf("failed to parse template %s: %w", path, err)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return tmpl, nil
}

func generateIndex(tmpl *template.Template, data config.TemplateData) error {
	file, err := os.Create("public/index.html")
	if err != nil {
		return fmt.Errorf("failed to create index.html: %w", err)
	}
	defer file.Close()

	// Execute the "index" template
	if err := tmpl.ExecuteTemplate(file, "index", data); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	return nil
}

func copyAssets() error {
	// Copy assets directory
	if err := copyDir("assets", "public/assets"); err != nil {
		return fmt.Errorf("failed to copy assets directory: %w", err)
	}
	return nil
}

func copyCNAME() error {
	// Check if CNAME exists
	if _, err := os.Stat("CNAME"); os.IsNotExist(err) {
		return nil // CNAME doesn't exist, skip
	}

	content, err := os.ReadFile("CNAME")
	if err != nil {
		return fmt.Errorf("failed to read CNAME: %w", err)
	}

	if err := os.WriteFile("public/CNAME", content, 0644); err != nil {
		return fmt.Errorf("failed to write CNAME: %w", err)
	}

	return nil
}

// copyDir recursively copies a directory
func copyDir(src, dst string) error {
	// Get file info
	info, err := os.Stat(src)
	if err != nil {
		return err
	}

	// Create destination directory
	if err := os.MkdirAll(dst, info.Mode()); err != nil {
		return err
	}

	// Read directory contents
	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			if err := copyDir(srcPath, dstPath); err != nil {
				return err
			}
		} else {
			if err := copyFile(srcPath, dstPath); err != nil {
				return err
			}
		}
	}

	return nil
}

// copyFile copies a single file
func copyFile(src, dst string) error {
	content, err := os.ReadFile(src)
	if err != nil {
		return err
	}

	info, err := os.Stat(src)
	if err != nil {
		return err
	}

	return os.WriteFile(dst, content, info.Mode())
}

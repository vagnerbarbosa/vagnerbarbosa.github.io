// Generator builds the static site from templates and config
package main

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/html"
	"github.com/tdewolff/minify/v2/js"
	"github.com/vagnerbarbosa/vagnerbarbosa.github.io/internal/config"
)

// minifier instance configured for optimal compression
var minifier *minify.M

func init() {
	minifier = minify.New()

	// JavaScript: aggressive minification
	minifier.AddFunc("application/javascript", js.Minify)
	minifier.AddFunc("text/javascript", js.Minify)

	// CSS: clean minification preserving semantics
	minifier.AddFunc("text/css", css.Minify)

	// HTML: minify while keeping essential whitespace
	minifier.Add("text/html", &html.Minifier{
		KeepDocumentTags: true,
		KeepEndTags:      true,
		KeepQuotes:       true,
		KeepWhitespace:   false,
	})
}

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

	// Generate index.html (minified)
	data := cfg.ToTemplateData()
	if err := generateIndex(tmpl, data); err != nil {
		return fmt.Errorf("failed to generate index: %w", err)
	}

	// Copy and minify static assets
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
	// Execute template to buffer first
	var buf bytes.Buffer
	if err := tmpl.ExecuteTemplate(&buf, "index", data); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	// Minify HTML before writing
	minified, err := minifier.Bytes("text/html", buf.Bytes())
	if err != nil {
		// Fallback to unminified if minification fails
		minified = buf.Bytes()
	}

	if err := os.WriteFile("public/index.html", minified, 0644); err != nil {
		return fmt.Errorf("failed to write index.html: %w", err)
	}

	// Print stats
	savings := len(buf.Bytes()) - len(minified)
	if savings > 0 {
		percent := float64(savings) * 100 / float64(len(buf.Bytes()))
		fmt.Printf("HTML minified: %d bytes -> %d bytes (saved %d bytes, %.1f%%)\n",
			len(buf.Bytes()), len(minified), savings, percent)
	}

	return nil
}

func copyAssets() error {
	if err := copyAndMinifyDir("assets", "public/assets"); err != nil {
		return fmt.Errorf("failed to copy assets directory: %w", err)
	}
	return nil
}

func copyCNAME() error {
	if _, err := os.Stat("CNAME"); os.IsNotExist(err) {
		return nil
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

// copyAndMinifyDir recursively copies a directory, minifying supported file types
func copyAndMinifyDir(src, dst string) error {
	info, err := os.Stat(src)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(dst, info.Mode()); err != nil {
		return err
	}

	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			if err := copyAndMinifyDir(srcPath, dstPath); err != nil {
				return err
			}
		} else {
			if err := copyAndMinifyFile(srcPath, dstPath); err != nil {
				return err
			}
		}
	}

	return nil
}

// copyAndMinifyFile copies a file, minifying if it's a supported type
func copyAndMinifyFile(src, dst string) error {
	content, err := os.ReadFile(src)
	if err != nil {
		return err
	}

	info, err := os.Stat(src)
	if err != nil {
		return err
	}

	// Determine mime type and minify if supported
	ext := strings.ToLower(filepath.Ext(src))
	var mimeType string

	switch ext {
	case ".js":
		mimeType = "application/javascript"
	case ".css":
		mimeType = "text/css"
	case ".html", ".htm":
		mimeType = "text/html"
	}

	if mimeType != "" {
		minified, err := minifier.Bytes(mimeType, content)
		if err == nil && len(minified) < len(content) {
			savings := len(content) - len(minified)
			percent := float64(savings) * 100 / float64(len(content))
			relPath, _ := filepath.Rel("assets", src)
			fmt.Printf("Minified %s: %d -> %d bytes (saved %d bytes, %.1f%%)\n",
				relPath, len(content), len(minified), savings, percent)
			content = minified
		}
	}

	return os.WriteFile(dst, content, info.Mode())
}

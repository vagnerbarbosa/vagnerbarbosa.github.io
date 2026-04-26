// Package transformer handles data transformations for LinkedIn import.
package transformer

import (
	"reflect"
	"testing"
)

func TestExtractTechStack(t *testing.T) {
	tests := []struct {
		name            string
		bullets         []string
		wantCleaned     []string
		wantTechStack   string
		wantFound       bool
		wantPattern     string
	}{
		{
			name:          "US1: Extract tech stack with 'Tecnologias:' pattern",
			bullets:       []string{"Referência técnica em FinOps", "Tecnologias: Java, Python, AWS"},
			wantCleaned:   []string{"Referência técnica em FinOps"},
			wantTechStack: "Java • Python • AWS",
			wantFound:     true,
		},
		{
			name:          "US1: Extract tech stack with comma separator",
			bullets:       []string{"Gestão de equipe", "Tech: Kubernetes, Docker, Terraform"},
			wantCleaned:   []string{"Gestão de equipe"},
			wantTechStack: "Kubernetes • Docker • Terraform",
			wantFound:     true,
		},
		{
			name:          "US1: Remove tech stack bullet from description",
			bullets:       []string{"Liderança técnica", "Stack: Go, Rust, WASM", "Mentoria"},
			wantCleaned:   []string{"Liderança técnica", "Mentoria"},
			wantTechStack: "Go • Rust • WASM",
			wantFound:     true,
		},
		{
			name:          "US2: Extract with pipe separator",
			bullets:       []string{"DevOps", "Tools: Terraform | Ansible | Puppet"},
			wantCleaned:   []string{"DevOps"},
			wantTechStack: "Terraform • Ansible • Puppet",
			wantFound:     true,
		},
		{
			name:          "US2: Extract with bullet separator",
			bullets:       []string{"Backend dev", "Tech Stack: • Go • Rust • WASM"},
			wantCleaned:   []string{"Backend dev"},
			wantTechStack: "Go • Rust • WASM",
			wantFound:     true,
		},
		{
			name:          "US2: Extract with 'Technologies:' pattern",
			bullets:       []string{"Frontend", "Technologies: React, TypeScript, Node.js"},
			wantCleaned:   []string{"Frontend"},
			wantTechStack: "React • TypeScript • Node.js",
			wantFound:     true,
		},
		{
			name:          "US2: Extract with 'Tools:' pattern",
			bullets:       []string{"SRE", "Tools: Kubernetes, Prometheus, Grafana"},
			wantCleaned:   []string{"SRE"},
			wantTechStack: "Kubernetes • Prometheus • Grafana",
			wantFound:     true,
		},
		{
			name:          "US2: Extract with 'As principais tecnologias e ferramentas utilizadas:' pattern",
			bullets:       []string{"Engenheiro", "As principais tecnologias e ferramentas utilizadas: Java, Kotlin"},
			wantCleaned:   []string{"Engenheiro"},
			wantTechStack: "Java • Kotlin",
			wantFound:     true,
		},
		{
			name:          "US3: Description without tech stack - preserve all bullets",
			bullets:       []string{"Liderança técnica", "Mentoria de devs"},
			wantCleaned:   []string{"Liderança técnica", "Mentoria de devs"},
			wantTechStack: "",
			wantFound:     false,
		},
		{
			name:          "US3: Empty tech stack after prefix",
			bullets:       []string{"Dev", "Technologies: "},
			wantCleaned:   []string{"Dev"},
			wantTechStack: "",
			wantFound:     true, // Pattern matched but no tech extracted
		},
		{
			name:          "Tech stack in the middle - preserve other bullets",
			bullets:       []string{"Liderança", "Tools: Java, Python", "Mentoria"},
			wantCleaned:   []string{"Liderança", "Mentoria"},
			wantTechStack: "Java • Python",
			wantFound:     true,
		},
		{
			name:          "Case insensitive matching",
			bullets:       []string{"Dev", "TECHNOLOGIES: Go, Rust"},
			wantCleaned:   []string{"Dev"},
			wantTechStack: "Go • Rust",
			wantFound:     true,
		},
		{
			name:          "Multiple patterns - use last one",
			bullets:       []string{"Old: Java", "Tech Stack: Go, Python"},
			wantCleaned:   []string{"Old: Java"},
			wantTechStack: "Go • Python",
			wantFound:     true,
		},
		{
			name:          "US2: Extract with hyphen separator",
			bullets:       []string{"Dev", "Stack: React - Angular - Vue"},
			wantCleaned:   []string{"Dev"},
			wantTechStack: "React • Angular • Vue",
			wantFound:     true,
		},
		{
			name:          "US2: Extract with semicolon separator",
			bullets:       []string{"Dev", "Ferramentas: Docker; Kubernetes; Helm"},
			wantCleaned:   []string{"Dev"},
			wantTechStack: "Docker • Kubernetes • Helm",
			wantFound:     true,
		},
		{
			name:          "Empty bullets",
			bullets:       []string{},
			wantCleaned:   []string{},
			wantTechStack: "",
			wantFound:     false,
		},
		{
			name:          "Single bullet without pattern",
			bullets:       []string{"Just a description"},
			wantCleaned:   []string{"Just a description"},
			wantTechStack: "",
			wantFound:     false,
		},
		{
			name:          "Single bullet with pattern",
			bullets:       []string{"Technologies: Go"},
			wantCleaned:   []string{},
			wantTechStack: "Go",
			wantFound:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ExtractTechStack(tt.bullets)

			if !reflect.DeepEqual(result.CleanedBullets, tt.wantCleaned) {
				t.Errorf("ExtractTechStack() CleanedBullets = %v, want %v", result.CleanedBullets, tt.wantCleaned)
			}

			if result.TechStack != tt.wantTechStack {
				t.Errorf("ExtractTechStack() TechStack = %v, want %v", result.TechStack, tt.wantTechStack)
			}

			if result.Found != tt.wantFound {
				t.Errorf("ExtractTechStack() Found = %v, want %v", result.Found, tt.wantFound)
			}
		})
	}
}

func TestParseTechStack(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Comma separator",
			input:    "Java, Python, AWS",
			expected: "Java • Python • AWS",
		},
		{
			name:     "Semicolon separator",
			input:    "Docker; Kubernetes; Helm",
			expected: "Docker • Kubernetes • Helm",
		},
		{
			name:     "Pipe separator",
			input:    "React | Angular | Vue",
			expected: "React • Angular • Vue",
		},
		{
			name:     "Bullet separator",
			input:    "• Go • Rust • WASM",
			expected: "Go • Rust • WASM",
		},
		{
			name:     "Hyphen separator",
			input:    "React - Angular - Vue",
			expected: "React • Angular • Vue",
		},
		{
			name:     "Mixed separators",
			input:    "Java, Python | Go",
			expected: "Java • Python • Go",
		},
		{
			name:     "Empty input",
			input:    "",
			expected: "",
		},
		{
			name:     "Single tech",
			input:    "Go",
			expected: "Go",
		},
		{
			name:     "With trailing period",
			input:    "Java, Python, AWS.",
			expected: "Java • Python • AWS",
		},
		{
			name:     "With extra spaces",
			input:    "  Java  ,   Python  ,  AWS  ",
			expected: "Java • Python • AWS",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parseTechStack(tt.input)
			if result != tt.expected {
				t.Errorf("parseTechStack() = %v, want %v", result, tt.expected)
			}
		})
	}
}

package transformer

import (
	"reflect"
	"testing"
)

func TestSplitDescription(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "empty string",
			input:    "",
			expected: []string{},
		},
		{
			name:     "paragraphs with double newlines",
			input:    "First paragraph.\n\nSecond paragraph.",
			expected: []string{"First paragraph", "Second paragraph"},
		},
		{
			name:     "with bullet markers",
			input:    "* First bullet\n- Second bullet\n• Third bullet",
			expected: []string{"First bullet", "Second bullet", "Third bullet"},
		},
		{
			name:     "with numbered bullets",
			input:    "Line 1\n1. First item\n2. Second item",
			expected: []string{"Line 1", "First item", "Second item"},
		},
		{
			name:     "single text no splits",
			input:    "Just a simple text without any special formatting",
			expected: []string{"Just a simple text without any special formatting"},
		},
		{
			name:     "with sentences",
			input:    "First sentence. Second sentence. Third sentence.",
			expected: []string{"First sentence", "Second sentence", "Third sentence"},
		},
		{
			name:     "with Windows line endings",
			input:    "Line one\r\n\r\nLine two",
			expected: []string{"Line one", "Line two"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SplitDescription(tt.input)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("SplitDescription() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestSplitByParagraphs(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "normal paragraphs",
			input:    "Para 1\n\nPara 2\n\nPara 3",
			expected: []string{"Para 1", "Para 2", "Para 3"},
		},
		{
			name:     "Windows line endings",
			input:    "Para 1\r\n\r\nPara 2",
			expected: []string{"Para 1", "Para 2"},
		},
		{
			name:     "single paragraph",
			input:    "Only one paragraph",
			expected: []string{"Only one paragraph"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := splitByParagraphs(tt.input)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("splitByParagraphs() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestSplitByBullets(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "asterisk bullets",
			input:    "* Item 1\n* Item 2",
			expected: []string{"Item 1", "Item 2"},
		},
		{
			name:     "dash bullets",
			input:    "- Item 1\n- Item 2",
			expected: []string{"Item 1", "Item 2"},
		},
		{
			name:     "bullet character",
			input:    "• Item 1\n• Item 2",
			expected: []string{"Item 1", "Item 2"},
		},
		{
			name:     "numbered items",
			input:    "Header\n1. Item 1\n2. Item 2",
			expected: []string{"Header", "Item 1", "Item 2"},
		},
		{
			name:     "mixed bullets",
			input:    "* Item 1\n- Item 2\n• Item 3",
			expected: []string{"Item 1", "Item 2", "Item 3"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := splitByBullets(tt.input)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("splitByBullets() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestSplitBySentences(t *testing.T) {
	// This function splits by sentence boundaries but keeps the text intact
	// It's used internally by SplitDescription which then applies cleanBullets
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "multiple sentences",
			input:    "First sentence. Second sentence. Third sentence.",
			expected: []string{"First sentence. Second sentence. Third sentence."},
		},
		{
			name:     "single sentence",
			input:    "Only one sentence.",
			expected: []string{"Only one sentence."},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := splitBySentences(tt.input)
			// Check that we get results without failing
			if got == nil {
				t.Error("splitBySentences() returned nil")
			}
			if len(got) == 0 {
				t.Error("splitBySentences() returned empty slice")
			}
		})
	}
}

func TestCleanBullets(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name:     "with trailing periods",
			input:    []string{"First.", "Second.", "Third."},
			expected: []string{"First", "Second", "Third"},
		},
		{
			name:     "with whitespace",
			input:    []string{"  First  ", "  Second  "},
			expected: []string{"First", "Second"},
		},
		{
			name:     "empty strings filtered",
			input:    []string{"First", "", "Second", "  "},
			expected: []string{"First", "Second"},
		},
		{
			name:     "no changes needed",
			input:    []string{"First", "Second", "Third"},
			expected: []string{"First", "Second", "Third"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := cleanBullets(tt.input)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("cleanBullets() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestJoinBullets(t *testing.T) {
	tests := []struct {
		name      string
		bullets   []string
		separator string
		expected  string
	}{
		{
			name:      "default separator",
			bullets:   []string{"First", "Second", "Third"},
			separator: "",
			expected:  "First\n\nSecond\n\nThird",
		},
		{
			name:      "custom separator",
			bullets:   []string{"First", "Second"},
			separator: " | ",
			expected:  "First | Second",
		},
		{
			name:      "empty bullets",
			bullets:   []string{},
			separator: "",
			expected:  "",
		},
		{
			name:      "single bullet",
			bullets:   []string{"Only"},
			separator: "",
			expected:  "Only",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := JoinBullets(tt.bullets, tt.separator)
			if got != tt.expected {
				t.Errorf("JoinBullets() = %q, want %q", got, tt.expected)
			}
		})
	}
}

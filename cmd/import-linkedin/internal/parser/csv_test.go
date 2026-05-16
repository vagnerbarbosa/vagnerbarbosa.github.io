package parser

import (
	"os"
	"strings"
	"testing"
)

func TestNewCSVReaderFromIO(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{"valid csv", "Col1,Col2\nVal1,Val2", false},
		{"empty input", "", true},
		{"only header", "Col1,Col2", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := strings.NewReader(tt.input)
			_, err := NewCSVReaderFromIO(r)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewCSVReaderFromIO() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCSVReader_Next(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []map[string]string
	}{
		{
			"regular records",
			"Name,Age\nAlice,30\nBob,25",
			[]map[string]string{
				{"Name": "Alice", "Age": "30"},
				{"Name": "Bob", "Age": "25"},
			},
		},
		{
			"irregular records (too short)",
			"Name,Age,City\nAlice,30,NY\nBob,25",
			[]map[string]string{
				{"Name": "Alice", "Age": "30", "City": "NY"},
				{"Name": "Bob", "Age": "25", "City": ""},
			},
		},
		{
			"irregular records (too long)",
			"Name,Age\nAlice,30,Extra",
			[]map[string]string{
				{"Name": "Alice", "Age": "30"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := strings.NewReader(tt.input)
			reader, err := NewCSVReaderFromIO(r)
			if err != nil {
				t.Fatalf("failed to create CSV reader: %v", err)
			}

			var results []map[string]string
			for {
				row, err := reader.Next()
				if err != nil {
					break
				}
				results = append(results, row)
			}

			if len(results) != len(tt.expected) {
				t.Errorf("got %d records, want %d", len(results), len(tt.expected))
			}

			for i := range results {
				for k, v := range tt.expected[i] {
					if results[i][k] != v {
						t.Errorf("record %d column %s = %q, want %q", i, k, results[i][k], v)
					}
				}
			}
		})
	}
}

func TestCSVReader_GetHeader(t *testing.T) {
	input := "Name,Age\nAlice,30"
	r := strings.NewReader(input)
	reader, _ := NewCSVReaderFromIO(r)
	header := reader.GetHeader()

	expected := map[string]int{"Name": 0, "Age": 1}
	if len(header) != len(expected) {
		t.Errorf("header length = %d, want %d", len(header), len(expected))
	}
	for k, v := range expected {
		if header[k] != v {
			t.Errorf("header[%s] = %d, want %d", k, header[k], v)
		}
	}
}

func TestCSVReader_HasColumn(t *testing.T) {
	input := "Name,Age\nAlice,30"
	r := strings.NewReader(input)
	reader, _ := NewCSVReaderFromIO(r)

	if !reader.HasColumn("Name") {
		t.Error("expected HasColumn(\"Name\") to be true")
	}
	if reader.HasColumn("City") {
		t.Error("expected HasColumn(\"City\") to be false")
	}
}

func TestCSVReader_Close(t *testing.T) {
	input := "Name,Age\nAlice,30"
	r := strings.NewReader(input)
	reader, _ := NewCSVReaderFromIO(r)
	if err := reader.Close(); err != nil {
		t.Errorf("Close() error = %v", err)
	}
}

func TestCSVReader_ValidateColumns(t *testing.T) {
	input := "Name,Age\nAlice,30"
	r := strings.NewReader(input)
	reader, _ := NewCSVReaderFromIO(r)

	tests := []struct {
		name     string
		required []string
		expected []string
	}{
		{"all present", []string{"Name", "Age"}, []string(nil)},
		{"some missing", []string{"Name", "City"}, []string{"City"}},
		{"all missing", []string{"City", "Country"}, []string{"City", "Country"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			missing := reader.ValidateColumns(tt.required)
			if len(missing) != len(tt.expected) {
				t.Errorf("ValidateColumns() missing = %v, want %v", missing, tt.expected)
			}
		})
	}
}

func TestNewCSVReader(t *testing.T) {
	content := "Name,Age\nAlice,30"
	tmpFile, err := os.CreateTemp("", "testcsv*.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.Write([]byte(content)); err != nil {
		t.Fatal(err)
	}
	tmpFile.Close()

	t.Run("valid file", func(t *testing.T) {
		_, err := NewCSVReader(tmpFile.Name())
		if err != nil {
			t.Errorf("NewCSVReader() error = %v", err)
		}
	})

	t.Run("non-existent file", func(t *testing.T) {
		_, err := NewCSVReader("non_existent_file.csv")
		if err == nil {
			t.Error("expected error for non-existent file")
		}
	})
}

package integration

import (
	"fmt"
	"os"
	"reflect"

	"gopkg.in/yaml.v3"
)

// CompareYAML performs a semantic comparison of two YAML files by unmarshaling them
// into generic maps and using reflect.DeepEqual. This avoids failures due to
// key ordering or insignificant whitespace differences.
func CompareYAML(current, golden []byte) (bool, error) {
	var currentMap, goldenMap interface{}

	if err := yaml.Unmarshal(current, &currentMap); err != nil {
		return false, fmt.Errorf("failed to unmarshal current YAML: %w", err)
	}

	if err := yaml.Unmarshal(golden, &goldenMap); err != nil {
		return false, fmt.Errorf("failed to unmarshal golden YAML: %w", err)
	}

	if reflect.DeepEqual(currentMap, goldenMap) {
		return true, nil
	}

	return false, nil
}

// UpdateGoldenFile overwrites the golden reference file with the current output.
// This is used when the changes in the YAML output are intentional and should
// become the new baseline.
func UpdateGoldenFile(path string, content []byte) error {
	return os.WriteFile(path, content, 0644)
}

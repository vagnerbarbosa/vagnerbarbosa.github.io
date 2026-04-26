// LinkedIn Import CLI
//
// A CLI tool to import professional data from LinkedIn CSV exports
// and merge them into a YAML configuration file.
//
// Usage:
//
//	linkedin-import import       # Import all data with interactive confirmation
//	linkedin-import import --yes # Auto-accept all changes
//	linkedin-import validate     # Validate CSV files
//	linkedin-import version      # Show version
//
// For more information, see: https://github.com/vagnerbarbosa/vagnerbarbosa.github.io
package main

import (
	"os"

	"github.com/vagnerbarbosa/vagnerbarbosa.github.io/cmd/import-linkedin/commands"
)

func main() {
	// Execute command and exit with appropriate code
	os.Exit(commands.Execute(os.Args[1:]))
}

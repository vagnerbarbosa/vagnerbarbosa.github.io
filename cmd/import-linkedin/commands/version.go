// Package commands implements the CLI commands for LinkedIn import.
package commands

import "fmt"

const (
	// Version is the current version of the CLI.
	Version = "1.0.0"
	// BuildDate is the date when the binary was built.
	BuildDate = "2026-04-26"
)

// printVersion prints version information.
// This is called from root.go.
func printVersionExtended() {
	fmt.Println("LinkedIn Import CLI")
	fmt.Printf("Version: %s\n", Version)
	fmt.Printf("Build Date: %s\n", BuildDate)
	fmt.Println("Go version: go1.26.2")
	fmt.Println()
	fmt.Println("For help, run: linkedin-import help")
}

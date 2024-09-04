/*
Copyright © 2024 Mauro Perna
*/

package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new [app name]",
	Short: "Creates a new Go Nest application",
	Long:  "This command generates a new Go Nest application with a basic structure (main.go, go.mod, and common structure).",
	Args:  cobra.ExactArgs(1), // Ensures that one argument is provided
	Run: func(cmd *cobra.Command, args []string) {
		appName := args[0]
		createApp(appName)
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}

// createApp scaffolds the new application structure
func createApp(appName string) {
	// Create the directory for the new app
	if err := os.Mkdir(appName, os.ModePerm); err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}

	// Navigate to the new directory
	if err := os.Chdir(appName); err != nil {
		fmt.Println("Error navigating to app directory:", err)
		return
	}

	// Initialize Go module
	cmd := exec.Command("go", "mod", "init", appName)
	if output, err := cmd.CombinedOutput(); err != nil {
		fmt.Println("Error initializing Go module:", err)
		fmt.Println(string(output))
		return
	}

	// Create basic app structure
	createFile("main.go", mainGoContent)
	createFile("common/README.md", commonReadmeContent)
	createFile("domain/README.md", domainReadmeContent)
	createFile("service/README.md", serviceReadmeContent)

	fmt.Printf("New Go Nest app '%s' created successfully!\n", appName)
}

// createFile creates a file with the given content
func createFile(path, content string) {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		fmt.Println("Error creating directories:", err)
		return
	}

	file, err := os.Create(path)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
}

var mainGoContent = `package main

import "fmt"

func main() {
    fmt.Println("Welcome to your Go Nest app!")
}
`

var commonReadmeContent = `# Common
This folder will contain common utilities, enums, and middlewares.
`

var domainReadmeContent = `# Domain
This folder will contain the core business logic of your application.
`

var serviceReadmeContent = `# Service
This folder will contain services like database connections and external APIs.
`

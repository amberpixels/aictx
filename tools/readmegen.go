package main

import (
	"bytes"
	"os"
	"os/exec"
	"text/template"

	"github.com/charmbracelet/log"
)

func main() {
	logger := log.NewWithOptions(os.Stderr, log.Options{
		Prefix: "readmegen 📚 ",
	})

	logger.Info("Generating readme.md from template...")
	defer func() {
		logger.Info("readme.md generated successfully.")
	}()

	// Read the template file.
	tmplData, err := os.ReadFile("tools/readme.template.md")
	if err != nil {
		logger.Fatalf("Error reading readme.template.md: %v", err)
	}

	// Run "aictx --help" and capture its output.
	cmd := exec.Command("./build/aictx", "--help")
	helpOutput, err := cmd.CombinedOutput()
	if err != nil {
		logger.Errorf("Failed running './aictx --help': %v", err)
	}

	// Prepare data for the template.
	data := struct {
		Usage string
	}{
		Usage: string(helpOutput),
	}

	// Parse the template.
	tmpl, err := template.New("readme").Parse(string(tmplData))
	if err != nil {
		logger.Fatalf("Failed parsing template: %v", err)
	}

	// Execute the template.
	var output bytes.Buffer
	if err := tmpl.Execute(&output, data); err != nil {
		logger.Fatalf("Failed executing template: %v", err)
	}

	if err := os.WriteFile("readme.md", output.Bytes(), 0600); err != nil {
		logger.Fatalf("Failed writing README.md: %v", err)
	}
}

package main

import (
	"encoding/json"
	"os"
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
)

// transformToJSON handles the Transform (T) command to parse JSON strings and open them in a new fx instance
func (m *model) transformToJSON() (tea.Model, tea.Cmd) {
	// Get the current cursor value
	value := m.cursorValue()
	if value == "" {
		return m, nil
	}

	// Try to parse the value as JSON
	var jsonData interface{}
	err := json.Unmarshal([]byte(value), &jsonData)
	if err != nil {
		// If it's not valid JSON, return without doing anything
		return m, nil
	}

	// Create a temporary file to store the JSON
	tmpFile, err := os.CreateTemp("", "fx-json-*.json")
	if err != nil {
		return m, nil
	}
	defer tmpFile.Close()

	// Write the JSON to the temporary file with proper formatting
	encoder := json.NewEncoder(tmpFile)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(jsonData); err != nil {
		os.Remove(tmpFile.Name())
		return m, nil
	}

	// Get the path to fx executable
	fxPath, err := os.Executable()
	if err != nil {
		// Fall back to "fx" in PATH
		fxPath = "fx"
	}

	// Open the JSON in a new fx instance
	execCmd := exec.Command(fxPath, tmpFile.Name())
	return m, tea.ExecProcess(execCmd, func(err error) tea.Msg {
		// Clean up the temporary file after fx exits
		os.Remove(tmpFile.Name())
		return nil
	})
}
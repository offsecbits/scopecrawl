package outputhandler

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// writeToFile handles writing to either .txt or .json
func writeToFile(filename string, content []string, format string) error {
	// Create the "output" directory if it doesn't exist
	if err := os.MkdirAll("output", os.ModePerm); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Create full file path
	filePath := filepath.Join("output", filename)

	var data []byte
	var err error

	if format == "json" {
		data, err = json.MarshalIndent(content, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal JSON: %w", err)
		}
	} else {
		// Join lines for TXT format
		data = []byte(strings.Join(content, "\n"))
	}

	// Write file
	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

// SaveLinks saves extracted links to file named after the base domain
func SaveLinks(baseURL string, links []string, format string) error {
        // Handle empty format by defaulting to txt
        if format != "json" {
                format = "txt"
        }


	safeName := sanitizeFilename(baseURL)
	extension := ".txt"
	if format == "json" {
		extension = ".json"
	}
	filename := safeName + extension
	return writeToFile(filename, links, format)
}

// sanitizeFilename ensures the filename is safe for writing
func sanitizeFilename(url string) string {
	url = strings.TrimPrefix(url, "https://")
	url = strings.TrimPrefix(url, "http://")
	url = strings.ReplaceAll(url, "/", "_")
	url = strings.ReplaceAll(url, ":", "_")
	return url
}

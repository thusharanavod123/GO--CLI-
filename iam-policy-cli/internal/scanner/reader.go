package scanner

import (
	"fmt"
	"os"
)

// FileData holds the path and the actual text content of a scanned file
type FileData struct {
	Path    string `json:"path"`
	Content string `json:"content"`
}

// ReadFiles opens each file path and reads its text content into memory
func ReadFiles(filePaths []string) ([]FileData, error) {
	var parsedFiles []FileData

	for _, path := range filePaths {
		// Read the entire file
		content, err := os.ReadFile(path)
		if err != nil {
			fmt.Printf("⚠️ Skipping file %s due to error: %v\n", path, err)
			continue
		}

		// Save the data into our container
		parsedFiles = append(parsedFiles, FileData{
			Path:    path,
			Content: string(content),
		})
	}

	return parsedFiles, nil
}
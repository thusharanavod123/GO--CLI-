package scanner

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// FindCodeFiles walks through a directory and returns a list of relevant file paths.
func FindCodeFiles(rootDir string) ([]string, error) {
	var files []string

	// Define the file extensions we care about for AWS/Infrastructure code
	validExtensions := map[string]bool{
		".py": true, // Python (Boto3)
		".tf": true, // Terraform
		".go": true, // Go (AWS SDK)
	}

	// filepath.WalkDir goes through every folder and file inside rootDir
	err := filepath.WalkDir(rootDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// 1. Skip hidden folders (like .git) and massive dependency folders (node_modules)
		if d.IsDir() {
			name := d.Name()
			if strings.HasPrefix(name, ".") || name == "node_modules" || name == "vendor" {
				return filepath.SkipDir
			}
			return nil
		}

		// 2. Check if the file ends with an extension we care about
		ext := filepath.Ext(path)
		if validExtensions[ext] {
			files = append(files, path)
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to scan directory: %w", err)
	}

	return files, nil
}
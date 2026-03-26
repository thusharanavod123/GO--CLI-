package cmd

import (
	"fmt"
	"path/filepath"

	"iam-policy-cli/internal/scanner" // <-- Importing our new package!

	"github.com/spf13/cobra"
)

var scanCmd = &cobra.Command{
	Use:   "scan [directory]",
	Short: "Scans a directory for code files",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		targetDir := "."
		if len(args) > 0 {
			targetDir = args[0]
		}

		absPath, err := filepath.Abs(targetDir)
		if err != nil {
			fmt.Println("Error reading directory path:", err)
			return
		}

		fmt.Println("🔍 Scanning directory:", absPath)

		// Call our new finding logic
		files, err := scanner.FindCodeFiles(absPath)
		if err != nil {
			fmt.Println("❌ Error scanning files:", err)
			return
		}

		// Print the results
		if len(files) == 0 {
			fmt.Println("⚠️  No relevant code files (.py, .tf, .go) found in this directory.")
			return
		}

		fmt.Printf("✅ Found %d files to analyze:\n", len(files))
		for _, file := range files {
			fmt.Println("  📄", file)
		}
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)
}
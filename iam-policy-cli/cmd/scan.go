package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
)

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:   "scan [directory]",
	Short: "Scans a directory for code files",
	Long: `The scan command looks through the specified directory (or the current directory if none is provided) 
to find infrastructure and application code.`,
	
	// This ensures the user provides either 0 or 1 argument (the folder path)
	Args: cobra.MaximumNArgs(1),
	
	Run: func(cmd *cobra.Command, args []string) {
		// Default to the current directory (".") if the user didn't type one
		targetDir := "."
		if len(args) > 0 {
			targetDir = args[0]
		}

		// Convert it to an absolute path so we know exactly where we are looking
		absPath, err := filepath.Abs(targetDir)
		if err != nil {
			fmt.Println("Error reading directory path:", err)
			return
		}

		fmt.Println("🔍 Initializing scan in directory:")
		fmt.Println("👉", absPath)
		fmt.Println("\n(Next up: We will add the logic to read the .tf and .py files in this folder!)")
	},
}

func init() {
	// This crucial line attaches the 'scan' command to your base 'root' command
	rootCmd.AddCommand(scanCmd)
}
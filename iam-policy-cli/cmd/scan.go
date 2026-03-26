package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
)


var scanCmd = &cobra.Command{
	Use:   "scan [directory]",
	Short: "Scans a directory for code files",
	Long: `The scan command looks through the specified directory (or the current directory if none is provided) 
to find infrastructure and application code.`,
	
	
	Args: cobra.MaximumNArgs(1),
	
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

		fmt.Println("🔍 Initializing scan in directory:")
		fmt.Println("👉", absPath)
		fmt.Println("\n(Next up: We will add the logic to read the .tf and .py files in this folder!)")
	},
}

func init() {
	
	rootCmd.AddCommand(scanCmd)
}
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "iam-policy-cli",
	Short: "A CLI tool to automatically generate AWS IAM policies from code",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🚀 Welcome to the IAM Policy Generator CLI!")
		fmt.Println("Use 'iam-policy-cli --help' to see available commands.")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
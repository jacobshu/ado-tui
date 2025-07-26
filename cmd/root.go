package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ado",
	Short: "A Terminal UI for interacting with Azure DevOps",
	Long:  "A TUI for managing Azure DevOps resources.",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize()
	rootCmd.AddCommand(workitemCmd)
}



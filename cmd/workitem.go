package cmd

import (
	"github.com/jacobshu/ado-tui/tui"
	"github.com/spf13/cobra"
)

var workitemCmd = &cobra.Command{
	Use:   "workitem",
	Short: "Manage Azure DevOps work items",
	Run: func(cmd *cobra.Command, args []string) {
		tui.RunWorkItem()
	},
}

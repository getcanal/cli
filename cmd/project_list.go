package cmd

import (
	"canal/util"
	"github.com/spf13/cobra"
)

// projectListCmd represents the project list command
var projectListCmd = &cobra.Command{
	Use:   "list",
	Short: "Shows projects you have access to",
	Run: func(cmd *cobra.Command, args []string) {
		util.PrintlnInfo("The projects you have access to:")
		util.PrintlnProjects()
	},
}

func init() {
	projectCmd.AddCommand(projectListCmd)
}

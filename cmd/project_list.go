package cmd

import (
	"cli/api/projects"
	"cli/util"
	"fmt"
	"github.com/spf13/cobra"
)

// projectListCmd represents the project list command
var projectListCmd = &cobra.Command{
	Use:   "list",
	Short: "Shows projects you have access to",
	Run: func(cmd *cobra.Command, args []string) {
		token, err := util.UserToken()
		if err != nil {
			util.PrintlnError(err)
			return
		}

		projects, err := api.ProjectList(token)
		if err != nil {
			util.PrintlnError(err)
			return
		}

		util.PrintlnInfo("The projects you have access to:")
		for i, project := range projects {
			fmt.Printf("%v. %v\n", i+1, project.Id)
		}

	},
}

func init() {
	projectCmd.AddCommand(projectListCmd)
}

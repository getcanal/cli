package cmd

import (
	"cli/api/auth"
	"cli/util"
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// projectUseCmd represents the project use <project> command
var projectUseCmd = &cobra.Command{
	Use:   "use",
	Short: "Use project of your choice",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		project := args[0]
		token, err := util.UserToken()
		if err != nil {
			util.PrintlnError(err)
			return
		}

		projectToken, err := auth.LoginProject(token, project)
		if err != nil {
			util.PrintlnError(err)
			return
		}

		err = util.StoreProjectToken(util.ProjectName(project), projectToken)
		if err != nil {
			util.PrintlnError(err)
			return
		}

		err = util.UseProject(util.ProjectName(project))
		if err != nil {
			util.PrintlnError(err)
			return
		}

		fmt.Printf("Switched %v to %v", color.CyanString("project"), color.CyanString(project))
	},
}

func init() {
	projectCmd.AddCommand(projectUseCmd)
}

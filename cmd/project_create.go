package cmd

import (
	api "canal/api/projects"
	"canal/util"
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// projectCreateCmd represents the project create command
var projectCreateCmd = &cobra.Command{
	Use:   "create <project>",
	Short: "Create a new project in Canal",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		project := util.ProjectName(args[0])
		token, err := util.UserToken()
		if err != nil {
			util.PrintlnError(err)
			return
		}

		util.PrintlnInfo("creating project... ")
		err = api.AddProject(api.Project{
			Id:          string(project),
			DisplayName: string(project),
		}, token)
		if err != nil {
			util.PrintlnError(err)
			return
		}

		fmt.Printf("Project %v has been successfully created\n", color.CyanString(string(project)))
		color.Cyan("Done!")
	},
}

func init() {
	projectCmd.AddCommand(projectCreateCmd)
}

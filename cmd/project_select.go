package cmd

import (
	"canal/util"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// projectSelectCmd represents the project select command
var projectSelectCmd = &cobra.Command{
	Use:   "select <project>",
	Short: "Select project of your choice",
	Run: func(cmd *cobra.Command, args []string) {
		token, err := util.UserToken()
		if err != nil {
			util.PrintlnError(err)
			return
		}

		err = util.SelectProject(token)
		if err != nil {
			util.PrintlnError(err)
			return
		}

		color.Cyan("Done!")
	},
}

func init() {
	projectCmd.AddCommand(projectSelectCmd)
}

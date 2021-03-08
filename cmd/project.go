package cmd

import (
	"canal/util"
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// projectCmd represents the project command
var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "Shows currently selected project",
	Run: func(cmd *cobra.Command, args []string) {
		project, err := util.CurrentProject()
		if err != nil {
			fmt.Printf("%v: project not selected\n", color.CyanString("Canal"))
			return
		}
		fmt.Printf("%v: Current project: %v\n", color.CyanString("Canal"), color.CyanString(string(project)))
	},
}

func init() {
	rootCmd.AddCommand(projectCmd)
}

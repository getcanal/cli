package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Displays version of installed Canal CLI",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(GetCurrentVersion())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func GetCurrentVersion() string {
	return "0.0.6" // ci-version-check
}

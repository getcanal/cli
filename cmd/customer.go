package cmd

import (
	"github.com/spf13/cobra"
)

// customerCmd represents the customer command
var customerCmd = &cobra.Command{
	Use: "customer",
}

func init() {
	rootCmd.AddCommand(customerCmd)
}

package cmd

import (
	"cli/api/auth"
	"cli/util"
	"fmt"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login <email>",
	Short: "Log in to Canal",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		passwordPrompt := promptui.Prompt{
			Label: "Password",
			Mask:  '*',
		}
		email := args[0]

		fmt.Printf("%v: Welcome back!\n", color.CyanString("Canal"))
		password, err := passwordPrompt.Run()
		if err != nil {
			util.PrintlnError(err)
			return
		}

		fmt.Print("Logging in... ")
		token, err := auth.Login(auth.Credentials{
			Email:    email,
			Password: password,
		})
		if err != nil {
			util.PrintlnError(err)
			return
		}

		err = util.StoreUserToken(util.Email(email), token)
		if err != nil {
			util.PrintlnError(err)
			return
		}

		err = util.ClearProjects()
		if err != nil {
			util.PrintlnError(err)
			return
		}

		color.Cyan("done")
		fmt.Printf("Logged %s as %s", color.CyanString("in"), color.CyanString("%v", email))
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}

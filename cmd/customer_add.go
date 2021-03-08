package cmd

import (
	customersApi "canal/api/customers"
	"canal/util"
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// customerAddCmd represents the customer add command
var customerAddCmd = &cobra.Command{
	Use:   "add email:<email> name:<name> phone:<phone>",
	Short: "Adds a new customer",
	Args: func(cmd *cobra.Command, args []string) error {
		_, err := util.EmailArg(args)
		return err
	},
	Run: func(cmd *cobra.Command, args []string) {
		project, err := util.CurrentProject()

		email, _ := util.EmailArg(args)
		name, _ := util.NameArg(args)
		phone, _ := util.PhoneArg(args)

		if err != nil {
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

			project, err = util.CurrentProject()
			if err != nil {
				util.PrintlnError(err)
				return
			}
		}

		util.PrintlnInfo(fmt.Sprintf("waiting %v Canal", color.CyanString("for")))
		fmt.Printf("Adding %v... ", email)

		token, err := util.ProjectToken(util.ProjectName(project))
		if err != nil {
			util.PrintlnError(err)
			return
		}

		err = customersApi.AddCustomer(token, customersApi.Customer{
			Email:    email,
			Name:     name,
			LastName: name,
			Phone:    phone,
		})
		if err != nil {
			util.PrintlnError(err)
			return
		}

		fmt.Printf(" %v!", color.CyanString("done!"))
	},
}

func init() {
	customerCmd.AddCommand(customerAddCmd)
}

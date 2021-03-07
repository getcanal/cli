package cmd

import (
	"canal/api/auth"
	api "canal/api/customers"
	"canal/util"
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search email:<email>",
	Short: "Queries Canal for the information you need",
	Long: fmt.Sprintf(
		"Queries Canal for the information you need\n" +
			"e.g. canal search email:<email> will return users matching given email query\n\n" +
			"supported query fields:\n" +
			"- email\n"),
	Args: func(cmd *cobra.Command, args []string) error {
		_, err := util.EmailArg(args)
		return err
	},
	Run: func(cmd *cobra.Command, args []string) {
		project, err := util.CurrentProject()
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
		email, _ := util.EmailArg(args)

		util.PrintlnInfo(fmt.Sprintf("waiting %v Canal", color.CyanString("for")))

		token, err := util.ProjectToken(project)
		if err != nil {
			userToken, err := util.UserToken()
			if err != nil {
				util.PrintlnError(err)
				return
			}
			projectToken, err := auth.LoginProject(userToken, string(project))
			err = util.StoreProjectToken(project, projectToken)
			if err != nil {
				util.PrintlnError(err)
				return
			}
			token = projectToken
		}

		customers, err := api.CustomerList(token)
		if err != nil {
			util.PrintlnError(err)
			return
		}

		filtered := filter(customers, func(customer api.Customer) bool {
			return strings.Contains(strings.ToLower(customer.Email), strings.ToLower(email))
		})

		fmt.Printf("Found %v... %v!\n", color.CyanString(strconv.Itoa(len(filtered))), color.CyanString("done"))
		fmt.Println(filtered)
	},
}

func filter(customers []api.Customer, test func(api.Customer) bool) (matching []api.Customer) {
	for _, s := range customers {
		if test(s) {
			matching = append(matching, s)
		}
	}
	return matching
}

func init() {
	rootCmd.AddCommand(searchCmd)
}

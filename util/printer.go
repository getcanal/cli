package util

import (
	api "cli/api/projects"
	"fmt"
	"github.com/fatih/color"
)

func PrintlnError(err error) {
	fmt.Printf("%s %v\n", color.RedString("Error:"), err)
}

func PrintlnInfo(info string) {
	fmt.Printf("%v: %v\n", color.CyanString("Canal"), info)
}

func PrintlnProjects() {
	token, err := UserToken()
	if err != nil {
		PrintlnError(err)
		return
	}

	projects, err := api.ProjectList(token)
	if err != nil {
		PrintlnError(err)
		return
	}

	for i, project := range projects {
		fmt.Printf("%v. %v\n", i+1, project.Id)
	}
}

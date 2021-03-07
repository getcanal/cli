package util

import (
	"canal/api/auth"
	projectsApi "canal/api/projects"
	"github.com/manifoldco/promptui"
)

func SelectProject(token auth.UserToken) error {
	projects, err := projectsApi.ProjectList(token)
	if err != nil {
		PrintlnError(err)
		return nil
	}
	var projectNames []string
	for i := range projects {
		projectNames = append(projectNames, projects[i].Id)
	}

	PrintlnInfo("please, select a project you have access to")
	prompt := promptui.Select{
		Items: projectNames,
	}
	_, selectedProject, err := prompt.Run()
	if err != nil {
		PrintlnError(err)
		return nil
	}

	projectToken, err := auth.LoginProject(token, selectedProject)
	if err != nil {
		PrintlnError(err)
		return nil
	}

	err = StoreProjectToken(ProjectName(selectedProject), projectToken)
	if err != nil {
		PrintlnError(err)
		return nil
	}

	err = UseProject(ProjectName(selectedProject))
	if err != nil {
		PrintlnError(err)
		return nil
	}

	return nil
}

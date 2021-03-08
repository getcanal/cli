package api

import (
	"canal/api/auth"
	"errors"
	"github.com/go-resty/resty/v2"
	"sort"
	"strings"
)

func ProjectList(user auth.UserToken) ([]Project, error) {
	client := resty.New()
	projectsSuccess := projectsResponseSuccess{}
	projectsError := projectsResponseError{}

	res, err := client.R().
		SetAuthToken(string(user)).
		SetResult(&projectsSuccess).
		SetError(&projectsError).
		Get("https://api.trycanal.com/v1/projects")

	if err != nil {
		return nil, err
	}

	if res.IsSuccess() {
		sort.Slice(projectsSuccess.Projects, func(i, j int) bool {
			compare := strings.Compare(strings.ToUpper(projectsSuccess.Projects[i].Id), strings.ToUpper(projectsSuccess.Projects[j].Id))
			return compare < 1
		})

		return projectsSuccess.Projects, nil
	}

	return nil, errors.New(projectsError.Message)
}

func AddProject(project Project, token auth.UserToken) error {
	client := resty.New()
	projectAddSuccess := projectAddResponseSuccess{}
	projectAddError := projectsResponseError{}

	res, err := client.R().
		SetAuthToken(string(token)).
		SetBody(projectAddRequestBody{
			Namespace:   project.Id,
			DisplayName: project.DisplayName,
		}).
		SetResult(&projectAddSuccess).
		SetError(&projectAddError).
		Post("https://api.trycanal.com/v1/projects")

	if err != nil {
		return err
	}

	if res.IsSuccess() {
		return nil
	}

	return errors.New(projectAddError.Message)
}

type projectsResponseError struct {
	Message string `json:"message"`
}

type projectsResponseSuccess struct {
	Projects []Project `json:"data"`
}

type projectAddRequestBody struct {
	Namespace   string `json:"namespace"`
	DisplayName string `json:"display_name"`
}

type projectAddResponseSuccess struct {
	Data data `json:"data"`
}

type data struct {
}

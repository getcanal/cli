package api

import (
	"cli/api/auth"
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

type projectsResponseError struct {
	Message string `json:"message"`
}

type projectsResponseSuccess struct {
	Projects []Project `json:"data"`
}

package auth

import (
	"errors"
	"github.com/go-resty/resty/v2"
)

func Login(credentials Credentials) (UserToken, error) {
	client := resty.New()
	loginSuccess := loginResponseSuccess{}
	loginError := loginResponseError{}

	res, err := client.R().
		SetBody(credentials).
		SetResult(&loginSuccess).
		SetError(&loginError).
		Post("https://api.trycanal.com/v1/auth/email/login")

	if err != nil {
		return UserToken(""), err
	}

	if res.IsSuccess() {
		return UserToken(loginSuccess.Data.Token), nil
	}

	return UserToken(""), errors.New(loginError.Message)
}

func LoginProject(token UserToken, project string) (ProjectToken, error) {
	client := resty.New()
	loginSuccess := loginResponseSuccess{}
	loginError := loginResponseError{}
	credentials := ProjectCredentials{Namespace: project}

	res, err := client.R().
		SetBody(credentials).
		SetAuthToken(string(token)).
		SetResult(&loginSuccess).
		SetError(&loginError).
		Post("https://api.trycanal.com/v1/auth/namespace")

	if err != nil {
		return ProjectToken(""), err
	}

	if res.IsSuccess() {
		return ProjectToken(loginSuccess.Data.Token), nil
	}

	return ProjectToken(""), errors.New(loginError.Message)
}

type loginResponseError struct {
	Message string `json:"message"`
}

type loginResponseSuccess struct {
	Data data `json:"data"`
}

type data struct {
	Token string `json:"authToken"`
}

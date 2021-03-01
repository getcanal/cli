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

type loginResponseError struct {
	Message string `json:"message"`
}

type loginResponseSuccess struct {
	Data data `json:"data"`
}

type data struct {
	Token UserToken `json:"authToken"`
}

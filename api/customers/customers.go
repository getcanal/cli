package api

import (
	"canal/api/auth"
	"errors"
	"github.com/go-resty/resty/v2"
)

func AddCustomer(project auth.ProjectToken, customer Customer) error {
	client := resty.New()
	customerAddSuccess := customerAddResponseSuccess{}
	customerAddError := customerAddResponseError{}

	res, err := client.R().
		SetAuthToken(string(project)).
		SetBody(customer).
		SetResult(&customerAddSuccess).
		SetError(&customerAddError).
		Post("https://api.trycanal.com/v1/customers")

	if err != nil {
		return err
	}

	if res.IsSuccess() {
		return nil
	}

	return errors.New(customerAddError.Message)
}

type customerAddResponseError struct {
	Message string `json:"message"`
}

type customerAddResponseSuccess struct {
	Data data `json:"data"`
}

type data struct {
}

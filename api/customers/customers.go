package api

import (
	"canal/api/auth"
	"errors"
	"github.com/go-resty/resty/v2"
	"sort"
	"strings"
)

func CustomerList(project auth.ProjectToken) ([]Customer, error) {
	client := resty.New()
	customersSuccess := customersResponseSuccess{}
	customersError := customersResponseError{}

	res, err := client.R().
		SetAuthToken(string(project)).
		SetResult(&customersSuccess).
		SetError(&customersError).
		Get("https://api.trycanal.com/v1/customers")

	if err != nil {
		return nil, err
	}

	if res.IsSuccess() {
		sort.Slice(customersSuccess.Customers, func(i, j int) bool {
			compare := strings.Compare(strings.ToUpper(customersSuccess.Customers[i].Email), strings.ToUpper(customersSuccess.Customers[j].Email))
			return compare < 1
		})

		return customersSuccess.Customers, nil
	}

	return nil, errors.New(customersError.Message)
}

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

type customersResponseError struct {
	Message string `json:"message"`
}

type customersResponseSuccess struct {
	Customers []Customer `json:"data"`
}

type customerAddResponseError struct {
	Message string `json:"message"`
}

type customerAddResponseSuccess struct {
	Data data `json:"data"`
}

type data struct{}

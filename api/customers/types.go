package api

import (
	"canal/util"
	"encoding/json"
	"os"
)

// TODO update customer model in the backend, CLI and UI
type Customer struct {
	Email    string `json:"email"`
	Name     string `json:"first_name"`
	LastName string `json:"last_name"`
	Phone    string `json:"phone"`
}

func (c Customer) String() string {
	bytes, err := json.MarshalIndent(c, "", "    ")
	if err != nil {
		util.PrintlnError(err)
		os.Exit(1)
	}
	return string(bytes)
}

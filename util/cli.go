package util

import (
	"errors"
	"strings"
)

func EmailArg(args []string) (string, error) {
	return argValue(args, "email")
}

func NameArg(args []string) (string, error) {
	return argValue(args, "name")
}

func PhoneArg(args []string) (string, error) {
	return argValue(args, "phone")
}

func argValue(args []string, argName string) (string, error) {
	for i := range args {
		prefix := argName + ":"
		if strings.HasPrefix(args[i], prefix) {
			return strings.TrimPrefix(args[i], prefix), nil
		}
	}
	return "", errors.New(argName + " not provided")
}

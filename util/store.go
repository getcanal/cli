package util

import (
	"cli/api/auth"
	"io/ioutil"
	"os"
)

func StoreProjectToken(token auth.ProjectToken, project string) error {
	err := assertCanalDirExists()
	if err != nil {
		return err
	}

	path, err := projectTokenFilePath(project)
	if err != nil {
		return err
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}

	_, err = file.WriteString(string(token))
	if err != nil {
		return err
	}

	return nil
}

func GetProjectToken(project string) (auth.ProjectToken, error) {
	path, err := projectTokenFilePath(project)
	if err != nil {
		return auth.ProjectToken(""), err
	}

	token, err := ioutil.ReadFile(path)
	if err != nil {
		return auth.ProjectToken(""), err
	}

	return auth.ProjectToken(token), nil
}

func StoreUserToken(token auth.UserToken) error {
	err := assertCanalDirExists()
	if err != nil {
		return err
	}

	path, err := userTokenFilePath()
	if err != nil {
		return err
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}

	_, err = file.WriteString(string(token))
	if err != nil {
		return err
	}

	return nil
}

func GetUserToken() (auth.UserToken, error) {
	path, err := userTokenFilePath()
	if err != nil {
		return auth.UserToken(""), err
	}

	token, err := ioutil.ReadFile(path)
	if err != nil {
		return auth.UserToken(""), err
	}

	return auth.UserToken(token), nil
}

func userTokenFilePath() (string, error) {
	canal, err := canalDirPath()
	if err != nil {
		return "", err
	}

	return canal + string(os.PathSeparator) + "user_token", nil
}

func projectTokenFilePath(project string) (string, error) {
	canal, err := canalDirPath()
	if err != nil {
		return "", err
	}

	return canal + string(os.PathSeparator) + project + "_project_token", nil
}

func canalDirPath() (string, error) {
	dir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	path := dir + string(os.PathSeparator) + ".canal"

	return path, nil
}

func assertCanalDirExists() error {
	path, err := canalDirPath()
	if err != nil {
		return err
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

package util

import (
	"cli/api/auth"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

const ConfigFileName = "config"

type CanalConfig struct {
	User     Email                             `json:"email"`
	Project  ProjectName                       `json:"project"`
	Token    auth.UserToken                    `json:"token"`
	Projects map[ProjectName]auth.ProjectToken `json:"projects"`
}
type Email string
type ProjectName string

func Config() (CanalConfig, error) {
	config := CanalConfig{}

	path, err := CanalConfigPath()
	if err != nil {
		return config, err
	}

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return config, err
	}

	err = json.Unmarshal(bytes, &config)
	if err != nil {
		return config, err
	}

	if config.Projects == nil {
		config.Projects = map[ProjectName]auth.ProjectToken{}
	}

	return config, err
}

func StoreConfig(config CanalConfig) error {
	bytes, err := json.Marshal(config)
	if err != nil {
		return err
	}

	path, err := CanalConfigPath()
	if err != nil {
		return err
	}

	return ioutil.WriteFile(path, bytes, os.ModePerm)
}

func ProjectToken(project ProjectName) (auth.ProjectToken, error) {
	config, err := Config()
	if err != nil {
		return auth.ProjectToken(""), err
	}

	token := config.Projects[project]
	if token == "" {
		return auth.ProjectToken(""), errors.New("project authentication token not found in configuration")
	}

	return auth.ProjectToken(token), nil
}

func StoreProjectToken(project ProjectName, token auth.ProjectToken) error {
	config, err := Config()
	if err != nil {
		return err
	}

	if config.Projects == nil {
		projects := map[ProjectName]auth.ProjectToken{}
		projects[project] = token
		config.Projects = projects
	} else {
		config.Projects[project] = token
	}

	return StoreConfig(config)
}

func CurrentProject() (ProjectName, error) {
	config, err := Config()
	if err != nil {
		return "", err
	}

	project := config.Project
	if project == "" {
		return "", errors.New("project not selected")
	}

	return project, nil
}

func UseProject(name ProjectName) error {
	config, err := Config()
	if err != nil {
		return err
	}

	config.Project = name

	return StoreConfig(config)
}

func ClearProjects() error {
	config, err := Config()
	if err != nil {
		return err
	}

	config.Projects = map[ProjectName]auth.ProjectToken{}
	config.Project = ""

	return StoreConfig(config)
}

func UserToken() (auth.UserToken, error) {
	config, err := Config()
	if err != nil {
		return auth.UserToken(""), err
	}

	token := config.Token
	if token == "" {
		return auth.UserToken(""), errors.New("user authentication token not found in configuration")
	}

	return auth.UserToken(token), nil
}

func StoreUserToken(user Email, token auth.UserToken) error {
	config, err := Config()
	if err != nil {
		return err
	}

	config.User = user
	config.Token = token

	return StoreConfig(config)
}

func CurrentUser() (string, error) {
	config, err := Config()
	if err != nil {
		return "", err
	}

	token := config.Token
	if token == "" {
		return "", errors.New("current user email not found in configuration")
	}

	return "", nil
}

func InitializeCanalConfig() error {
	if !CanalDirExists() {
		path, err := CanalDirPath()
		if err != nil {
			return err
		}

		err = os.Mkdir(path, os.ModePerm)
		if err != nil {
			return err
		}
	}

	path, err := CanalConfigPath()
	if err != nil {
		return err
	}

	_, err = os.Create(path)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path, []byte("{}"), os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func CanalConfigPath() (string, error) {
	path, err := CanalDirPath()
	if err != nil {
		return "", err
	}
	return path + string(os.PathSeparator) + ConfigFileName + ".json", nil
}

func CanalDirPath() (string, error) {
	dir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return dir + string(os.PathSeparator) + ".canal", nil
}

func CanalDirExists() bool {
	path, err := CanalDirPath()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return pathExists(path)
}

func CanalConfigExists() bool {
	path, err := CanalConfigPath()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return pathExists(path)
}

func pathExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	} else {
		return false
	}
}

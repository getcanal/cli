package util

import (
	"cli/api/auth"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"os"
)

const ConfigFileName = "config"

type CanalConfig struct {
	User     Email                             `yaml:"user"`
	Project  ProjectName                       `yaml:"project"`
	Token    auth.UserToken                    `yaml:"token"`
	Projects map[ProjectName]auth.ProjectToken `json:"projects"`
}
type Email string
type ProjectName string

func LoadConfig() (*CanalConfig, error) {
	config := CanalConfig{}
	err := viper.Unmarshal(&config)
	return &config, err
}

func StoreProjectToken(project ProjectName, token auth.ProjectToken) error {
	viper.Set("projects."+string(project), token)
	err := viper.WriteConfig()
	if err != nil {
		return err
	}
	return nil
}

func GetProjectToken(project ProjectName) (auth.ProjectToken, error) {
	token := viper.GetString("projects." + string(project))
	if token == "" {
		return auth.ProjectToken(""), errors.New("project authentication token not found in configuration")
	}
	return auth.ProjectToken(token), nil
}

func CurrentProject() (string, error) {
	project := viper.GetString("project")
	if project == "" {
		return "", errors.New("project not selected")
	}
	return project, nil
}

type Empty struct {
}

func ClearProjects() error {
	viper.Set("project", "")
	viper.Set("projects", map[ProjectName]auth.ProjectToken{})
	return viper.WriteConfig()
}

func StoreUserToken(user Email, token auth.UserToken) error {
	viper.Set("user", user)
	viper.Set("token", token)
	err := viper.WriteConfig()
	if err != nil {
		return err
	}
	return nil
}

func GetUserToken() (auth.UserToken, error) {
	token := viper.GetString("token")
	if token == "" {
		return auth.UserToken(""), errors.New("user authentication token not found in configuration")
	}
	return auth.UserToken(token), nil
}

func GetCurrentUser() (string, error) {
	token := viper.GetString("user")
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

	return nil
}

func CanalConfigPath() (string, error) {
	path, err := CanalDirPath()
	if err != nil {
		return "", err
	}
	return path + string(os.PathSeparator) + ConfigFileName + ".yml", nil
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

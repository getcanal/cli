package util

import (
	"cli/api/auth"
	"fmt"
	"github.com/spf13/viper"
	"os"
)

const ConfigFileName = "config"

func StoreProjectToken(token auth.ProjectToken, project string) error {
	viper.Set(project+"_project_token", token)
	err := viper.WriteConfig()
	if err != nil {
		return err
	}

	return nil
}

func GetProjectToken(project string) (auth.ProjectToken, error) {
	token := viper.GetString(project + "_project_token")
	return auth.ProjectToken(token), nil
}

func StoreUserToken(token auth.UserToken) error {
	viper.Set("user_token", token)
	err := viper.WriteConfig()
	if err != nil {
		return err
	}

	return nil
}

func GetUserToken() (auth.UserToken, error) {
	token := viper.GetString("user_token")
	return auth.UserToken(token), nil
}

func InitializeCanalConfig() error {
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

	path := dir + string(os.PathSeparator) + ".canal"

	return path, nil
}

func CanalConfigExists() bool {
	path, err := CanalConfigPath()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if _, err := os.Stat(path); err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	} else {
		return false
	}
}

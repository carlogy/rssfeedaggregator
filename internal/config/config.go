package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (c Config) SetUser(username string) {

	if len(username) == 0 {
		fmt.Println("Error: Username can't be empty ")
	}

	c.CurrentUserName = username

	err := write(c)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Updated file")
}

func write(cfg Config) error {

	data, err := json.Marshal(cfg)
	if err != nil {

		return err
	}

	configPath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	err = os.WriteFile(configPath, data, os.ModeAppend)
	if err != nil {
		return fmt.Errorf("Error writing config file:\n%w", err)
	}

	return nil
}

package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const configFileName = ".aggregatorConfig.json"

// bootdev ver:
// const configFileName = ".gatorconfig.json"

func Read() (Config, error) {

	filePath, err := getConfigFilePath()

	if err != nil {
		return Config{}, err
	}

	f, err := os.ReadFile(filePath)
	if err != nil {
		return Config{}, err
	}

	config := Config{}

	if err := json.Unmarshal(f, &config); err != nil {
		return Config{}, fmt.Errorf("Error decoding config file:\n %w", err)
	}

	return config, nil

}

func getConfigFilePath() (string, error) {

	homeDir, err := os.UserHomeDir()

	if err != nil {
		return "", fmt.Errorf("Encountered the following error while attempting to get user home directory path:\n%w", err)
	}

	configPath := fmt.Sprintf("%s/%s", homeDir, configFileName)

	fileExists := checkFileExists(configPath)

	if !fileExists {
		return "", fmt.Errorf("%s does not exist", configPath)
	}

	return configPath, nil

}

func checkFileExists(filepath string) bool {
	_, err := os.Stat(filepath)

	if err != nil {
		return false
	}

	return true
}

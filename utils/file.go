package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// File writing and Reading utils

func GetConfigPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to find the Home Directory: %v", err)
	}

	configPath := filepath.Join(homeDir, FLENV_CONFIG_FILENAME)
	return configPath, nil
}

func ReadConfigFile() (Flenv, error) {
	configPath, err := GetConfigPath()
	if err != nil {
		return Flenv{}, fmt.Errorf("failure retrieving the config path: %v", err)
	}
	jsonData, err := os.ReadFile(configPath)
	if err != nil {
		return Flenv{}, fmt.Errorf("failed to read config file %v", err)
	}

	var config Flenv
	err = json.Unmarshal(jsonData, &config)
	if err != nil {
		return Flenv{}, fmt.Errorf("can't unmarshal file: %v", err)
	}
	return config, nil
}

func WriteNewConfigFile(config Flenv) error {
	configPath, err := GetConfigPath()
	if err != nil {
		return fmt.Errorf("failure getting config path: %v", err)
	}

	jsonData, err := json.MarshalIndent(config, "", " ")
	if err != nil {
		return fmt.Errorf("JSON Marshal Error: %v", err)
	}

	// Define permissions using constants from the os package
	permissions := os.FileMode(0o644)
	err = os.WriteFile(configPath, jsonData, permissions)
	if err != nil {
		return fmt.Errorf("failed writing the config file: %v", err)
	}
	return nil
}

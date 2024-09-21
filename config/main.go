package config

import (
	"fmt"
	"log"
	"os"
)

func ValidateAndProcess(newConfigName string, list bool) {
	if newConfigName == "" && list == false {
		fmt.Println("Error: Either the create or list flags must be provided")
		os.Exit(1)
	}

	if newConfigName != "" && list == true {
		fmt.Println("Only one of create or list flag should be provided, and not a combination")
		os.Exit(1)
	}

	ProcessConfig(newConfigName, list)
}

func ProcessConfig(newConfigName string, list bool) {
	if newConfigName != "" {
		err := CreateNewConfig(newConfigName)
		if err != nil {
			log.Fatalf("Failed to create a new stage: %v", err)
		}
	}

	if list {
		err := ListConfigs()
		if err != nil {
			log.Fatalf("failed to list current configs: %v", err)
		}
	}
}

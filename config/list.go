package config

import (
	"fmt"

	"github.com/yvanflorian/flenv/utils"
)

func ListConfigs() error {
	currentConfig, err := utils.ReadConfigFile()
	if err != nil {
		return fmt.Errorf("failure reading config file: %v", err)
	}

	fmt.Println("Available config(s) across stage(s): ")
	oneStage := currentConfig.Stages[0]
	for _, config := range oneStage.Configs {
		fmt.Printf(" > %v\n", config.Name)
	}
	return nil
}

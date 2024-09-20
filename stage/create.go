package stage

import (
	"fmt"
	"log"
	"os"

	"github.com/yvanflorian/flenv/utils"
)

// Handle Creation of stage
// even when there's no file
//

func CreateNewStage(name string) error {
	log.Println("Creating a new Stage:", name)

	configPath, err := utils.GetConfigPath()
	if err != nil {
		return fmt.Errorf("Error getting config file: %v", err)
	}

	_, err = os.Stat(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			log.Println("File not existing creating it")
			newConfig := utils.NewConfig(name)
			return utils.WriteNewConfigFile(newConfig)
		} else {
			return fmt.Errorf("cannot find the file. Error: %v", err)
		}
	}

	existingConfig, err := utils.ReadConfigFile()
	if err != nil {
		return fmt.Errorf("failure reading config file: %v", err)
	}
	return AppendNewStage(name, existingConfig)
}

func AppendNewStage(stageName string, config utils.Flenv) error {
	log.Println("appending new stage to existing config", stageName)
	var stages []utils.Stage

	for _, stage := range config.Stages {
		if stage.StageName == stageName {
			return fmt.Errorf("Stage already exists in your config!")
		}
		stages = append(stages, stage)
	}
	stages = append(stages, utils.Stage{StageName: stageName})
	newConfig := utils.Flenv{
		Stages: stages,
	}
	return utils.WriteNewConfigFile(newConfig)
}

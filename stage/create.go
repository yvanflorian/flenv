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
	//first stage
	if len(config.Stages) == 0 {
		newConfig := utils.Flenv{
			Stages: stages,
		}
		return utils.WriteNewConfigFile(newConfig)
	}

	existingStage := config.Stages[0]
	var newConfigs []utils.Config
	for _, conf := range existingStage.Configs {
		var newVars []utils.Variable
		for _, val := range conf.Variables {
			newVar := utils.Variable{
				Key:   val.Key,
				Value: "",
			}
			newVars = append(newVars, newVar)
		}
		newConfigs = append(newConfigs, utils.Config{
			Name:      conf.Name,
			Variables: newVars,
		})
	}
	newStage := utils.Stage{
		StageName: stageName,
		Configs:   newConfigs,
	}

	for _, stage := range config.Stages {
		if stage.StageName == stageName {
			return fmt.Errorf("Stage already exists in your config!")
		}
		stages = append(stages, stage)
	}
	stages = append(stages, newStage)
	newConfig := utils.Flenv{
		Stages: stages,
	}
	return utils.WriteNewConfigFile(newConfig)
}

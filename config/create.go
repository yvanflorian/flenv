package config

import (
	"bufio"
	"fmt"
	"os"

	"github.com/yvanflorian/flenv/utils"
)

func CreateNewConfig(name string) error {
	configPath, err := utils.GetConfigPath()
	if err != nil {
		return fmt.Errorf("Error getting config file: %v", err)
	}

	_, err = os.Stat(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("cannot find the config file. Use the 'stage' subcommand to create a new config file")
		} else {
			return fmt.Errorf("cannot find the file. Error: %v", err)
		}
	}

	existingConfig, err := utils.ReadConfigFile()
	if err != nil {
		return fmt.Errorf("failure reading config file: %v", err)
	}
	err = checkConflict(name, existingConfig)
	if err != nil {
		return fmt.Errorf("Conflict  Detected!", err)
	}

	newConf, err := loopNewConfig(name, existingConfig)
	if err != nil {
		return fmt.Errorf("error writing new config: %v", err)
	}

	return utils.WriteNewConfigFile(newConf)
}

func loopNewConfig(name string, existingConfig utils.Flenv) (utils.Flenv, error) {
	scanner := bufio.NewScanner(os.Stdin)
	var stages []utils.Stage

	fmt.Print("> Enter the variable name: ")
	scanner.Scan()
	key := scanner.Text()
	// Prompt for name
	for _, stage := range existingConfig.Stages {
		fmt.Printf("> Enter the variable value for '%v' and the stage '%v': ", key, stage.StageName)
		scanner.Scan()
		val := scanner.Text()

		fmt.Printf("Config Name is: %s\n", key)
		fmt.Printf("Config Value is: %s\n", val)
		newConfig := utils.Config{
			Name:      name,
			Variables: append([]utils.Variable{}, utils.Variable{Key: key, Value: val}),
		}

		stages = append(stages, utils.Stage{
			StageName: stage.StageName,
			Configs:   append(stage.Configs, newConfig),
		})
	}
	return utils.Flenv{Stages: stages}, nil
}

func checkConflict(name string, existingConfig utils.Flenv) error {
	if len(existingConfig.Stages) == 0 {
		return fmt.Errorf("Current Stage doesn't have a stage set. Please review!")
	}

	oneStage := existingConfig.Stages[0]
	for _, config := range oneStage.Configs {
		if config.Name == name {
			return fmt.Errorf("There's an existing Config with the same name. Verify with the List flag")
		}
	}
	return nil
}

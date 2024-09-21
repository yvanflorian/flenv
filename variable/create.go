package variable

import (
	"bufio"
	"fmt"
	"os"

	"github.com/yvanflorian/flenv/utils"
)

func CreateNewVariable(name string, config string) error {
	// fmt.Printf("creating %v for config %v", name, config)
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

	err = checkConflict(name, config, existingConfig)
	if err != nil {
		return fmt.Errorf("Conflict  Detected!", err)
	}

	newConf, err := loopNewVariable(name, config, existingConfig)
	if err != nil {
		return fmt.Errorf("error writing new config: %v", err)
	}

	return utils.WriteNewConfigFile(newConf)
}

func loopNewVariable(newVarName string, configName string, existingConfig utils.Flenv) (utils.Flenv, error) {
	scanner := bufio.NewScanner(os.Stdin)
	var stages []utils.Stage

	// Prompt for name
	for _, iStage := range existingConfig.Stages {
		fmt.Printf("> Enter the variable value for '%v' and the stage '%v': ", newVarName, iStage.StageName)
		scanner.Scan()
		val := scanner.Text()

		fmt.Printf("Variable name is: %s\n", newVarName)
		fmt.Printf("Variable value is: %s\n", val)
		var newConfigs []utils.Config
		for _, iConfig := range iStage.Configs {
			varList := append([]utils.Variable{}, iConfig.Variables...)
			if iConfig.Name == configName {
				varList = append(varList, utils.Variable{
					Key:   newVarName,
					Value: val,
				})
			}
			newConfigs = append(newConfigs, utils.Config{
				Name:      iConfig.Name,
				Variables: varList,
			})
		}

		stages = append(stages, utils.Stage{
			StageName: iStage.StageName,
			Configs:   newConfigs,
		})
	}
	return utils.Flenv{Stages: stages}, nil
}

func checkConflict(varName string, configName string, existingConfig utils.Flenv) error {
	if len(existingConfig.Stages) == 0 {
		return fmt.Errorf("Current Stage doesn't have a stage set. Please review!")
	}

	oneStage := existingConfig.Stages[0]
	for _, config := range oneStage.Configs {
		// loop vars of this config
		if configName == config.Name {
			for _, val := range config.Variables {
				if val.Key == varName {
					return fmt.Errorf("There's an existing Config with the same name. Verify with the List flag")
				}
			}
		}
	}
	return nil
}

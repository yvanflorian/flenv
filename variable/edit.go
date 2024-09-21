package variable

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/yvanflorian/flenv/utils"
)

func EditVariable(pConfig string, pVarName string, pStage string) error {
	log.Println("editing variable...")
	existingConfig, err := utils.ReadConfigFile()
	if err != nil {
		return fmt.Errorf("failure reading existing config: %v", err)
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("> Enter the variable value for variable '%v': ", pVarName)
	scanner.Scan()
	val := scanner.Text()

	var newStages []utils.Stage
	for _, iStage := range existingConfig.Stages {
		if iStage.StageName == pStage {
			var newConfigs []utils.Config
			for _, iConfig := range iStage.Configs {
				if iConfig.Name == pConfig {
					// search for variable to edit
					var newVars []utils.Variable
					for _, iVar := range iConfig.Variables {
						if iVar.Key == pVarName {
							newVars = append(newVars, utils.Variable{Key: iVar.Key, Value: val})
						} else {
							newVars = append(newVars, iVar)
						}
					}
					newConfigs = append(newConfigs, utils.Config{
						Name:      iConfig.Name,
						Variables: newVars,
					})
				} else {
					newConfigs = append(newConfigs, iConfig)
				}
			}
			newStages = append(newStages, utils.Stage{
				StageName: iStage.StageName,
				Configs:   newConfigs,
			})
		} else {
			newStages = append(newStages, iStage)
		}
	}
	newConf := utils.Flenv{Stages: newStages}
	return utils.WriteNewConfigFile(newConf)
}

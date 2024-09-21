package variable

import (
	"fmt"
	"os"

	"github.com/yvanflorian/flenv/utils"
)

func ShowVariable(pConfig string, pVarName string) error {
	config, err := utils.ReadConfigFile()
	if err != nil {
		return fmt.Errorf("failed to read config file...%v", err)
	}

	countFound := 0
	for _, stage := range config.Stages {
		for _, config := range stage.Configs {
			if config.Name == pConfig {
				for _, cVar := range config.Variables {
					if cVar.Key == pVarName {
						fmt.Printf(" > Stage: %v Variable value: %v\n", stage.StageName, cVar.Value)
						countFound++
					}
				}
			}
		}
	}

	if countFound == 0 {
		fmt.Printf("Error: Variable '%v' was not found under config '%v'\n", pVarName, pConfig)
		os.Exit(1)
	}

	return nil
}

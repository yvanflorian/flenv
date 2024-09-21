package variable

import (
	"fmt"

	"github.com/yvanflorian/flenv/utils"
)

func ShowVariable(pConfig string, pVarName string) error {
	config, err := utils.ReadConfigFile()
	if err != nil {
		return fmt.Errorf("failed to read config file...%v", err)
	}

	for _, stage := range config.Stages {
		for _, config := range stage.Configs {
			if config.Name == pConfig {
				for _, cVar := range config.Variables {
					if cVar.Key == pVarName {
						fmt.Printf(" > Stage: %v Variable value: %v\n", stage.StageName, cVar.Value)
					}
				}
			}
		}
	}

	return nil
}

package variable

import (
	"fmt"

	"github.com/yvanflorian/flenv/utils"
)

func ListVariable(pConfig string, pStage string) error {
	config, err := utils.ReadConfigFile()
	if err != nil {
		return fmt.Errorf("failed to read config file...%v", err)
	}

	for _, stage := range config.Stages {
		if pStage == stage.StageName {
			//when stage specified
			for _, config := range stage.Configs {
				if config.Name == pConfig {
					for _, cVar := range config.Variables {
						fmt.Printf(" > Variable: %v\n", cVar.Key)
					}
				}
			}
		}
	}

	return nil
}

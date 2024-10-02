package stage

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/yvanflorian/flenv/utils"
)

func SetStage(stageName string) error {
	config, err := utils.ReadConfigFile()
	if err != nil {
		return err
	}
	for _, stage := range config.Stages {
		if stage.StageName == stageName {
			for _, config := range stage.Configs {
				confPrefix := strings.ToUpper(config.Name)
				for _, val := range config.Variables {
					varPrefix := strings.ToUpper(val.Key)
					envName := fmt.Sprintf("%v_%v_%v", utils.FLENV_CONFIG_PREFIX, confPrefix, varPrefix)
					log.Println("Env name:", envName)
					err := os.Setenv(envName, string(val.Value))
					if err != nil {
						return fmt.Errorf("failed to set env: %v", err)
					}
				}
			}
		}
	}
	fmt.Println("Environments set successfully")
	return nil
}

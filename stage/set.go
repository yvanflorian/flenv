package stage

import (
	"fmt"

	"github.com/yvanflorian/flenv/utils"
)

func SetStage(pStageName string) error {
	config, err := utils.ReadConfigFile()
	if err != nil {
		return fmt.Errorf("failed to read config file...%v", err)
	}

	var stages []string
	countExists := 0
	for _, stage := range config.Stages {
		stages = append(stages, stage.StageName)
		if stage.StageName == pStageName {
			countExists++
		}
	}
	if countExists == 0 {
		return fmt.Errorf("Stage: %v not part of your configuration", pStageName)
	}

	fmt.Printf("export FLENV_STAGE_ENVIRONMENT=%v\n", pStageName)
	return nil
}

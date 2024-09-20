package stage

import (
	"fmt"
	"log"

	"github.com/yvanflorian/flenv/utils"
)

func ListStages() ([]string, error) {
	log.Println("List down stages")
	config, err := utils.ReadConfigFile()
	if err != nil {
		return nil, fmt.Errorf("failed to read config file...%v", err)
	}

	var stages []string
	for _, stage := range config.Stages {
		stages = append(stages, stage.StageName)
	}

	return stages, nil
}

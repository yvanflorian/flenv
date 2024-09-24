package config

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/yvanflorian/flenv/utils"
)

func Handle(args []string) {
	configCommands := flag.NewFlagSet("config", flag.ExitOnError)
	configname := configCommands.String("create", "", "config Name to Create")
	configList := configCommands.Bool("list", false, "List Available configs in all stages")

	configCommands.Usage = utils.ShowDocsConfig

	configCommands.Parse(args)
	// Check if any flags were provided
	utils.NoEmptyFlags(configCommands)
	ValidateAndProcess(*configname, *configList)
}

func ValidateAndProcess(newConfigName string, list bool) {
	if newConfigName == "" && list == false {
		fmt.Println("Error: Either the create or list flags must be provided")
		os.Exit(1)
	}

	if newConfigName != "" && list == true {
		fmt.Println("Only one of create or list flag should be provided, and not a combination")
		os.Exit(1)
	}

	ProcessConfig(newConfigName, list)
}

func ProcessConfig(newConfigName string, list bool) {
	if newConfigName != "" {
		err := CreateNewConfig(newConfigName)
		if err != nil {
			log.Fatalf("Failed to create a new stage: %v", err)
		}
	}

	if list {
		err := ListConfigs()
		if err != nil {
			log.Fatalf("failed to list current configs: %v", err)
		}
	}
}

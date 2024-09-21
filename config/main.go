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

	configCommands.Usage = func() {
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "USAGE\n")
		fmt.Fprintf(os.Stderr, "  flenv config [flags]\n")
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "AVAILABLE FLAGS\n")
		fmt.Fprintf(os.Stderr, "  create: Create a new config. Multiple Prompts will follow to create the same config for all stages\n")
		fmt.Fprintf(os.Stderr, "  list: View & Manage specific configuration variables\n")
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "EXAMPLE\n")
		fmt.Fprintf(os.Stderr, " #Create a new config called database\n")
		fmt.Fprintf(os.Stderr, " $flenv config --create database\n")
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, " #List down all config that flenv is managing\n")
		fmt.Fprintf(os.Stderr, " $flenv config --list\n")
		fmt.Fprintf(os.Stderr, "\n")
		// configCommands.PrintDefaults()
	}
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

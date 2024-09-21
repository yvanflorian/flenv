package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/yvanflorian/flenv/config"
	"github.com/yvanflorian/flenv/stage"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Expected either 'stage' , 'config' or 'variable'")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "stage":
		stageCommands := flag.NewFlagSet("stage", flag.ExitOnError)
		stagename := stageCommands.String("create", "", "Stage Name to Create")
		stageList := stageCommands.Bool("list", false, "List the available stages")
		stageSet := stageCommands.String("set", "", "Set all environment variables for the given stagename")
		stageCommands.Parse(os.Args[2:])
		noEmptyFlags(stageCommands)
		stage.ValidateAndProcessStage(*stagename, *stageList, *stageSet)
	case "config":
		configCommands := flag.NewFlagSet("config", flag.ExitOnError)
		configname := configCommands.String("create", "", "config Name to Create")
		configList := configCommands.Bool("list", false, "List Available configs in all stages")
		configCommands.Parse(os.Args[2:])
		// Check if any flags were provided
		noEmptyFlags(configCommands)
		config.ValidateAndProcess(*configname, *configList)
	default:
		fmt.Println("Wrong command: Either 'stage', 'config' or 'variable'")
		os.Exit(1)
	}
}

func noEmptyFlags(cmd *flag.FlagSet) {
	if cmd.NFlag() == 0 {
		fmt.Printf("Error: The '%v' subcommand requires flags. Please review!", cmd.Name())
		cmd.PrintDefaults()
		os.Exit(1)
	}
}

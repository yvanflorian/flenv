package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/yvanflorian/flenv/stage"
)

func config(name string) {
	fmt.Println("config name to be created...", name)
}

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
		// Check if any flags were provided
		if stageCommands.NFlag() == 0 {
			fmt.Println("Error: The 'stage' subcommand requires flags")
			stageCommands.PrintDefaults()
			os.Exit(1)
		}
		stage.ValidateAndProcessStage(*stagename, *stageList, *stageSet)
	case "config":
		configCommands := flag.NewFlagSet("config", flag.ExitOnError)
		configname := configCommands.String("create", "", "config Name to Create")
		configCommands.Parse(os.Args[2:])
		config(*configname)
	default:
		fmt.Println("Wrong command: Either 'stage', 'config' or 'variable'")
		os.Exit(1)
	}
}

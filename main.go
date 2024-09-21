package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/yvanflorian/flenv/config"
	"github.com/yvanflorian/flenv/stage"
	"github.com/yvanflorian/flenv/variable"
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
	case "variable":
		varCommands := flag.NewFlagSet("variable", flag.ExitOnError)
		varConfig := varCommands.String("config", "", "Config Name that holds this variable")
		varCreate := varCommands.String("create", "", "Variable name to Create")
		varShow := varCommands.String("show", "", "Display the Variable value in given stages")
		varEdit := varCommands.String("edit", "", "Variable name to edit")
		varStage := varCommands.String("stage", "", "Stage that owns the config and the variable")
		varCommands.Parse(os.Args[2:])
		noEmptyFlags(varCommands)
		variable.ValidateAndProcess(*varCreate, *varShow, *varEdit, *varConfig, *varStage)
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

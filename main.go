package main

import (
	"flag"
	"fmt"
	"os"
)

func validateAndProcessStage(name string, list bool, set string) {
	if name == "" && list == false && set == "" {
		fmt.Println("Either the create, list or set flags must be provided")
		os.Exit(1)
	}
	if name != "" && (list == true || set != "") {
		fmt.Println("Only one of create, list or set flag should be provided, and not a combination")
		os.Exit(1)
	}
	if set != "" && (list == true || name != "") {
		fmt.Println("Only one of create, list or set flag should be provided, and not a combination")
		os.Exit(1)
	}
	if list == true && (name != "" || set != "") {
		fmt.Println("Only one of create, list or set flag should be provided, and not a combination")
		os.Exit(1)
	}
	fmt.Println("Stage Command valid, proceed to process", name, list, set)

}

func config(name string) {
	fmt.Println("config name to be created...", name)
}

func main() {
	fmt.Println("Main here...")

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
		validateAndProcessStage(*stagename, *stageList, *stageSet)
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

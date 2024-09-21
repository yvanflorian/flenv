package stage

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/yvanflorian/flenv/utils"
)

func Handle(args []string) {
	stageCommands := flag.NewFlagSet("stage", flag.ExitOnError)
	stagename := stageCommands.String("create", "", "Stage Name to Create")
	stageList := stageCommands.Bool("list", false, "List the available stages")
	stageSet := stageCommands.String("set", "", "Set all environment variables for the given stagename")

	stageCommands.Usage = func() {
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "USAGE\n")
		fmt.Fprintf(os.Stderr, "  flenv stage [flags]\n")
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "AVAILABLE FLAGS\n")
		fmt.Fprintf(os.Stderr, "  create: Create a new stage.\n")
		fmt.Fprintf(os.Stderr, "  list: View list of stages managed by flenv.\n")
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "EXAMPLE\n")
		fmt.Fprintf(os.Stderr, " #Create a new Stage called prod\n")
		fmt.Fprintf(os.Stderr, " $flenv stage --create prod\n")
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, " #List down all stages that flenv is managing\n")
		fmt.Fprintf(os.Stderr, " $flenv config --list\n")
		fmt.Fprintf(os.Stderr, "\n")
		// configCommands.PrintDefaults()
	}
	stageCommands.Parse(os.Args[2:])
	utils.NoEmptyFlags(stageCommands)
	ValidateAndProcessStage(*stagename, *stageList, *stageSet)
}

func ValidateAndProcessStage(name string, list bool, set string) {
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
	ProcessStage(name, list, set)
}

func ProcessStage(name string, list bool, set string) {
	if name != "" {
		err := CreateNewStage(name)
		if err != nil {
			log.Fatalf("Failed to create a new stage: %v", err)
		}
	}

	if list {
		stages, err := ListStages()
		if err != nil {
			log.Fatalf("Failure to List stages: %v", err)
		}
		fmt.Println("Available stage(s):")
		for _, stage := range stages {
			fmt.Println(" >", stage)
		}
	}

	if set != "" {
		log.Println("TODO: setting envs belonging to stage", set)
	}
}

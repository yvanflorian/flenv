package variable

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/yvanflorian/flenv/utils"
)

func Handle(args []string) {
	varCommands := flag.NewFlagSet("variable", flag.ExitOnError)
	varConfig := varCommands.String("config", "", "Config Name that holds this variable")
	varCreate := varCommands.String("create", "", "Variable name to Create")
	varShow := varCommands.String("show", "", "Display the Variable value in given stages")
	varEdit := varCommands.String("edit", "", "Variable name to edit")
	varStage := varCommands.String("stage", "", "Stage that owns the config and the variable")

	varCommands.Usage = utils.ShowDocsVariable

	varCommands.Parse(os.Args[2:])
	utils.NoEmptyFlags(varCommands)
	ValidateAndProcess(*varCreate, *varShow, *varEdit, *varConfig, *varStage)
}

func ValidateAndProcess(
	create string,
	show string,
	edit string,
	config string,
	stage string,
) {
	if create == "" && config == "" && show == "" && edit == "" {
		fmt.Println("Either the create, config, show or edit flags must be provided")
		os.Exit(1)
	}

	err := onlyOneFlag(create, show, edit)
	if err != nil {
		fmt.Println("Only one of 'create', 'show' or 'edit' flag should be provided, and not a combination")
		os.Exit(1)
	}

	if config == "" {
		fmt.Println("A configuration that owns this variable must be provided with the 'config' flag!")
		os.Exit(1)
	}

	if edit != "" && stage == "" {
		fmt.Println("Editing a variable requires specifying the stage")
		os.Exit(1)
	}
	ProcessVariable(create, show, edit, config, stage)
}

func ProcessVariable(
	create string,
	show string,
	edit string,
	config string,
	stage string,
) {
	// log.Printf("Processing create %v show %v edit %v config %v\n", create, show, edit, config)
	if create != "" {
		err := CreateNewVariable(create, config)
		if err != nil {
			log.Fatalf("Failed to create a new stage: %v", err)
		}
	}

	if show != "" {
		err := ShowVariable(config, show)
		if err != nil {
			log.Fatalf("Failure to show variable. Error: %v", err)
		}
	}

	if edit != "" {
		err := EditVariable(config, edit, stage)
		if err != nil {
			log.Fatalf("Failure to show variable. Error: %v", err)
		}
	}
}

// Only one flag should be provided
// Either 'create', 'show' or 'edit' but not both at the same time...
func onlyOneFlag(create string, show string, edit string) error {
	countNonEmpty := 0
	if create != "" {
		countNonEmpty++
	}

	if show != "" {
		countNonEmpty++
	}

	if edit != "" {
		countNonEmpty++
	}
	if countNonEmpty > 1 {
		return errors.New("More than one flag provided...")
	}
	return nil
}

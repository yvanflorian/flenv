package stage

import (
	"fmt"
	"log"
	"os"
)

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
}

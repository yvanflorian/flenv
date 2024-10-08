package printer

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/yvanflorian/flenv/utils"
)

//print out the entire config

func Handle(args []string) {
	printCommand := flag.NewFlagSet("print", flag.ExitOnError)
	printCommand.Usage = utils.ShowDocsPrint
	printCommand.Parse(args)
	printProcess()
}

func printProcess() {
	currentConfig, err := utils.ReadConfigFile()
	if err != nil {
		log.Fatalf("failure reading config file: %v", err)
	}

	prettyJson, err := json.MarshalIndent(currentConfig, "", " ")
	if err != nil {
		log.Fatalf("Error MarshalIndent(): %v", err)
	}

	fmt.Println(string(prettyJson))

}

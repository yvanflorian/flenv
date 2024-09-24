package main

import (
	"fmt"
	"os"

	"github.com/yvanflorian/flenv/config"
	"github.com/yvanflorian/flenv/stage"
	"github.com/yvanflorian/flenv/utils"
	"github.com/yvanflorian/flenv/variable"
)

func main() {
	if len(os.Args) < 2 {
		utils.ShowDocsMain()
		os.Exit(0)
	}

	switch os.Args[1] {
	case "stage":
		stage.Handle(os.Args[2:])
	case "config":
		config.Handle(os.Args[2:])
	case "variable":
		variable.Handle(os.Args[2:])
	default:
		fmt.Println("Wrong command: Either 'stage', 'config' or 'variable'")
		os.Exit(1)
	}
}

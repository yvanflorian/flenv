package main

import (
	"fmt"
	"os"

	"github.com/yvanflorian/flenv/config"
	"github.com/yvanflorian/flenv/stage"
	"github.com/yvanflorian/flenv/variable"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println(" ")
		fmt.Println("Flenv: A Basic environment variables setter/manager for your terminal.")
		fmt.Println("")
		fmt.Println("USAGE")
		fmt.Println("  flenv <command> [flags]")
		fmt.Println("")
		fmt.Println("COMMANDS")
		fmt.Println("  stage: View, Manage your stages")
		fmt.Println("  config: View & Manage your stage configurations")
		fmt.Println("  variable: View & Manage specific configuration variables")
		fmt.Println("")
		fmt.Println("LEARN MORE:")
		fmt.Println(" Use `flenv <command> --help for more information about the given command.`")
		fmt.Println("")
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

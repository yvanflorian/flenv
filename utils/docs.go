package utils

import (
	"fmt"
	"os"
)

//Show docs using the Cobra style

// flenv
func ShowDocsMain() {
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
}

// flenv stage --help
func ShowDocsStage() {
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
	fmt.Fprintf(os.Stderr, " $flenv stage --list\n")
	fmt.Fprintf(os.Stderr, "\n")
}

// flenv config --help
func ShowDocsConfig() {
	fmt.Fprintf(os.Stderr, "\n")
	fmt.Fprintf(os.Stderr, "USAGE\n")
	fmt.Fprintf(os.Stderr, "  flenv config [flags]\n")
	fmt.Fprintf(os.Stderr, "\n")
	fmt.Fprintf(os.Stderr, "AVAILABLE FLAGS\n")
	fmt.Fprintf(os.Stderr, "  create: Create a new config. Multiple Prompts will follow to create the same config for all stages\n")
	fmt.Fprintf(os.Stderr, "  list: View & Manage specific configuration variables\n")
	fmt.Fprintf(os.Stderr, "\n")
	fmt.Fprintf(os.Stderr, "EXAMPLE\n")
	fmt.Fprintf(os.Stderr, " #Create a new config called database\n")
	fmt.Fprintf(os.Stderr, " $flenv config --create database\n")
	fmt.Fprintf(os.Stderr, "\n")
	fmt.Fprintf(os.Stderr, " #List down all config that flenv is managing\n")
	fmt.Fprintf(os.Stderr, " $flenv config --list\n")
	fmt.Fprintf(os.Stderr, "\n")
}

// flenv variable --help
func ShowDocsVariable() {
	fmt.Fprintf(os.Stderr, "\n")
	fmt.Fprintf(os.Stderr, "USAGE\n")
	fmt.Fprintf(os.Stderr, "  flenv variable [flags] [targeted-flags]\n")
	fmt.Fprintf(os.Stderr, "\n")
	fmt.Fprintf(os.Stderr, "GENERAL FLAGS\n")
	fmt.Fprintf(os.Stderr, "  create: Create a new Variable\n")
	fmt.Fprintf(os.Stderr, "  show: Display the value of the given variable name\n")
	fmt.Fprintf(os.Stderr, "  edit: Modify the value of a given variable name\n")
	fmt.Fprintf(os.Stderr, "\n")
	fmt.Fprintf(os.Stderr, "TARGETED FLAGS\n")
	fmt.Fprintf(os.Stderr, "  config: Specific Config that owns the variable\n")
	fmt.Fprintf(os.Stderr, "  stage: Stage name that owns the config and the variable specified\n")
	fmt.Fprintf(os.Stderr, "\n")
	fmt.Fprintf(os.Stderr, "EXAMPLE\n")
	fmt.Fprintf(os.Stderr, " #Create a new variable called host for the redis config\n")
	fmt.Fprintf(os.Stderr, " $flenv variable --create host --config redis\n")
	fmt.Fprintf(os.Stderr, "\n")
	fmt.Fprintf(os.Stderr, " #Display the variable value for the redis host \n")
	fmt.Fprintf(os.Stderr, " $flenv variable --show host --config redis --stage prod \n")
	fmt.Fprintf(os.Stderr, "\n")
	fmt.Fprintf(os.Stderr, " #Modify the Redis Hostname for the production stage\n")
	fmt.Fprintf(os.Stderr, " $flenv variable --edit host --config redis --stage prod\n")
	fmt.Fprintf(os.Stderr, "\n")
	fmt.Fprintf(os.Stderr, "\n")
	fmt.Fprintf(os.Stderr, " #List all variables available for a given config\n")
	fmt.Fprintf(os.Stderr, " $flenv variable --list --config redis\n")
	fmt.Fprintf(os.Stderr, "\n")

}

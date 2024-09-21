package utils

import (
	"flag"
	"fmt"
	"os"
)

func NoEmptyFlags(cmd *flag.FlagSet) {
	if cmd.NFlag() == 0 {
		fmt.Printf("Error: The '%v' subcommand requires flags. Please review!\n", cmd.Name())
		// cmd.PrintDefaults()
		cmd.Usage()
		os.Exit(1)
	}
}

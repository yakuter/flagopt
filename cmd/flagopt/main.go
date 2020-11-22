package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/yakuter/flagopt/config"
)

func main() {
	// Create a FlagSet and set the usage
	fs := flag.NewFlagSet(filepath.Base(os.Args[0]), flag.ExitOnError)

	// Configure the options from the flags/config file
	opts, err := config.ConfigureOptions(fs, os.Args[1:])
	if err != nil {
		config.PrintUsageErrorAndDie(err)
	}

	// If -help flag is defined, print help
	if opts.ShowHelp {
		config.PrintHelpAndDie()
	}

	fmt.Println(opts)
}

package config

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/fatih/color"
)

var usageStr = `
Usage: flagopt [options]
Options:
    -n, --name <name>       Your name. i.e. John
    -e, --email <email>     Your email address. i.e. john@doe.com
    -a, --age <age>         Your age. i.e. 26
    -m, --mask <mask>	    (Default False) Do you wear mask? i.e. true
    -h, --help              (Optional) Show help message
`

// PrintUsageErrorAndDie ...
func PrintUsageErrorAndDie(err error) {
	color.Red(err.Error())
	fmt.Println(usageStr)
	os.Exit(1)
}

// PrintHelpAndDie ...
func PrintHelpAndDie() {
	fmt.Println(usageStr)
	os.Exit(0)
}

// Options is main value holder for command line arguments/flags.
type Options struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
	Mask     bool   `json:"mask"`
	ShowHelp bool   `json:"show_help"`
}

// ConfigureOptions accepts a flag set and augments it with project's
// specific flags. On success, an options structure is returned configured
// based on the selected flags.
func ConfigureOptions(fs *flag.FlagSet, args []string) (*Options, error) {

	// Create empty options
	opts := &Options{}

	// Define flags
	fs.StringVar(&opts.Name, "n", "", "Your name. i.e. John")
	fs.StringVar(&opts.Name, "name", "", "Your full name. i.e. John Doe")
	fs.StringVar(&opts.Email, "e", "", "Your email address. i.e. john@doe.com")
	fs.StringVar(&opts.Email, "email", "", "Your email address. i.e. john@doe.com")
	fs.IntVar(&opts.Age, "a", 0, "Your age. i.e. 26")
	fs.IntVar(&opts.Age, "age", 0, "Your age. i.e. 26")
	fs.BoolVar(&opts.Mask, "m", false, "Do you wear mask? i.e. True")
	fs.BoolVar(&opts.Mask, "mask", false, "Do you wear mask? i.e. True")
	fs.BoolVar(&opts.ShowHelp, "h", false, "Show help message")
	fs.BoolVar(&opts.ShowHelp, "help", false, "Show help message")

	// Parse arguments and check for errors
	if err := fs.Parse(args); err != nil {
		return nil, err
	}

	// If it is not help and other args are empty, return error
	if (opts.ShowHelp == false) && (opts.Name == "" || opts.Email == "" || opts.Age == 0) {
		err := errors.New("please specify all arguments")
		return nil, err
	}

	return opts, nil
}

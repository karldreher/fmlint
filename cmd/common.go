package cmd

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

// handleErrors checks if any errors occurred during the execution of the command.
// If so, it prints the error message and exits the program.
// If --warn-only is set, it prints the error message and continues.
// When the function is called, it expects a boolean value indicating whether
// or not an error occurred.
func handleErrors(hasError bool) {
	warn := viper.GetString("warn")
	if hasError && warn != "true" {
		os.Exit(1)
	}
	if hasError && warn == "true" {
		log.Println("Warning: 1 or more errors occurred but --warn-only is set.")
	}

}

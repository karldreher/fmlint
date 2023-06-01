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

// ruleEnabled checks if a rule is enabled.
// It expects a rule ID as input.
// Rules are always enabled unless disabled in the config file.
// Rules are disabled if they are in the yaml config file in the map "disabled_rules"."
// If the rule is disabled, it returns false.
// If the rule is enabled, it returns true.
func ruleEnabled(ruleID string) bool {
	config := viper.GetStringSlice("disabled_rules")
	for _, rule := range config {
		if rule == ruleID {
			return false
		}
	}
	return true
}

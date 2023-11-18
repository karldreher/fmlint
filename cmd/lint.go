package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// lintCmd represents the lint command.
// By itself this does nothing and exits fatally.
// Note, subcommands of "lint" are still cobra.Commands but should be added to lintCmd.
var lintCmd = &cobra.Command{
	Use:   "lint",
	Short: "Execute lint subcommands.",
	Long:  `Lint front matter based on provided rules.  To find avaialble rules, run "fmlint list".`,
	Run: func(cmd *cobra.Command, args []string) {
		// if there were no arguments, print usage
		if len(args) == 0 {
			fmt.Println("Error: No lint subcommand provided")
			//nolint:errcheck
			cmd.Help()
			os.Exit(1)
		}
		// if all was specified,
		if cmd.CalledAs() == "lint all" {
			//lint all
			println("hello world")

		}
	},
}

func init() {
	rootCmd.AddCommand(lintCmd)
}

// handleErrors checks if any errors occurred during the execution of the command.
// If so, it prints the error message and exits the program.
// If --warn-only is set, it prints the error message and continues.
// If lint was called with "lint all", it does not exit and returns an error instead.
// When the function is called, it expects a boolean value indicating whether
// or not an error occurred.
func handleErrors(hasError bool) {
	warn := viper.GetViper().GetString("warn")
	//lint_all only is set by allCmd.
	lint_all := viper.GetViper().GetBool("lint_all")
	if hasError && !lint_all && warn != "true" {
		os.Exit(1)
	}
	if hasError && warn == "true" {
		log.Println("Warning: 1 or more errors occurred but --warn-only is set.")
	}
	if hasError && lint_all {
		// Special viper setting used internally.  Not exposed as a flag.
		// While it could technically be used through a config file,
		// it would be abnormal and not as-designed.
		// If lint_all is set, it is considered a "soft" failure per-rule, but then
		// allCmd will register the failure and exit abnormally.
		viper.Set("lint_all_fail", true)
	}
}

// ruleEnabled checks if a rule is enabled.
// It expects a rule ID as input.
// Rules are always enabled unless disabled in the config file.
// Rules are disabled if they are in the yaml config file in the map "disabled_rules"."
// If the rule is disabled, it returns false.
// If the rule is enabled, it returns true.
func ruleEnabled(ruleID string) bool {
	config := viper.GetViper().GetStringSlice("disabled_rules")
	for _, rule := range config {
		if rule == ruleID {
			log.Printf("Rule \"%s\" is disabled, not linting", ruleID)
			return false
		}
	}
	return true
}

// A function which takes a filepath and returns a boolean value.
type lintRule func(path string) bool

// Generic function to walk directories and evaluate lint rules.
// Requires that a lint rule function is passed to it.  This should be a function which
// returns true/false based on rule evaluation.
func evaluateRules(fn lintRule) {
	folder := viper.GetViper().GetString("folder")
	//Sets hasErr to false, until altered by individual lint rules
	hasErr := false
	err := filepath.Walk(folder,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil
			}
			check := fn(path)
			if !check {
				hasErr = true
			}
			return nil
		})
	//Handle errors from the filepath walk
	if err != nil {
		log.Println(err)
	}
	//Handle errors from the lint function, if more than zero errors are present
	handleErrors(hasErr)
}

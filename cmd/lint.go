/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// lintCmd represents the lint command.
// By itself this does nothing and exits fatally.
// Note, subcommands of "lint" are still cobra.Commands but should be added to lintCmd.
var lintCmd = &cobra.Command{
	Use:   "lint",
	Short: "Execute lint subcommands",
	Long:  `Lint front matter based on provided rules.  To find avaialble rules, run "fmlint list".`,
	Run: func(cmd *cobra.Command, args []string) {
		// if there were no arguments, print usage
		if len(args) == 0 {
			cmd.Usage()
			fmt.Println("Error: No lint subcommand provided")
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(lintCmd)
}

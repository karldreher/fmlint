package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// allCmd represents the all command
var allCmd = &cobra.Command{
	Use:         "all",
	Annotations: map[string]string{"rule-id": "none"},
	Short:       "Run all lint commands.",
	Long:        ``,
	Run: func(cmd *cobra.Command, args []string) {
		viper.Set("lint_all", true)
		//RUn all commands with rule-id annotation.
		cmdList := lintCmd.Commands()
		for _, command := range cmdList {
			if command.Annotations["rule-id"] != "none" {
				command.Run(cmd, args)
			}

		}
		// If a failure was encountered and recorded by handleError, exit.
		if viper.GetBool("lint_all_fail") {
			os.Exit(1)
		}
	},
}

func init() {
	lintCmd.AddCommand(allCmd)

}

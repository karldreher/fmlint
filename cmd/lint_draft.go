package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// draftCmd represents the draft command
var draftCmd = &cobra.Command{
	Use:         "draft",
	Annotations: map[string]string{"rule-id": "draft-enabled"},
	Short:       "Checks that draft is not set to \"true\" for all files.",
	Long: `Draft mode should not be enabled.  
For the most part, draft mode should be avoided before content goes into production.
If draft: false, this lint rule will pass.  
If draft: true, this lint will trigger a failure.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO
		fmt.Println("draft called")
	},
}

func init() {
	// Note, subcommands of "lint" are still cobra.Commands but should be added to lintCmd.
	lintCmd.AddCommand(draftCmd)
}

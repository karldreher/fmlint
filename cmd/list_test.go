package cmd

import (
	"testing"

	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Test Command",
	Long: `This is only a test command and is not added to the root command,
	outside of the test context.
	The only purpose for this command is to test the rule-id annotation.`,
	Run: func(cmd *cobra.Command, args []string) {
	}}

func init() {
	rootCmd.AddCommand(testCmd)
}

// This test is not specific to list.go.  All commands added to this program
// should have a rule-id annotation.  In the case that a  command does not have
// a "categorizable" rule, rule-id should be annotated with "none".
func TestMain(t *testing.T) {
	cmdList := rootCmd.Commands()
	for _, command := range cmdList {
		//Print command name
		ruleID := command.Annotations["rule-id"]
		println(ruleID)
		if ruleID == "" {
			t.Errorf("command %s has no rule-id annotation", command.Name())
			t.Fail()
		}
	}

}

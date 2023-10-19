package cmd

import (
	"testing"
)

// All sub-commands added to the lintCmd
// should have a rule-id annotation.  In the case that a  command does not have
// a "categorizable" rule, rule-id should be annotated with "none".
func Test_lintCmd_rule_id_annotations(t *testing.T) {
	cmdList := lintCmd.Commands()
	for _, command := range cmdList {
		// Ensure all sub-commands of lintCmd have a rule-id annotation.
		ruleID := command.Annotations["rule-id"]
		if ruleID == "" {
			t.Errorf("command %s has no rule-id annotation", command.Name())
			t.Fail()
		}
	}

}

// This test is not specific to lint.go.
// rootCmd entries are expected not to have a rule-id annotation.
func Test_rootCmd_rule_id_annotations(t *testing.T) {
	cmdList := rootCmd.Commands()
	for _, command := range cmdList {

		// If the command has command.Annotations["rule-id"], fail the test.
		// This is not expected on root command entries.
		ruleID := command.Annotations["rule-id"]
		if ruleID != "" {
			t.Errorf("command %s has rule-id annotation", command.Name())
			t.Fail()
		}
	}

}

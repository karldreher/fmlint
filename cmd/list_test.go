package cmd

import (
	"testing"
)

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

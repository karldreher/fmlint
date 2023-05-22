package cmd

import (
	"testing"
)

func Test_passing_checkTags(t *testing.T) {
	// Checks that a file with properly sorted tags passes.
	if checkTags("../test/tags_pass.md") == false {
		t.Error("Expected tags_pass.md to pass but it failed")
		t.Fail()
	}
}
func Test_failing_checkTags(t *testing.T) {
	// Checks that a file with improperly sorted tags fails.
	if checkTags("../test/tags_fail.md") == true {
		t.Error("Expected tags_fail.md to fail but it passed")
		t.Fail()
	}
}

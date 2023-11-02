package cmd

import (
	"log"
	"os"
	"testing"
)

// Checks that a file with properly sorted tags passes.
func Test_passing_checkTags(t *testing.T) {
	// Redirect log output to /dev/null to avoid test result pollution.
	null, _ := os.Open(os.DevNull)
	log.SetOutput(null)
	if checkTags("../test/lint_tags/tags_pass.md") == false {
		t.Error("Expected tags_pass.md to pass but it failed")
		t.Fail()
	}
}

// Checks that a file with improperly sorted tags fails.
func Test_failing_checkTags(t *testing.T) {
	// Redirect log output to /dev/null to avoid test result pollution.
	null, _ := os.Open(os.DevNull)
	log.SetOutput(null)
	if checkTags("../test/lint_tags/tags_fail.md") == true {
		t.Error("Expected tags_fail.md to fail but it passed")
		t.Fail()
	}
}

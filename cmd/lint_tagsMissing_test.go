package cmd

import (
	"log"
	"os"
	"testing"
)

func Test_passing_checkTagsMissing(t *testing.T) {
	// Redirect log output to /dev/null to avoid test result pollution.
	null, _ := os.Open(os.DevNull)
	log.SetOutput(null)
	if checkTagsPresent("../test/lint_tagsMissing/pass.md") == false {
		t.Error("Expected tags_pass.md to pass but it failed")
		t.Fail()
	}
}

// Create the other cases for fail.md and fail2.md.
func Test_failing_checkTagsMissing(t *testing.T) {
	// Redirect log output to /dev/null to avoid test result pollution.
	// null, _ := os.Open(os.DevNull)
	// log.SetOutput(null)
	if checkTagsPresent("../test/lint_tagsMissing/fail.md") == true {
		t.Error("Expected tags_fail.md to fail but it passed")
		t.Fail()
	}
	if checkTagsPresent("../test/lint_tagsMissing/fail2.md") == true {
		t.Error("Expected tags_fail2.md to fail but it passed")
		t.Fail()
	}
}

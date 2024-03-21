package cmd

import (
	"log"
	"os"
	"testing"
)

// Checks that a file with properly sorted tags passes.
func Test_passing_checkTagSort(t *testing.T) {
	// Redirect log output to /dev/null to avoid test result pollution.
	null, _ := os.Open(os.DevNull)
	log.SetOutput(null)
	if checkTagSort("../test/lint_tags_sort/tags_pass.md") == false {
		t.Error("Expected tags_pass.md to pass but it failed")
		t.Fail()
	}
}

// Checks that a file with improperly sorted tags fails.
func Test_failing_checkTagSort(t *testing.T) {
	// Redirect log output to /dev/null to avoid test result pollution.
	null, _ := os.Open(os.DevNull)
	log.SetOutput(null)
	if checkTagSort("../test/lint_tags_sort/tags_fail.md") == true {
		t.Error("Expected tags_fail.md to fail but it passed")
		t.Fail()
	}
}

// Checks that a file that does not have tags set does not trigger the rule.
func Test_passing_checkTagSort_noTags(t *testing.T) {
	// Redirect log output to /dev/null to avoid test result pollution.
	null, _ := os.Open(os.DevNull)
	log.SetOutput(null)
	if checkTagSort("../test/lint_tags_sort/tags_not_set.md") == false {
		t.Error("Expected tags_not_set.md to pass but it failed")
		t.Fail()
	}
}

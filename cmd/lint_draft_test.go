package cmd

import (
	"log"
	"os"
	"testing"
)

// Checks that a file with draft: false will pass.
func Test_passing_checkDraft(t *testing.T) {
	// Redirect log output to /dev/null to avoid test result pollution.
	null, _ := os.Open(os.DevNull)
	log.SetOutput(null)
	if checkDraft("../test/lint_draft/draft_pass.md") == false {
		t.Error("Expected draft_pass.md to pass but it failed")
		t.Fail()
	}
}

// Checks that a file with draft: true will fail.
func Test_failing_checkDraft(t *testing.T) {
	// Redirect log output to /dev/null to avoid test result pollution.
	null, _ := os.Open(os.DevNull)
	log.SetOutput(null)
	if checkDraft("../test/lint_draft/draft_fail.md") == true {
		t.Error("Expected draft_fail.md to fail but it passed")
		t.Fail()
	}
}

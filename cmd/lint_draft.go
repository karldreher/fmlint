package cmd

import (
	"bytes"
	"log"
	"os"

	"github.com/adrg/frontmatter"
	"github.com/spf13/cobra"
)

// draftCmd represents the draft command
var draftCmd = &cobra.Command{
	Use:         "draft",
	Annotations: map[string]string{"rule-id": "draft-enabled"},
	Short:       "Checks that draft is not set to \"true\" for all files.",
	Long: `Draft mode should not be enabled.  
The check is designed to avoid drafts being enabled prior to the release of the site.  
If draft: true, this lint rule will trigger a failure.  
If draft: false, this lint rule will pass.`,
	Run: func(cmd *cobra.Command, args []string) {
		evaluateRules(checkDraft)
	},
}

func init() {
	// Note, subcommands of "lint" are still cobra.Commands but should be added to lintCmd.
	lintCmd.AddCommand(draftCmd)
}

// checkDraft checks if draft: "false" or unset.
// Returns true if draft is not set to true, otherwise false.
func checkDraft(file string) bool {
	var matter struct {
		Title string `yaml:"title"`
		Draft bool   `yaml:"draft"`
	}
	b, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	rest, err := frontmatter.Parse(bytes.NewReader(b), &matter)
	if err != nil {
		log.Println(rest, err)
	}
	if matter.Draft == true {
		log.Printf("Draft mode enabled. {\"file\": %q, \"draft\": %t}", file, matter.Draft)
		return false
	}
	return true
}

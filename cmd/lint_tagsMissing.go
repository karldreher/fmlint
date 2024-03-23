package cmd

import (
	"bytes"
	"log"
	"os"

	"github.com/adrg/frontmatter"
	"github.com/spf13/cobra"
)

func init() {
	lintCmd.AddCommand(tagsMissingCmd)
}

// tagsMissingCmd represents the tagsMissing command
var tagsMissingCmd = &cobra.Command{
	Use:         "tags-missing",
	Annotations: map[string]string{"rule-id": "tags-missing"},
	Short:       "Lint for missing (null) tags",
	Long: `Tags in frontmatter are expected to be a YAML list.
	This command checks to ensure at least one tag is present.`,
	Run: func(cmd *cobra.Command, args []string) {
		if ruleEnabled("tags-sorted") {
			evaluateRules(checkTagsPresent)
		}
	},
}

// checkTagsPresent checks to see if the file has at least one tag.
// Returns true if the file has at least one item in the tags list.
// Returns false otherwise.
func checkTagsPresent(file string) bool {
	var matter struct {
		Tags []string `yaml:"tags"`
	}
	b, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	rest, err := frontmatter.Parse(bytes.NewReader(b), &matter)
	if err != nil {
		log.Println(rest, err)
	}
	if matter.Tags != nil {
		// If any item in the tags list is empty, return false.
		for _, tag := range matter.Tags {
			if len(tag) == 0 {
				log.Printf("Tags are not present.  {\"file\": %q, \"tags\": %q}", file, matter.Tags)
				return false
			}
		}
		// Otherwise, return true.
		if len(matter.Tags) > 0 {

			return true
		}

	}
	// TODO Not ideal that this log is repeated in this function,
	// Consider refactoring but possibly considering all lint failure logging in a central facility
	log.Printf("Tags are not present.  {\"file\": %q, \"tags\": %q}", file, matter.Tags)
	return false

}

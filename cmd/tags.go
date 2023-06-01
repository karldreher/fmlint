package cmd

import (
	"bytes"
	"log"
	"os"
	"path/filepath"
	"sort"

	"github.com/adrg/frontmatter"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(tagsCmd)
}

var tagsCmd = &cobra.Command{
	Use:         "tags",
	Annotations: map[string]string{"rule-id": "tags-sorted"},
	Short:       "Lint tag sorting",
	Long: `Tags in frontmatter are expected to be a YAML list.
	This command checks to ensure they are sorted alphabetically.`,
	Run: func(cmd *cobra.Command, args []string) {
		hasErr := false
		//recursively walk the "content" directory and find all the files
		//that have a frontmatter
		folder := viper.GetString("folder")
		err := filepath.Walk(folder,
			func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if info.IsDir() {
					return nil
				}
				check := checkTags(path)
				if !check {
					hasErr = true
				}
				return nil
			})
		if err != nil && ruleEnabled("tags-sorted") {
			log.Println(err)
		}
		handleErrors(hasErr)
	},
}

// checkTags checks if the tags are sorted alphabetically.
// Returns true if sorted, false if not.
func checkTags(file string) bool {
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

	// Check if tags are sorted
	sortedTags := sort.SliceIsSorted(matter.Tags, func(i, j int) bool {
		return matter.Tags[i] < matter.Tags[j]
	})
	if !sortedTags {
		log.Printf("Tags are not sorted.  {\"file\": %q, \"tags\": %q}", file, matter.Tags)
		return false
	}
	return true
}

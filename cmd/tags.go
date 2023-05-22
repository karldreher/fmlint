package cmd

import (
	"bytes"
	"fmt"
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
	Use:   "tags",
	Short: "Lint tags in frontmatter",
	Long: `Tags are expected to be a YAML list.
	This subcommand checks to ensure they are sorted alphabetically.`,
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
				if check == false {
					hasErr = true
				}
				return nil
			})
		if err != nil {
			log.Println(err)
		}
		if hasErr {
			os.Exit(1)
		}
	},
}

func checkTags(file string) bool {
	var matter struct {
		Name string   `yaml:"name"`
		Tags []string `yaml:"tags"`
	}
	b, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	rest, err := frontmatter.Parse(bytes.NewReader(b), &matter)
	if err != nil {
		fmt.Println(rest, err)
	}

	// Check if tags are sorted
	sortedTags := sort.SliceIsSorted(matter.Tags, func(i, j int) bool {
		return matter.Tags[i] < matter.Tags[j]
	})
	if sortedTags == false {

		fmt.Println("Tags are not sorted.", "tags:", matter.Tags, "file:", file)
		fmt.Println(file, matter)
		return false
	}
	return true
}
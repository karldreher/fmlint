/*
Copyright © 2023 Karl Dreher

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
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

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "fmlint",
	Short: "Lint your front-matter markdown files to ensure that tags are sorted.",
	Long: `

`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
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

	//Check if tags are sorted
	sortedTags := sort.SliceIsSorted(matter.Tags, func(i, j int) bool {
		return matter.Tags[i] < matter.Tags[j]
	})
	if sortedTags == false {

		fmt.Println("Tags are not sorted.", "tags:", matter.Tags, "file:", file)
		fmt.Println(file, matter)
		//os.Exit(1)
		return false
	}
	return true
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize()

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.fmlint.yaml)")
	rootCmd.PersistentFlags().StringP("folder", "f", "./content", "Folder to recursively scan for front-matter markdown files.")
	viper.BindPFlag("folder", rootCmd.PersistentFlags().Lookup("folder"))
}

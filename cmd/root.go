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
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "fmlint",
	Short: "Lint your front-matter markdown files.",
	Long: `Lint your front-matter markdown files.
This must be run with a subcommand, as detailed below.
Currently, the only supported subcommand is "tags".
  For more information, try 'fmlint tags --help'.

`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			// This is configured to exit with a non-zero exit code,
			// to prevent "false negative" results which appear to be passing, but are actually misconfigurations.
			os.Exit(1)
		}
	},
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
	rootCmd.PersistentFlags().StringP("folder", "f", "./content", "Folder to recursively scan for frontmatter markdown files.")
	rootCmd.PersistentFlags().BoolP("warn-only", "", false, "Do not fail if errors are encountered, but print warnings.")

	viper.BindPFlag("folder", rootCmd.PersistentFlags().Lookup("folder"))
	viper.BindPFlag("warn", rootCmd.PersistentFlags().Lookup("warn-only"))

}

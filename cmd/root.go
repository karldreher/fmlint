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
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "fmlint",
	Short: "Lint your markdown front-matter.",
	Long: `Lint your markdown front-matter.
For more information, try '--help' after any sub-command.

`,

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			err := cmd.Help()
			if err != nil {
				panic(err)
			}

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

func SetVersionInfo(version, commit, date string) {
	rootCmd.Version = fmt.Sprintf("%s (Built on %s)", version, date)
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.fmlint.yaml)")
	rootCmd.PersistentFlags().StringP("folder", "f", "./content", "Folder to recursively scan for frontmatter markdown files.")
	rootCmd.PersistentFlags().BoolP("warn-only", "", false, "Do not fail if errors are encountered, but print warnings.")
	//nolint:errcheck
	viper.BindPFlag("folder", rootCmd.PersistentFlags().Lookup("folder"))
	//nolint:errcheck
	viper.BindPFlag("warn", rootCmd.PersistentFlags().Lookup("warn-only"))

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
		_, err := os.Stat(cfgFile)
		// If the specified file does not exist, fail.
		cobra.CheckErr(err)

	} else {
		// use the working directory as a primary search path
		pwd, err := os.Getwd()
		cobra.CheckErr(err)
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search directories for a config file.
		// The first added paths take precedence.
		viper.AddConfigPath(pwd)
		viper.AddConfigPath(home)

		viper.SetConfigType("yaml")
		viper.SetConfigName(".fmlint")
		// when the config file is found, print it
		if viper.ConfigFileUsed() != "" {
			fmt.Printf("Config file name: %s\n", viper.ConfigFileUsed())
		}
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		log.Printf("Using config file: %s", viper.ConfigFileUsed())

	}
}

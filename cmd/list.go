package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List lint rules",
	Long:  `Lists all commands and their corresponding lint rules.`,
	Run: func(cmd *cobra.Command, args []string) {
		//Get all commands
		cmdList := lintCmd.Commands()
		fmt.Printf("%s \t %s \t\t %s\n", "Command", "Rule-ID", "Short Description")
		//print a separator
		fmt.Println("--------------------------------------------------------------------------")
		for _, command := range cmdList {

			if command.Annotations != nil {
				//If command.Annotations has key "rule-id", print it
				if _, ok := command.Annotations["rule-id"]; ok {
					fmt.Printf("%s \t\t %s \t\t %s\n", command.Name(), command.Annotations["rule-id"], command.Short)
				}
			}

		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

}

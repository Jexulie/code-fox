package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:     "codefox",
	Short:   "Code snippet manager.",
	Long:    "Code snippet manager with search & tags",
	Version: "1.0",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`Welcome to CodeFox!`)
		fmt.Println(`This cli tool is for managing code snippets`)
		fmt.Println(`-h for help.`)
	},
}

// InitializeCLI initializes CLI commands
func init() {
	RootCmd.AddCommand(saveCmd)
	addSaveCmdFlags()

	RootCmd.AddCommand(listCmd)
	addListCmdFlags()

	RootCmd.AddCommand(getCmd)

	RootCmd.AddCommand(editCmd)
	addEditCmdFlags()

	RootCmd.AddCommand(deleteCmd)
}

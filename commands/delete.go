package commands

import (
	"code-fox/snippet"
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

/**
Deletes snippets - cmd "delete"
*/

var deleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Delete a snippet by id",
	Long:  "Delete a snippet by id",
	Aliases: []string{
		"d",
	},
	Run: func(cmd *cobra.Command, args []string) {
		idStrValue := args[0]
		idValue, err := strconv.ParseInt(idStrValue, 10, 64)
		if err != nil {
			fmt.Println("invalid id")
		}

		snip, err := snippet.GetById(idValue)
		if err != nil {
			fmt.Println("snippet not found")
			return
		}

		if err := snip.Delete(); err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("snippet deleted")
	},
}

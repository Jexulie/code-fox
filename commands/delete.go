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
	Run: func(cmd *cobra.Command, args []string) {
		idStrValue := args[0]
		idValue, err := strconv.ParseInt(idStrValue, 10, 64)
		if err != nil {
			fmt.Println("invalid id")
		}

		snip := snippet.Snippet{
			Id: idValue,
		}

		if err := snip.Delete(); err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("snippet deleted")
	},
}

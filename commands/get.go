package commands

import (
	"code-fox/snippet"
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

/**
Gets snippet - cmd "get"
*/

var getCmd = &cobra.Command{
	Use:   "get [id|title]",
	Short: "Gets snippet by id or title",
	Long:  "Gets snippet by id or title",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var snip *snippet.Snippet

		idStrValue := args[0]
		idValue, err := strconv.ParseInt(idStrValue, 10, 64)
		if err != nil {
			// try to get by title
			snip, err = snippet.GetByTitle(idStrValue)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
		} else {
			snip, err = snippet.GetById(idValue)
			if err != nil {
				fmt.Println(err)
				return
			}
		}

		if snip.Id == 0 {
			fmt.Println("snippet not found")
			return
		}

		fmt.Println()
		fmt.Printf("Title: %s\n", snip.Title)
		fmt.Printf("Lang: %s\n", snip.Language)
		fmt.Printf("Tags: %s\n", snip.Tags)
		fmt.Printf("Created: %s\n", snip.CreatedAt.Format("2006-01-02 15:04:05"))
		if snip.UpdatedAt != nil {
			fmt.Printf("Updated: %s\n", snip.UpdatedAt.Format("2006-01-02 15:04:05"))
		}
		fmt.Println()
		fmt.Println(snip.Code)
	},
}
